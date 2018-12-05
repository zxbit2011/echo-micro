package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/micro/go-micro/client"
	"net/http"
)

func Home(c echo.Context) error {
	return c.HTML(http.StatusOK, fmt.Sprintf(`<h1>欢迎来到Echo Web Micro [website - web]的官网！</h1><p>%s</p><p><a href='/page/index.html'>进入主页</a></p>`,
		c.Request().Form.Encode()))
}

func ApiHome(c echo.Context) error {
	r := c.Request()
	var response json.RawMessage
	req := client.NewRequest("go.micro.echo.website", "home", r)

	err := client.Call(context.TODO(), req, &response)
	if err != nil {
		return c.HTML(http.StatusBadRequest,err.Error())
	}
	return c.HTML(http.StatusOK, string(response))
}
