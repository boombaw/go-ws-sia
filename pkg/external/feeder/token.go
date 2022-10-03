package feeder

import (
	"os"

	"github.com/boombaw/go-ws-sia/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetToken() FeederResponse {
	var response FeederResponse

	url := os.Getenv("FEEDER_URL")

	payload := Token{
		Act:      GET_TOKEN,
		Username: "031036",
		Password: "Ubhara_j4y4@lldikti3",
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &response)

	if err != nil {
		return FeederResponse{
			ErrorCode: 500,
			ErrorDesc: "Gagal Mendapatkan Token",
		}
	}

	if response.ErrorCode != 0 {
		return FeederResponse{
			ErrorCode: response.ErrorCode,
			ErrorDesc: response.ErrorDesc,
		}
	}

	return response
}
