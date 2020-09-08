package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main(){


	ginRoute :=gin.Default()



	ginRoute.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"data":"index",
		})
	})

	server := web.NewService(
		web.Address(":8000"),
		web.Handler(ginRoute),

	)

	server.Run()
}
