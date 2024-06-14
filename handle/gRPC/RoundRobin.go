package gRPC

import (
	"apitest/protos"
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var connPool []*grpc.ClientConn
var connPoolMutex sync.Mutex

func init() {
	connPool = make([]*grpc.ClientConn, len(serverPool))
}

func GetRoundRobinPhone(c *fiber.Ctx) error {
	id := c.Params("id")

	connPoolMutex.Lock()
	conn := connPool[counter%3]
	if conn == nil {
		var err error
		conn, err = grpc.Dial(serverPool[counter%3], grpc.WithInsecure())
		if err != nil {
			connPoolMutex.Unlock()
			return err
		}
		connPool[counter%3] = conn
	}
	counter++
	connPoolMutex.Unlock()

	cc := protos.NewLookupServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3300*time.Millisecond)
	defer cancel()

	req := &protos.LookupReq{PhoneNumber: id}
	res, err := cc.Lookup(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
