package action

import (
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/repository"
	"github.com/boombaw/go-ws-sia/pkg/util"
	"github.com/gofiber/fiber/v2"
)

// ListTransaksiNilai is the responsible for listing all the nilai published
func ListTransaksiNilai(kdJadwal string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewNilaiRepository()

	arg := repository.SyncNilaiParams{
		KdJadwal: kdJadwal,
	}

	nilai, err := repo.ListTransaksiNilai(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": nilai,
	}

	resp.Event = "list_nilai"
	resp.Message = string(util.ToJson(data))

	return resp
}

func ListJadwal(kdJadwal string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewNilaiRepository()

	arg := repository.SyncNilaiParams{
		KdJadwal: kdJadwal,
	}

	jadwal, err := repo.ListJadwal(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": jadwal,
	}

	resp.Event = "list_jadwal"
	resp.Message = string(util.ToJson(data))

	return resp
}
