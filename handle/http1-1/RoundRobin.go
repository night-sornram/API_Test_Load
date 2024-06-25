package http1_1

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
	"http://localhost:8011",
	"http://localhost:8012",
	"http://localhost:8013",
	// Add more servers as needed
}

func GetRoundRobinPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("%s/phone?number=%s", serverPool[counter%3], id)

	counter++

	client := http.Client{
		Timeout: 8 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		if response != nil {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		return c.SendStatus(fiber.StatusRequestTimeout)
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
