package http2

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var counter = 1

var serverPool = []string{
	"http://localhost:8082",
	"http://localhost:8085",

	// Add more servers as needed
}

func GetRoundRobinPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("%s/phone?number=%s", serverPool[counter%2], id)

	counter++

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
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
