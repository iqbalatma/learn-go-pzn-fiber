package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
