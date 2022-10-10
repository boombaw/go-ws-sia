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

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	err2 := json.Unmarshal([]byte(body), &feederResponse)

	if err2 != nil {
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", err)
	}

	if err != nil {
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", err)
	}

	if feederResponse.ErrorCode != 0 {
		return model.FeederListKelas{}, fmt.Errorf("error GET LIST KELAS : %v", feederResponse.ErrorDesc)
	}

	for _, v := range feederResponse.Data {
		b, _ := json.Marshal(v)
		_ = json.Unmarshal(b, &list_kelas)
	}

	return list_kelas, nil
}
