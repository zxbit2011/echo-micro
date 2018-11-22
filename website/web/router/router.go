package router

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/echo-micro/website/web/handler"
)

func Router() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.Static("/","html")
	e.GET("/", handler.Home)
	e.GET("/home", handler.Home)
	api:=e.Group("/api")
	{
		api.GET("/", handler.ApiHome)
		api.GET("/home", handler.ApiHome)
	}
	return e
}
