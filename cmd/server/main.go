package main

import (
	rdb "github.com/boombaw/go-ws-sia/pkg/database/redis"
	r "github.com/boombaw/go-ws-sia/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
	rdb.RedisConn()

}

func main() {
	app := fiber.New()
	r.Routes(app)
}
