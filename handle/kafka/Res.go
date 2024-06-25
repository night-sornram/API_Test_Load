package kafka

type Response struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}
