package gRPC

type Response struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}

var counter = 0

var serverPool = []string{
	"mock-lookup:8031",
	"mock-lookup:8032",
	"mock-lookup:8033",
}
