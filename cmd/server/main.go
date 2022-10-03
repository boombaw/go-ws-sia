package main

import (
	r "github.com/boombaw/go-ws-sia/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	app := fiber.New()
	r.Routes(app)
}
