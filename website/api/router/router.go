package router

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/echo-micro/website/api/handler"
)

func Router() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.GET("/", handler.Home)
	e.GET("/home", handler.Home)
	e.GET("/info", handler.Info)
	return e
}
