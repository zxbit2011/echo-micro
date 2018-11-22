package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Info(c echo.Context) error {
	return c.HTML(http.StatusOK, fmt.Sprintf(`<h1>Info：%s</h1>`, c.FormValue("name")))
}
