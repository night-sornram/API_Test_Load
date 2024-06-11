package main

import (
	"apitest/handle"
	"github.com/gofiber/fiber/v2"
)

var Counter int

func main() {
	app := fiber.New()

	app.Get("/default/phone/:id", handle.GetDefaultPhone)
	app.Get("/fasthttp/phone/:id", handle.GetFastHTTPPhone)
	app.Get("/roundrobin/phone/:id", handle.GetRoundRobinPhone)
	app.Get("/fasthttpgoroutine/phone/:id", handle.GetFastHTTPPGoroutinePhone)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
