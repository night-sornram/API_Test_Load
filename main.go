package main

import (
	"apitest/handle/gRPC"
	"apitest/handle/http1-1"
	"apitest/handle/http2"
	"github.com/gofiber/fiber/v2"
)

var Counter int

func main() {
	app := fiber.New()

	app.Get("/http1-1/default/phone/:id", http1_1.GetDefaultPhone)
	app.Get("/http1-1/fasthttp/phone/:id", http1_1.GetFastHTTPPhone)
	app.Get("/http1-1/roundrobin/phone/:id", http1_1.GetRoundRobinPhone)

	app.Get("/http2/default/phone/:id", http2.GetDefaultPhone)
	app.Get("/http2/fasthttp/phone/:id", http2.GetFastHTTPPhone)
	app.Get("/http2/roundrobin/phone/:id", http2.GetRoundRobinPhone)

	app.Get("/grpc/default/phone/:id", gRPC.GetDefaultPhone)
	app.Get("/grpc/roundrobin/phone/:id", gRPC.GetRoundRobinPhone)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
