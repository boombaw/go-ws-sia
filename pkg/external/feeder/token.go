package feeder

import (
	"context"
	"errors"
	"os"
	"time"

	rdb "github.com/boombaw/go-ws-sia/pkg/database/redis"
	"github.com/boombaw/go-ws-sia/pkg/util"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var ctx = context.Background()

func GetToken() FeederResponse {
	var response FeederResponse

	response.Data = make(map[string]interface{})

	tokenRedis, err := tokenRedis()

	if err != nil || tokenRedis == "" {
		tokenFeeder, err := tokenFeeder()

		if err != nil {
			response.ErrorCode = 1
			response.ErrorDesc = "Gagal Mendapatkan Token"
			return response
		}

		response.ErrorCode = 0
		response.ErrorDesc = "Berhasil Mendapatkan Token"
		response.Data["token"] = tokenFeeder
		SetRedis("token", tokenFeeder)

		return response
	}

	response.ErrorCode = 0
	response.ErrorDesc = "Berhasil Mendapatkan Token"
	response.Data["token"] = tokenRedis

	return response
}

func tokenRedis() (string, error) {

	redisClient, err := rdb.RedisConn()

	if err != nil {
		return "", err
	}

	token, err := redisClient.Get(ctx, "token").Result()
	if err == redis.Nil {

		tokenFeeder, err := tokenFeeder()

		if err != nil {
			return "", errors.New("gagal Mendapatkan Token")
		}

		SetRedis("token", tokenFeeder)

		return tokenFeeder, nil

	} else if err != nil {
		return "", errors.New("gagal Mendapatkan Token")
	}

	return token, nil
}

func tokenFeeder() (string, error) {

	var response FeederResponse

	url := os.Getenv("FEEDER_URL")

	payload := Token{
		Act:      GET_TOKEN,
		Username: os.Getenv("FEEDER_USERNAME"),
		Password: os.Getenv("FEEDER_PASSWORD"),
	}

	jsonPayload := string(util.ToJson(payload))

	_, body, err := gorequest.New().Post(url).Send(jsonPayload).End()

	_ = json.Unmarshal([]byte(body), &response)

	if err != nil {
		return "", errors.New("gagal Mendapatkan Token")

	}

	if response.ErrorCode != 0 {
		return "", errors.New(response.ErrorDesc)
	}

	SetRedis("token", response.Data["token"].(string))

	return response.Data["token"].(string), nil
}

func SetRedis(key, value string) {

	redisClient, err := rdb.RedisConn()

	if err != nil {
		return
	}

	redisClient.SetEX(ctx, "key", "value", 120*time.Second)

	err = redisClient.SetEX(ctx, key, value, 120*time.Second).Err()
	if err != nil {
		panic(err.Error())
	}
}
