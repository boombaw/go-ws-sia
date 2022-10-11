package action

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boombaw/go-ws-sia/pkg/external/feeder"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type listKelas struct{}

type ListKelas interface {
	List(arg model.FeederParams) (model.FeederListKelas, error)
}

// NewListKelas
func NewListKelas() *listKelas {
	return &listKelas{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (r *listKelas) List(arg model.FeederParams) (model.FeederListKelas, error) {

	var feederResponse feeder.FeederResponses
	var list_kelas model.FeederListKelas

	url := os.Getenv("FEEDER_URL")
	payload := feeder.GetPayload{
		Act:    feeder.GET_LIST_KELAS,
		Token:  arg.Token,
		Filter: arg.Data["filter"].(string),
		Limit:  arg.Data["limit"].(string),
		Order:  arg.Data["order"].(string),
	}

	jsonPayload := string(util.ToJson(payload))

	resp, body, err := gorequest.New().Timeout(5 * time.Minute).Post(url).Send(jsonPayload).End()

	if resp.StatusCode != 200 {
		return model.FeederListKelas{}, fmt.Errorf("error Koneksi Ke Feeder HTTP Header %v", resp.StatusCode)
	}

	err2 := json.Unmarshal([]byte(body), &feederResponse)

	if err2 != nil {
		log.Println("Error Kelas 2", err2)
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", err)
	}

	if err != nil {
		log.Println("Error Kelas 1", err)
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		log.Println("Error Feeder Response", feederResponse.ErrorDesc)
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", feederResponse.ErrorDesc)
	}

	for _, v := range feederResponse.Data {
		b, err := json.Marshal(v)

		if err != nil {
			log.Println("Error marshalling json")
		}

		err = json.Unmarshal(b, &list_kelas)
		if err != nil {
			log.Println("Error unmarshalling json")
		}
	}

	return list_kelas, nil
}
