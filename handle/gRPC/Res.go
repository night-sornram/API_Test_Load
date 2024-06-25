package gRPC

type Response struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}

var counter = 0

var serverPool = []string{
	"localhost:8031",
	"localhost:8032",
	"localhost:8033",
}
