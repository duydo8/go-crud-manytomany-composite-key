package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-manytomany-composite-key/config"
	"go-crud-manytomany-composite-key/initializers"
	"go-crud-manytomany-composite-key/model"
	"strconv"
	"time"
)

func FindAllProduct(context *gin.Context) {
	var product []model.Product
	if err := initializers.DB.Find(&product).Error; err != nil {
		context.JSON(400, gin.H{
			"err": err,
		})
	} else {
		config.CustomResponse(context, 200, "", product)
	}

}

func CreateProduct(context *gin.Context) {
	var productCreate model.ProductCreate
	context.BindJSON(&productCreate)
	var product model.Product
	product.ProductName = productCreate.ProductName
	product.Description = productCreate.Description
	product.Image = productCreate.Image
	product.VoteAverage = productCreate.VoteAverage
	product.CreateDate = time.Now()
	product.IsDeleted = false
	product.IsActive = true
	if err := initializers.DB.Create(&product).Error; err != nil {
		config.CustomResponse(context, 200, "", err)
	} else {
		config.CustomResponse(context, 200, "", product)
	}

}
func UpdateProduct(context *gin.Context) {
	var product, product1 model.Product
	context.ShouldBindJSON(&product)
	id := product.ID

	if err1 := initializers.DB.Find(&product1, id).Error; err1 != nil {
		config.CustomResponse(context, 400, "not found", "")
	} else {

		initializers.DB.Save(&product)
		config.CustomResponse(context, 200, "success", product)
	}

}
func FindByIdProduct(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		config.CustomResponse(context, 400, "id must be int", "")
	} else {
		var product model.Product
		if err1 := initializers.DB.Find(&product, id).Error; err1 != nil {
			config.CustomResponse(context, 400, "not found", "")
		} else {
			config.CustomResponse(context, 200, "success", product)
		}

	}
}
func DeleteProduct(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		config.CustomResponse(context, 400, "id must be int", "")
	} else {
		var tag model.Tag
		if err1 := initializers.DB.Delete(&tag, id).Error; err1 != nil {
			config.CustomResponse(context, 400, "not found", "")
		} else {
			config.CustomResponse(context, 200, "success", tag)
		}

	}
}
