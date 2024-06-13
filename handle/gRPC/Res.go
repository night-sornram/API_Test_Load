package gRPC

type Response struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}

var counter = 1

var serverPool = []string{
	"localhost:8083",
	"localhost:8086",
	"localhost:8089",
}
