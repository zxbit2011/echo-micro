package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"github.com/zxbit2011/echo-micro/website/api/router"
	"time"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.echo.website"),
		web.Version("v1.1"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Address(":8080"),
	)

	service.Handle("/", router.Router())

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
