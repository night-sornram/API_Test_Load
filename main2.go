package main

import (
	"apitest/protos"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

var counter = 1

var serverPool = []string{
	"localhost:8083",
	"localhost:8086",
	"localhost:8089",
}

var connPool []*grpc.ClientConn
var connPoolMutex sync.Mutex

func init() {
	connPool = make([]*grpc.ClientConn, len(serverPool))
}

func main() {
	r := gin.Default()

	gRPCGroup := r.Group("/grpc")
	gRPCGroup.GET("/roundrobin/phone/:id", getRoundRobinPhone)

	err := r.Run(":4000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func getRoundRobinPhone(c *gin.Context) {
	id := c.Param("id")

	connPoolMutex.Lock()
	conn := connPool[counter%3]
	if conn == nil {
		var err error
		conn, err = grpc.Dial(serverPool[counter%3], grpc.WithInsecure())
		if err != nil {
			connPoolMutex.Unlock()
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		connPool[counter%3] = conn
	}
	counter++
	connPoolMutex.Unlock()

	cc := protos.NewLookupServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3200*time.Millisecond)
	defer cancel()

	req := &protos.LookupReq{PhoneNumber: id}
	res, err := cc.Lookup(ctx, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
