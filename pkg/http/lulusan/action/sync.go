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

type syncLulusan struct{}

type SyncLulusan interface {
	InsertLulusan(arg model.FeederParams) (feeder.FeederResponse, error)
}

// NewSyncLulusan
func NewSyncLulusan() *syncLulusan {
	return &syncLulusan{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (a *syncLulusan) InsertLulusan(arg model.FeederParams) (feeder.FeederResponse, error) {
	var feederResponse feeder.FeederResponse

	url := os.Getenv("FEEDER_URL")
	payload := feeder.PostPayload{
		Act:   feeder.INSERT_LULUS_DO,
		Token: arg.Token,
		Record: feeder.Record{
			arg.Data["record"].(map[string]interface{}),
		},
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &feederResponse)

	if err != nil {
		return feeder.FeederResponse{}, fmt.Errorf("error Insert Lulusan : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		return feederResponse, fmt.Errorf("error Insert Lulusan : %v", feederResponse.ErrorDesc)
	}

	return feederResponse, nil
}
