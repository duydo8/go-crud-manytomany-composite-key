package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-manytomany-composite-key/config"
	"go-crud-manytomany-composite-key/initializers"
	"go-crud-manytomany-composite-key/model"
	"strconv"
)

func FindAll(context *gin.Context) {
	var tag []model.Tag
	if err := initializers.DB.Find(&tag).Error; err != nil {
		context.JSON(400, gin.H{
			"err": err,
		})
	} else {
		config.CustomResponse(context, 200, "", tag)
	}

}
func Create(context *gin.Context) {
	var tag model.Tag
	context.BindJSON(&tag)
	if err := initializers.DB.Create(&tag).Error; err != nil {
		config.CustomResponse(context, 200, "", err)
	} else {
		config.CustomResponse(context, 200, "", tag)
	}

}
func Update(context *gin.Context) {
	var tag, tag1 model.Tag
	context.ShouldBindJSON(&tag)
	id := tag.ID

	if err1 := initializers.DB.Find(&tag1, id).Error; err1 != nil {
		config.CustomResponse(context, 400, "not found", "")
	} else {

		initializers.DB.Save(&tag)
		config.CustomResponse(context, 200, "success", tag)
	}

}
func FindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		config.CustomResponse(context, 400, "id must be int", "")
	} else {
		var tag model.Tag
		if err1 := initializers.DB.Find(&tag, id).Error; err1 != nil {
			config.CustomResponse(context, 400, "not found", "")
		} else {
			config.CustomResponse(context, 200, "success", tag)
		}

	}
}
func Delete(context *gin.Context) {
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
