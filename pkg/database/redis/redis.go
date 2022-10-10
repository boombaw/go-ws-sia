package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisConn() (*redis.Client, error) {
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

	_, err := rdb.Ping(rdb.Context()).Result()

	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func Set(key string, value string, db ...int) error {
	rdb, err := RedisConn()

	if err != nil {
		return err
	}

	var expiration time.Duration

	if len(db) > 0 {
		rdb.Options().DB = db[0]

		if db[1] == 0 {
			expiration = 0
		} else {
			expiration = time.Duration(db[1]) * time.Second
		}

	} else {
		expiration = 10 * time.Minute
	}

	err = rdb.Set(rdb.Context(), key, value, expiration).Err()

	if err != nil {
		return err
	}

	rdb.Close()
	return nil
}
