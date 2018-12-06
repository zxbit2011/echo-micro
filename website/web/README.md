# Go ECHO + Micro 微服务 - 官网站点服务
> namespace：go.micro.echo.website - v1.1

### （在`website\web`执行）启动：
```
go run main.go
```
### 一、测试(`grpc调用方式，路径：account\api\proto\account`)：
```
$ curl -H 'Content-Type: application/json' -X GET "http://localhost:8081/proto/account/getName?name=hello_zxbit2011"
```
#### grpc执行结果
[![测试效果](test/1.png)](test/1.png)
### 二、测试ECHO API：
```
$ curl -H 'Content-Type: application/json' -X GET "http://localhost:8081/proto/account/getName?name=hello_zxbit2011"
```
#### grpc执行结果
[![测试效果](test/1.png)](test/2.png)
