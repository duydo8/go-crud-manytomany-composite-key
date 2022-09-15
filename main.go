package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			tag.POST("create", controller.Create)
			tag.PUT("update", controller.Update)
			tag.GET("findById", controller.FindById)
			tag.DELETE("delete", controller.Delete)
		}
		product := v1.Group("product")
		{
			product.GET("/", controller.FindAllProduct)
			product.POST("create", controller.CreateProduct)
			product.PUT("update", controller.UpdateProduct)
			product.GET("findById", controller.FindByIdProduct)
			product.DELETE("delete", controller.DeleteProduct)
		}
		tagProduct := v1.Group("tagProduct")
		{
			tagProduct.GET("/", controller.FindAllTagProduct)
			tagProduct.POST("create", controller.CreateTagProduct)
		}
		tagProductTest := v1.Group("tagProductTest")
		{
			tagProductTest.GET("/", controller.FindAllTagProductTest)
			tagProductTest.POST("create", controller.CreateTagProductTest)
		}
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
