package action

import (
	"fmt"
	"log"
	"os"

	"github.com/boombaw/go-ws-sia/pkg/external/feeder"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type syncAKM struct{}

type SyncAKM interface {
	UpdateAKM(arg model.FeederParams) (feeder.FeederResponse, error)
	StatusAKM(arg model.FeederParams) (model.FeederAKM, error)
}

// NewSyncAKM
func NewSyncAKM() *syncAKM {
	return &syncAKM{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (a *syncAKM) UpdateAKM(arg model.FeederParams) (feeder.FeederResponse, error) {
	var feederResponse feeder.FeederResponse

	url := os.Getenv("FEEDER_URL")
	payload := feeder.PutPayload{
		Act:    feeder.UPDATE_LIST_KULIAH_MHS,
		Token:  arg.Token,
		Key:    arg.Data["key"].(map[string]interface{}),
		Record: arg.Data["record"].(map[string]interface{}),
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &feederResponse)

	if err != nil {
		return feeder.FeederResponse{}, fmt.Errorf("error GET AKM : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		return feederResponse, fmt.Errorf("error GET AKM : %v", feederResponse.ErrorDesc)
	}

	return feederResponse, nil
}

func (a *syncAKM) StatusAKM(arg model.FeederParams) (model.FeederAKM, error) {
	var feederResponse feeder.FeederResponses
	var feederAKm model.FeederAKM

	url := os.Getenv("FEEDER_URL")
	payload := feeder.GetPayload{
		Act:    feeder.GET_AKM,
		Token:  arg.Token,
		Filter: arg.Data["filter"].(string),
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &feederResponse)

	if err != nil {
		return model.FeederAKM{}, fmt.Errorf("error GET AKM : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		return model.FeederAKM{}, fmt.Errorf("error GET AKM : %v", feederResponse.ErrorDesc)
	}

	for _, v := range feederResponse.Data {
		b, _ := json.Marshal(v)
		_ = json.Unmarshal(b, &feederAKm)
	}

	// feederAKm.IDStatusMahasiswa = feederResponse.Data[0]["id_status_mahasiswa"].(string)

	return feederAKm, nil
}

// sync akm gorequest
func SyncAKMService(token string, filter string) {
	url := os.Getenv("FEEDER_URL")

	// log.Printf("%v\n", filter)

	payload := feeder.GetPayload{
		Act:    "GetAktivitasKuliahMahasiswa",
		Token:  token,
		Filter: filter,
	}

	json := string(util.ToJson(payload))

	// res, body, err
	resp, _, err := gorequest.New().Post(url).Send(json).End()

	if err != nil {
		log.Println(err)
	}

	log.Println(resp)
	// log.Println(body)
	// req.Header.Add("Content-Type", "application/json")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

}
