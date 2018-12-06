package global

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"github.com/zxbit2011/echo-micro/website/web/router"
	"time"
)

var Service web.Service

func Start() {
	// create new web service
	Service = web.NewService(
		web.Name("go.micro.web.account"),
		web.Version("latest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Address(":8081"),
	)
	Service.Handle("/", router.Router())

	// initialise service
	if err := Service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := Service.Run(); err != nil {
		log.Fatal(err)
	}
}
