package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func RedisConn() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")

	redisConnString := fmt.Sprintf("%s:%s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConnString,
		Username: "",
		Password: pass, // no password set
		DB:       0,    // use default DB
	})

	RedisDB = rdb
	return RedisDB
}

func Set(key string, value string, db ...int) error {

	var expiration time.Duration

	if len(db) > 0 {
		RedisDB.Options().DB = db[0]

		if db[1] == 0 {
			expiration = 0
		} else {
			expiration = time.Duration(db[1]) * time.Second
		}

	} else {
		expiration = 10 * time.Minute
	}

	err := RedisDB.Set(RedisDB.Context(), key, value, expiration).Err()

	if err != nil {
		return err
	}

	return nil
}
