package router

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/echo-micro/website/web/handler"
)

func Router() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.Static("/", "html")
	e.GET("/", handler.Home)
	e.GET("/home", handler.Home)
	//Api - 模式的微服务
	api := e.Group("/api")
	{
		api.GET("/", handler.ApiHome)
		api.GET("/home", handler.ApiHome)
	}
	//Rpc - 模式的微服务
	proto := e.Group("/proto")
	{

		proto.GET("/account/getName", handler.GetAccountName)
	}
	return e
}
