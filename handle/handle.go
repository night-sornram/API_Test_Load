package handle

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

type Response struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}

func GetPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("http://localhost:8081/phone?number=%s", id)

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
