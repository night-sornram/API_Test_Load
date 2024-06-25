package gRPC

import (
	"apitest/protos"
	"context"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"log"
)

func GetDefaultPhone(r *rate.Limiter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !r.Allow() {
			return c.SendStatus(fiber.StatusTooManyRequests)
		}

		id := c.Params("id")
		conn, err := grpc.Dial("localhost:8031", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		cc := protos.NewLookupServiceClient(conn)

		req := &protos.LookupReq{PhoneNumber: id}
		res, err := cc.Lookup(context.Background(), req)
		if err != nil {
			//log.Fatalf("could not lookup: %v", err)
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}

		return c.JSON(res)
	}
}
