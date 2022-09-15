package controller

import (
	"go-crud-manytomany-composite-key/config"
	"go-crud-manytomany-composite-key/initializers"
	"go-crud-manytomany-composite-key/model"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTagProductTest(context *gin.Context) {

	var tagProductTest model.TagProductTest
	var tagProductCreate model.TagProductCreate
	var tag model.Tag
	var product model.Product
	context.ShouldBindJSON(&tagProductCreate)
	tagProductTest.CreateTime = time.Now()
	tagProductTest.TagId = tagProductCreate.TagId
	tagProductTest.ProductId = tagProductCreate.ProductId
	if err := initializers.DB.First(&tag, tagProductCreate.TagId).Error; err != nil {
		config.CustomResponse(context, 400, "", err)
	}
	tagProductTest.Tag = tag

	if err := initializers.DB.First(&product, tagProductCreate.ProductId).Error; err != nil {
		config.CustomResponse(context, 400, "", err)
	}
	tagProductTest.Product = product

	if err := initializers.DB.Create(&tagProductTest).Error; err != nil {
		config.CustomResponse(context, 400, "", err)
	}
	initializers.DB.Save(&tagProductTest)

	var tagProductTag []model.TagProductTest
	var productTag []model.TagProductTest
	tagProductTag = append(tag.TagProductTests, tagProductTest)
	productTag = append(product.TagProductTests, tagProductTest)
	tag.TagProductTests = tagProductTag
	product.TagProductTests = productTag

	initializers.DB.Save(&tag)
	initializers.DB.Save(&product)

	config.CustomResponse(context, 200, "success", tagProductTest)

}
func FindAllTagProductTest(context *gin.Context) {
	var tagProductTest []model.TagProductTest
	if err := initializers.DB.Last(&tagProductTest).Error; err != nil {
		config.CustomResponse(context, 400, "", err)
	} else {
		config.CustomResponse(context, 200, "", tagProductTest)
	}

}
