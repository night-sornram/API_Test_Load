package http1_1

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var counter = 1

var serverPool = []string{
	"http://mock-lookup:8011",
	"http://mock-lookup:8012",
	"http://mock-lookup:8013",
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

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseDataStr := strings.ReplaceAll(string(responseData), "\r", "")

	content := Response{}

	err = json.Unmarshal([]byte(responseDataStr), &content)
	if err != nil {
		return err
	}

	return c.JSON(content)
}
