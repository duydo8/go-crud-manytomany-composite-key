package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go-crud-manytomany-composite-key/controller"
	"go-crud-manytomany-composite-key/docs"
	"go-crud-manytomany-composite-key/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
