package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/antoniodipinto/ikisocket"
	"github.com/boombaw/go-ws-sia/pkg/external/feeder"

	akm "github.com/boombaw/go-ws-sia/pkg/http/akm/action"
	kelas "github.com/boombaw/go-ws-sia/pkg/http/kelas/action"
	mhs "github.com/boombaw/go-ws-sia/pkg/http/mhs/action"
	nilai "github.com/boombaw/go-ws-sia/pkg/http/nilai/action"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/repository"
	"github.com/boombaw/go-ws-sia/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

var (
	clients = make(map[string]string)
	// wsChan  = make(chan model.MessageObject)
)

type UserSocketID struct {
	UUID uuid.UUID
}

func Routes(app *fiber.App) {

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Setup the middleware to retrieve the data sent in first GET request
	app.Use(func(c *fiber.Ctx) error {
		if c.Get("host") == "ws.ubharajaya.ac.id" {
			c.Locals("Host", "ws.ubharajaya.ac.id")
			return c.Next()
		}
		// log.Println(websocket.IsWebSocketUpgrade(c))
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			// c.Locals("allowed", true)
			return c.Next()
		}

		log.Printf("Error Upgrade %v", c)
		// return fiber.ErrUpgradeRequired
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Use("/ws/:id", func(c *fiber.Ctx) error {
		log.Println("Host", c.Get("host"))
		if c.Get("host") == "ws.ubharajaya.ac.id" {
			c.Locals("Host", "ws.ubharajaya.ac.id")
			return c.Next()
		}

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		log.Println("IsWebSocketUpgrade", websocket.IsWebSocketUpgrade(c))
		return c.Status(403).SendString("Request origin not allowed")
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Hello, World!",
	// 	})
	// })

	// Multiple event handling supported
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		log.Println("Connection event 1 - User: ", fmt.Sprintf("%v", ep.Kws.GetStringAttribute("user_id")))
	})

	// Custom event handling supported
	ikisocket.On("CUSTOM_EVENT", func(ep *ikisocket.EventPayload) {
		log.Println("Custom event - User: ", fmt.Sprintf("%v", ep.Kws.GetStringAttribute("user_id")))
		// --->

		// DO YOUR BUSINESS HERE

		// --->
	})

	// On message event
	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {

		log.Println("Message event - User: ", fmt.Sprintf("%v - Message: %v", ep.Kws.GetStringAttribute("user_id"), string(ep.Data)))

		message := model.MessageObject{}

		// Unmarshal the json message
		// {
		//  "from": "<user-id>",
		//  "to": "<recipient-user-id>",
		//  "event": "CUSTOM_EVENT",
		//  "data": "hello"
		//}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println("Error marshalling json :", err)
			return
		}

		// Fire custom event based on some
		// business logic
		if message.Event != "" {
			ep.Kws.Fire(message.Event, []byte(message.Data))
		}

		switch message.Event {
		case "sync-akm":
			SyncAkm(ep, message)
		case "sync-akm-na":
			SyncAkmNA(ep, message)
		case "sync-update-nilai":
			SyncUpdateNilai(ep, message)
		}

	})

	// On disconnect event
	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		log.Printf("Disconnection event - User: %s\n", ep.Kws.GetStringAttribute("user_id"))
		ep.Kws.Close()
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		log.Printf("Close event - User: %s\n", ep.Kws.GetStringAttribute("user_id"))
		ep.Kws.Close()
	})

	// On error event
	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		log.Printf("Error %s - User: %s\n", ep.Data, ep.Kws.GetStringAttribute("user_id"))
		log.Printf("Error: %v\n", ep.Error)
		ep.Kws.Close()
		ep.Kws.Fire("error", []byte("Error Connection"))
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
	})

	app.Get("/ws/:id", ikisocket.New(func(kws *ikisocket.Websocket) {

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// ikisocket to manage Emit/EmitTo/Broadcast
		clients[userId] = kws.GetUUID()

		// Every websocket connection has an optional session key => value storage
		kws.SetAttribute("user_id", userId)

		//Broadcast to all the connected users the newcomer
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, ikisocket.TextMessage)

		//Write welcome message
		var response model.MessageResponse
		response.Event = "hello"
		response.Message = fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)
		response.ID = kws.GetUUID()

		jsonResponse := util.ToJson(response)

		kws.Emit(jsonResponse, ikisocket.TextMessage)
	}))

	port := ":" + os.Getenv("APP_PORT")
	log.Println("Listening on port : ", port)
	app.Listen(port)
}

type ParamsAKM struct {
	KdProdi  string `json:"kd_prodi"`
	Semester string `json:"semester"`
}

func SyncAkm(ep *ikisocket.EventPayload, message model.MessageObject) {
	var response model.MessageResponse

	var paramAKM ParamsAKM
	err := json.Unmarshal([]byte(message.Data), &paramAKM)
	if err != nil {
		fmt.Println("Error marshalling json :", err)
		return
	}

	response = akm.ListAKMService(paramAKM.KdProdi, paramAKM.Semester)

	type ListAkmJson struct {
		List []struct {
			model.Akm
		} `json:"list"`
	}

	var data ListAkmJson
	_ = json.Unmarshal([]byte(response.Message), &data)

	total := []byte(`{"total" : %v}`)
	response.Event = "total"
	response.Message = fmt.Sprintf(string(total), len(data.List))

	if ep.Kws.IsAlive() {
		err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		ep.Kws.Close()
	}

	for i, v := range data.List {
		tokenFeeder := feeder.GetToken()

		if tokenFeeder.ErrorCode == 0 {

			response.Event = message.Event
			response.Message = string(util.ToJson(v))

			status_mhs, err := get_status_akm(tokenFeeder.Data["token"].(string), v.NPM, paramAKM)
			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			id_registrasi_mahasiwa, err := get_id_registrasi(tokenFeeder.Data["token"].(string), v.NPM, paramAKM)
			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			ips, _ := strconv.ParseFloat(v.Ips, 64)
			ipk, _ := strconv.ParseFloat(v.Ipk, 64)
			sks, _ := strconv.ParseFloat(v.Sks, 64)
			total_sks, _ := strconv.ParseFloat(v.TotalSks, 64)
			biaya, _ := strconv.ParseInt(v.Biaya, 10, 32)

			arg := model.FeederParams{
				Token: tokenFeeder.Data["token"].(string),
				Data: map[string]interface{}{
					"key": map[string]interface{}{
						"id_registrasi_mahasiswa": id_registrasi_mahasiwa,
						"id_semester":             paramAKM.Semester,
					},
					"record": map[string]interface{}{
						"id_status_mahasiswa": status_mhs,
						"ips":                 ips,
						"ipk":                 ipk,
						"sks_semester":        sks,
						"total_sks":           total_sks,
						"biaya_kuliah_smt":    biaya,
					},
				},
			}

			updateAkm, err := akm.NewSyncAKM().UpdateAKM(arg)

			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			if updateAkm.ErrorCode != 0 {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-danger " style="font-size:0.8rem !important">Gagal</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}

				response.Message = string(util.ToJson(data))
			} else {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-success " style="font-size:0.8rem !important">Behasil</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}
				response.Message = string(util.ToJson(data))

				// only in feeder live not in sandbox
				// var p repository.UpdateSyncParams
				// p.ID, _ = strconv.Atoi(v.ID)
				// err = repository.NewAKMRepository().UpdateHasSync(p)

				// if err != nil {
				// 	log.Println("Error update status sync akm ", err.Error())
				// }
			}
		} else {
			response.Event = "error"
			response.Message = "Gagal Mendapatkan Token Feeder"
		}

		if ep.Kws.IsAlive() {
			// Emit the message directly to specified user
			err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			ep.Kws.Close()
		}

		// time.Sleep(1 * time.Second)
	}
}

func SyncAkmNA(ep *ikisocket.EventPayload, message model.MessageObject) {
	var response model.MessageResponse

	var paramAKM ParamsAKM
	err := json.Unmarshal([]byte(message.Data), &paramAKM)
	if err != nil {
		fmt.Println("Error marshalling json :", err)
		return
	}

	response = akm.ListAKMNAService(paramAKM.KdProdi, paramAKM.Semester)

	type ListAkmJson struct {
		List []struct {
			model.MhsNA
		} `json:"list"`
	}

	var data ListAkmJson
	_ = json.Unmarshal([]byte(response.Message), &data)

	total := []byte(`{"total" : %v}`)
	response.Event = "total"
	response.Message = fmt.Sprintf(string(total), len(data.List))

	if ep.Kws.IsAlive() {
		err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		ep.Kws.Close()
	}

	for i, v := range data.List {
		tokenFeeder := feeder.GetToken()

		if tokenFeeder.ErrorCode == 0 {

			var last_akm_param repository.LastAkmParams
			last_akm_param.Npm = v.NPM

			last_akm, err := repository.NewAKMRepository().LastAKM(last_akm_param)

			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			id_registrasi_mahasiwa, err := get_id_registrasi(tokenFeeder.Data["token"].(string), v.NPM, paramAKM)
			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			ips, _ := strconv.ParseFloat("0", 64)
			ipk, _ := strconv.ParseFloat(last_akm.Ipk, 64)
			sks, _ := strconv.ParseFloat("0", 64)
			total_sks, _ := strconv.ParseFloat(last_akm.TotalSks, 64)
			biaya, _ := strconv.ParseInt(last_akm.Biaya, 10, 32)

			arg := model.FeederParams{
				Token: tokenFeeder.Data["token"].(string),
				Data: map[string]interface{}{
					"key": map[string]interface{}{
						"id_registrasi_mahasiswa": id_registrasi_mahasiwa,
						"id_semester":             paramAKM.Semester,
					},
					"record": map[string]interface{}{
						"id_status_mahasiswa": "N",
						"ips":                 ips,
						"ipk":                 ipk,
						"sks_semester":        sks,
						"total_sks":           total_sks,
						"biaya_kuliah_smt":    biaya,
					},
				},
			}

			updateAkm, err := akm.NewSyncAKM().UpdateAKM(arg)

			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			if updateAkm.ErrorCode != 0 {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-danger " style="font-size:0.8rem !important">Gagal</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}

				response.Message = string(util.ToJson(data))
			} else {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-success " style="font-size:0.8rem !important">Behasil</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}
				response.Message = string(util.ToJson(data))

				// only in feeder live not in sandbox
				// var p repository.UpdateSyncParams
				// p.ID, _ = strconv.Atoi(v.ID)
				// err = repository.NewAKMRepository().UpdateHasSync(p)

				// if err != nil {
				// 	log.Println("Error update status sync akm ", err.Error())
				// }
			}

			// filter := `nim ~* '` + akm.NPM + `'`
			// action.SyncAKMService(tokenFeeder.Data["token"].(string), filter)
		} else {
			response.Event = "error"
			response.Message = "Gagal Mendapatkan Token Feeder"
		}

		if ep.Kws.IsAlive() {
			// Emit the message directly to specified user
			err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			ep.Kws.Close()
		}

		// time.Sleep(1 * time.Second)
	}
}

func SyncAkmDO(ep *ikisocket.EventPayload, message model.MessageObject) {
	var response model.MessageResponse

	var paramAKM ParamsAKM
	err := json.Unmarshal([]byte(message.Data), &paramAKM)
	if err != nil {
		fmt.Println("Error marshalling json :", err)
		return
	}

	response = akm.ListAKMDOService(paramAKM.KdProdi, paramAKM.Semester)

	type ListAkmJson struct {
		List []struct {
			model.MhsDO
		} `json:"list"`
	}

	var data ListAkmJson
	_ = json.Unmarshal([]byte(response.Message), &data)

	total := []byte(`{"total" : %v}`)
	response.Event = "total"
	response.Message = fmt.Sprintf(string(total), len(data.List))

	if ep.Kws.IsAlive() {
		err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		ep.Kws.Close()
	}

	for i, v := range data.List {
		tokenFeeder := feeder.GetToken()

		if tokenFeeder.ErrorCode == 0 {

			var last_akm_param repository.LastAkmParams
			last_akm_param.Npm = v.NPM

			last_akm, err := repository.NewAKMRepository().LastAKM(last_akm_param)

			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			id_registrasi_mahasiwa, err := get_id_registrasi(tokenFeeder.Data["token"].(string), v.NPM, paramAKM)
			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			ips, _ := strconv.ParseFloat("0", 64)
			ipk, _ := strconv.ParseFloat(last_akm.Ipk, 64)
			sks, _ := strconv.ParseFloat("0", 64)
			total_sks, _ := strconv.ParseFloat(last_akm.TotalSks, 64)
			biaya, _ := strconv.ParseInt(last_akm.Biaya, 10, 32)

			arg := model.FeederParams{
				Token: tokenFeeder.Data["token"].(string),
				Data: map[string]interface{}{
					"key": map[string]interface{}{
						"id_registrasi_mahasiswa": id_registrasi_mahasiwa,
						"id_semester":             paramAKM.Semester,
					},
					"record": map[string]interface{}{
						"id_status_mahasiswa": "N",
						"ips":                 ips,
						"ipk":                 ipk,
						"sks_semester":        sks,
						"total_sks":           total_sks,
						"biaya_kuliah_smt":    biaya,
					},
				},
			}

			updateAkm, err := akm.NewSyncAKM().UpdateAKM(arg)

			if err != nil {
				response.Event = "error"
				response.Message = err.Error()
			}

			if updateAkm.ErrorCode != 0 {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-danger " style="font-size:0.8rem !important">Gagal</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}

				response.Message = string(util.ToJson(data))
			} else {
				response.Event = message.Event

				updateAkm.Data["name"] = v.Name
				updateAkm.Data["npm"] = v.NPM
				updateAkm.Data["status"] = `<span class="badge rounded-pill bg-success " style="font-size:0.8rem !important">Behasil</span>`
				updateAkm.Data["order"] = i + 1

				data := fiber.Map{
					"error_code": updateAkm.ErrorCode,
					"error_desc": updateAkm.ErrorDesc,
					"list":       updateAkm.Data,
				}
				response.Message = string(util.ToJson(data))

				// only in feeder live not in sandbox
				// var p repository.UpdateSyncParams
				// p.ID, _ = strconv.Atoi(v.ID)
				// err = repository.NewAKMRepository().UpdateHasSync(p)

				// if err != nil {
				// 	log.Println("Error update status sync akm ", err.Error())
				// }
			}

			// filter := `nim ~* '` + akm.NPM + `'`
			// action.SyncAKMService(tokenFeeder.Data["token"].(string), filter)
		} else {
			response.Event = "error"
			response.Message = "Gagal Mendapatkan Token Feeder"
		}

		if ep.Kws.IsAlive() {
			// Emit the message directly to specified user
			err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			ep.Kws.Close()
		}

		// time.Sleep(1 * time.Second)
	}
}

type ParamsNilai struct {
	KdJadwal     string `json:"kd_jadwal"`
	KdProdi      string `json:"kd_prodi"`
	KdMatakuliah string `json:"kd_matakuliah"`
	Kelas        string `json:"kelas"`
	Semester     string `json:"semester"`
}

func SyncUpdateNilai(ep *ikisocket.EventPayload, message model.MessageObject) {
	var response model.MessageResponse

	var paramNilai ParamsNilai
	err := json.Unmarshal([]byte(message.Data), &paramNilai)
	if err != nil {
		fmt.Println("Error marshalling json :", err)
		return
	}

	response = nilai.ListTransaksiNilai(paramNilai.KdJadwal)

	type ListNilaiJson struct {
		List []struct {
			model.SyncNilai
		} `json:"list"`
	}

	var data ListNilaiJson
	_ = json.Unmarshal([]byte(response.Message), &data)

	total := []byte(`{"total" : %v}`)
	response.Event = "total"
	response.Message = fmt.Sprintf(string(total), len(data.List))

	if ep.Kws.IsAlive() {
		err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		ep.Kws.Close()
	}

	type ListJadwalJson struct {
		List []struct {
			model.Jadwal
		} `json:"list"`
	}
	jadwal := nilai.ListJadwal(paramNilai.KdJadwal)

	var dataJadwal ListJadwalJson
	_ = json.Unmarshal([]byte(jadwal.Message), &dataJadwal)

	paramNilai.KdMatakuliah = dataJadwal.List[0].KdMatakuliah
	paramNilai.Kelas = dataJadwal.List[0].Kelas
	paramNilai.Semester = dataJadwal.List[0].TahunAjaran

	tokenFeeder := feeder.GetToken()

	if tokenFeeder.ErrorCode != 0 {
		response.Event = "error"
		response.Message = "Gagal Mendapatkan Token Feeder"
	}

	id_kelas_kuliah, err := get_id_kelas(tokenFeeder.Data["token"].(string), paramNilai)

	if err != nil {
		response.Event = "error"
		response.Message = "Gagal Mendapatkan Id Kelas Kuliah"
	}

	for i, v := range data.List {

		var paramAKM ParamsAKM
		paramAKM.Semester = paramNilai.Semester

		id_registrasi_mahasiwa, err := get_id_registrasi(tokenFeeder.Data["token"].(string), v.NPM, paramAKM)
		if err != nil {
			response.Event = "error"
			response.Message = err.Error()
		}

		arg := model.FeederParams{
			Token: tokenFeeder.Data["token"].(string),
			Data: map[string]interface{}{
				"key": map[string]interface{}{
					"id_registrasi_mahasiswa": id_registrasi_mahasiwa,
					"id_kelas_kuliah":         id_kelas_kuliah,
				},
				"record": map[string]interface{}{
					"nilai_angka":  fmt.Sprintf("%v", v.NilaiAkhir),
					"nilai_huruf":  v.NilaiIndeks,
					"nilai_indeks": fmt.Sprintf("%f", v.Bobot),
				},
			},
		}

		updateNilai, err := nilai.NewSyncNilai().UpdateNilai(arg)

		if err != nil {
			response.Event = "error"
			response.Message = err.Error()
		}

		updateNilai.Data["name"] = v.Name
		updateNilai.Data["npm"] = v.NPM
		updateNilai.Data["order"] = i + 1
		updateNilai.Data["kd_matakuliah"] = v.KdMatakuliah
		updateNilai.Data["nilai_indeks"] = v.NilaiIndeks
		updateNilai.Data["nilai_akhir"] = v.NilaiAkhir

		if updateNilai.ErrorCode != 0 {
			response.Event = message.Event

			updateNilai.Data["status"] = `<span class="badge rounded-pill bg-danger " style="font-size:0.8rem !important">Gagal</span>`

			data := fiber.Map{
				"error_code": updateNilai.ErrorCode,
				"error_desc": updateNilai.ErrorDesc,
				"list":       updateNilai.Data,
			}

			response.Message = string(util.ToJson(data))
		} else {
			response.Event = message.Event

			updateNilai.Data["status"] = `<span class="badge rounded-pill bg-success " style="font-size:0.8rem !important">Behasil</span>`

			data := fiber.Map{
				"error_code": updateNilai.ErrorCode,
				"error_desc": updateNilai.ErrorDesc,
				"list":       updateNilai.Data,
			}
			response.Message = string(util.ToJson(data))
		}

		if ep.Kws.IsAlive() {
			// Emit the message directly to specified user
			err = ep.Kws.EmitTo(clients[message.To], util.ToJson(response), ikisocket.TextMessage)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			ep.Kws.Close()
		}

		// time.Sleep(1 * time.Second)
	}
}

func get_status_akm(token string, npm string, param ParamsAKM) (string, error) {

	var arg model.FeederParams
	arg.Token = token
	arg.Sms = repository.NewSmsProdiRepository().SMSProdi(repository.SmsParams{KdProdi: param.KdProdi})
	arg.Data = map[string]interface{}{
		"filter": fmt.Sprintf("nim ~* '%s' and id_semester = '%s'", npm, param.Semester),
	}
	fakm := akm.NewSyncAKM()

	feederAkm, err := fakm.StatusAKM(arg)
	if err != nil {
		return "", err
	}

	return feederAkm.IDStatusMahasiswa, nil

}

func get_id_registrasi(token string, npm string, param ParamsAKM) (string, error) {

	var arg model.FeederParams
	arg.Token = token
	arg.Sms = repository.NewSmsProdiRepository().SMSProdi(repository.SmsParams{KdProdi: param.KdProdi})
	arg.Data = map[string]interface{}{
		"filter": fmt.Sprintf("nim ~* '%s' ", npm),
		"limit":  "1",
		"order":  "id_periode_masuk DESC",
	}

	m := mhs.NewRiwayatPendidikan()

	r, err := m.List(arg)
	if err != nil {
		return "", err
	}

	return r.IDRegistrasiMahasiswa, nil

}

func get_id_kelas(token string, param ParamsNilai) (string, error) {

	var arg model.FeederParams
	arg.Token = token
	arg.Sms = repository.NewSmsProdiRepository().SMSProdi(repository.SmsParams{KdProdi: param.KdProdi})

	var filter string

	filter = " id_prodi ='" + arg.Sms + "'"
	filter += " AND kode_mata_kuliah = '" + strings.ReplaceAll(strings.TrimSpace(param.KdMatakuliah), "-", "") + "'"
	filter += " AND nama_kelas_kuliah = '" + strings.ReplaceAll(strings.TrimSpace(param.Kelas), "-", "") + "'"
	filter += " AND id_semester = '" + param.Semester + "'"

	arg.Data = map[string]interface{}{
		"filter": filter,
		"limit":  "1",
		"order":  "",
	}

	k := kelas.NewListKelas()

	r, err := k.List(arg)
	if err != nil {
		return "", err
	}

	return r.IDKelasKuliah, nil

}
