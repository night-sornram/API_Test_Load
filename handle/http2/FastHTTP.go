package http2

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"os"
	"strings"
)

func GetFastHTTPPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("http://localhost:8082/phone?number=%s", id)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	client := &fasthttp.Client{
		MaxConnsPerHost: 2000,
	}

	err = client.Do(req, resp)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData := resp.Body()

	responseDataStr := strings.ReplaceAll(string(responseData), "\r", "")

	content := Response{}

	err = json.Unmarshal([]byte(responseDataStr), &content)
	if err != nil {
		return err
	}

	return c.JSON(content)
}
