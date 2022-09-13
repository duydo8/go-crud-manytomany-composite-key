package main

import (
	"github.com/gin-gonic/gin"
	"go-crud-manytomany-composite-key/controller"
	"go-crud-manytomany-composite-key/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		tag := v1.Group("tag")
		{
			tag.GET("/", controller.FindAll)
			tag.POST("create")
		}
		product := v1.Group("product")
		{
			product.GET("")
		}
	}
	r.Run()
}
