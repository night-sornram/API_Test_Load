package http2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var counter = 1
var client = &http.Client{} // Reuse HTTP client

var serverPool = []string{
	"http://localhost:8082",
	"http://localhost:8085",
	"http://localhost:8088",
	// Add more servers as needed
}

func GetRoundRobinPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("%s/phone?number=%s", serverPool[counter%3], id)

	counter++

	ctx, cancel := context.WithTimeout(context.Background(), 3300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	bodyStr := strings.ReplaceAll(string(body), "\r", "")
	body = []byte(bodyStr)

	content := Response{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		return err
	}

	return c.JSON(content)
}
