package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	return c.HTML(http.StatusOK, fmt.Sprintf(`<h1>欢迎来到Echo Web Micro [website - api]的API服务！</h1><p>%s</p>`, c.Request().Form.Encode()))
}
