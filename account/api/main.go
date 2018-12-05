package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/zxbit2011/echo-micro/account/api/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "go.micro.echo.account Hello " + req.Name
	return nil
}

func main() {
	// 创建新服务。这里可以选择一些选项。
	service := micro.NewService(
		micro.Name("go.micro.echo.account"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
