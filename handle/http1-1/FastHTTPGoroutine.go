package http1_1

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net"
	"os"
	"strings"
	"time"
)

func GetFastHTTPPGoroutinePhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	url := fmt.Sprintf("http://localhost:8081/phone?number=%s", id)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	client := &fasthttp.Client{
		MaxConnsPerHost: 2000,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*5)
		},
	}

	// Create a channel to communicate the result of the goroutine
	resultCh := make(chan string)
	errorCh := make(chan error)

	go func() {
		err = client.Do(req, resp)
		if err != nil {
			errorCh <- err
			return
		}

		responseData := resp.Body()
		responseDataStr := strings.ReplaceAll(string(responseData), "\r", "")
		resultCh <- responseDataStr
	}()

	select {
	case responseDataStr := <-resultCh:
		content := Response{}
		err = json.Unmarshal([]byte(responseDataStr), &content)
		if err != nil {
			return err
		}
		return c.JSON(content)
	case err := <-errorCh:
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return nil
}
