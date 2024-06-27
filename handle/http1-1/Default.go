package http1_1

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetDefaultPhone(r *rate.Limiter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !r.Allow() {
			//return c.SendStatus(fiber.StatusTooManyRequests)
		}

		id := c.Params("id")

		url := fmt.Sprintf("http://mock-lookup:8011/phone?number=%s", id)

		response, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusRequestTimeout).JSON(err.Error())
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
}
