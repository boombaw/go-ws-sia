package action

import (
	"fmt"
	"os"

	"github.com/boombaw/go-ws-sia/pkg/external/feeder"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type syncNilai struct{}

type SyncNilai interface {
	UpdateNilai(arg model.FeederParams) (feeder.FeederResponse, error)
}

// NewSyncNilai
func NewSyncNilai() *syncNilai {
	return &syncNilai{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (a *syncNilai) UpdateNilai(arg model.FeederParams) (feeder.FeederResponse, error) {
	var feederResponse feeder.FeederResponse

	url := os.Getenv("FEEDER_URL")
	payload := feeder.PutPayload{
		Act:    feeder.UPDATE_NILAI_KELAS,
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
