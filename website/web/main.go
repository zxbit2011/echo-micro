package main

import (
	"github.com/zxbit2011/echo-micro/website/web/router"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.account"),
		web.Version("latest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Address(":8081"),
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
