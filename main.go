package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro-test/prodService"
)

func main(){


	csReg :=consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
		)

	ginRoute :=gin.Default()

	v1Group := ginRoute.Group("/v1")

	v1Group.Handle("GET","/prod", func(context *gin.Context) {
		context.JSON(200,prodService.NewProdList(5))
	})

	server := web.NewService(
		web.Name("prodService"),
		//web.Address(":8001"),
		web.Handler(ginRoute),
		web.Registry(csReg),
		)

	server.Init()
	// go run main.go  --server_address :8001
	server.Run()
}
