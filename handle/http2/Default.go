package http2

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

func GetDefaultPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("http://localhost:8021/phone?number=%s", id)

	client := http.Client{
		Timeout: 6 * time.Second,
	}

	response, err := client.Get(url)

	if err != nil {
		fmt.Println(err.Error())
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
