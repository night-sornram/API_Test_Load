package gRPC

import (
	"apitest/protos"
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"log"
)

func GetDefaultPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cc := protos.NewLookupServiceClient(conn)

	req := &protos.LookupReq{PhoneNumber: id}
	res, err := cc.Lookup(context.Background(), req)
	if err != nil {
		log.Fatalf("could not lookup: %v", err)
	}

	return c.JSON(res)

}
