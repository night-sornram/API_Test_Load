package main

import (
	"apitest/handle"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/phone/:id", handle.GetPhone)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
