package controller

import (
	"fmt"
	"go-crud-manytomany-composite-key/config"
	"go-crud-manytomany-composite-key/initializers"
	"go-crud-manytomany-composite-key/model"
	"time"

	"github.com/gin-gonic/gin"
)

func FindAllTagProduct(context *gin.Context) {
	var tagProduct []model.TagProduct
	if err := initializers.DB.Find(&tagProduct).Error; err != nil {
		context.JSON(400, gin.H{
			"err": err,
		})
	} else {
		config.CustomResponse(context, 200, "", tagProduct)
	}

}
func CreateTagProduct(context *gin.Context) {

	var tagProductCreate model.TagProductCreate
	var tag model.Tag
	var product model.Product
	context.BindJSON(&tagProductCreate)
	fmt.Println(tagProductCreate)
	var tagProduct model.TagProduct
	tagProduct.TagId = tagProductCreate.TagId
	tagProduct.ProductId = tagProductCreate.ProductId
	tagProduct.CreateTime = time.Now()
	if err := initializers.DB.First(&tag, tagProductCreate.TagId).Error; err != nil {
		config.CustomResponse(context, 200, "", err)
	}
	tagProduct.Tag = tag

	if err := initializers.DB.First(&product, tagProductCreate.TagId).Error; err != nil {
		config.CustomResponse(context, 200, "", err)
	}
	tagProduct.Product = product

	if err := initializers.DB.Create(&tagProduct).Error; err != nil {
		config.CustomResponse(context, 200, "", err)
	} else {
		config.CustomResponse(context, 200, "", tagProduct)
	}

}
