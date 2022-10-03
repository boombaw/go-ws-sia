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

type riwayatPendidikan struct{}

type RiwawyatPendidikan interface {
	List(arg model.FeederParams) (model.FeederRiwayatPendidikan, error)
}

// NewRiwayatPendidikan
func NewRiwayatPendidikan() *riwayatPendidikan {
	return &riwayatPendidikan{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (r *riwayatPendidikan) List(arg model.FeederParams) (model.FeederRiwayatPendidikan, error) {

	var feederResponse feeder.FeederResponses
	var riwayat model.FeederRiwayatPendidikan

	url := os.Getenv("FEEDER_URL")
	payload := feeder.GetPayload{
		Act:    feeder.GET_RIWAYAT_PENDIDIKAN_MHS,
		Token:  arg.Token,
		Filter: arg.Data["filter"].(string),
		Limit:  arg.Data["limit"].(string),
		Order:  arg.Data["order"].(string),
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &feederResponse)

	if err != nil {
		return model.FeederRiwayatPendidikan{}, fmt.Errorf("error GET AKM : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		return model.FeederRiwayatPendidikan{}, fmt.Errorf("error GET AKM : %v", feederResponse.ErrorDesc)
	}

	for _, v := range feederResponse.Data {
		b, _ := json.Marshal(v)
		_ = json.Unmarshal(b, &riwayat)
	}

	return riwayat, nil
}
