package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/micro/go-micro/client"
	"github.com/zxbit2011/echo-micro/account/api/proto"
	"net/http"
	"time"
)

func GetAccountName(c echo.Context) error {
	// New RPC client
	rpcClient := client.NewClient(client.RequestTimeout(time.Second * 120))
	// Create new greeter client
	greeter := proto.NewGreeterService("go.micro.echo.account", rpcClient)
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: c.FormValue("name")})
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, rsp.String())
}
