package gRPC

import (
	"apitest/protos"
	"context"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var connPool []*grpc.ClientConn
var connPoolMutex sync.Mutex

func init() {
	connPool = make([]*grpc.ClientConn, len(serverPool))
}

func GetRoundRobinPhone(r *rate.Limiter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !r.Allow() {
			//return c.SendStatus(fiber.StatusTooManyRequests)
		}
		id := c.Params("id")

		connPoolMutex.Lock()
		conn := connPool[counter%3]

		totalTimeout := 8 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
		defer cancel()

		if conn == nil {
			var err error
			conn, err = grpc.DialContext(ctx, serverPool[counter%3], grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				connPoolMutex.Unlock()
				return err
			}
			connPool[counter%3] = conn
		}
		counter++
		connPoolMutex.Unlock()

		cc := protos.NewLookupServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3200*time.Millisecond)
	defer cancel()

		req := &protos.LookupReq{PhoneNumber: id}
		res, err := cc.Lookup(callCtx, req)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}
