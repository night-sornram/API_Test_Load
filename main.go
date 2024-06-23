package main

import (
	"apitest/handle/gRPC"
	"apitest/handle/http1-1"
	"apitest/handle/http2"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	http1_1Group := app.Group("/http1-1")
	http1_1Group.Get("/default/phone/:id", http1_1.GetDefaultPhone)
	http1_1Group.Get("/fasthttp/phone/:id", http1_1.GetFastHTTPPhone)
	http1_1Group.Get("/roundrobin/phone/:id", http1_1.GetRoundRobinPhone)

	http2Group := app.Group("/http2")
	http2Group.Get("/default/phone/:id", http2.GetDefaultPhone)
	http2Group.Get("/fasthttp/phone/:id", http2.GetFastHTTPPhone)
	http2Group.Get("/roundrobin/phone/:id", http2.GetRoundRobinPhone)

	gRPCGroup := app.Group("/grpc")
	gRPCGroup.Get("/default/phone/:id", gRPC.GetDefaultPhone)
	gRPCGroup.Get("/roundrobin/phone/:id", gRPC.GetRoundRobinPhone)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
