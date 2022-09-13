package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-manytomany-composite-key/initializers"
	"go-crud-manytomany-composite-key/model"
)

func FindAll(context *gin.Context) {
	var tag []model.Tag
	if err := initializers.DB.Find(&tag).Error; err != nil {
		context.JSON(400, gin.H{
			"err": err,
		})
	} else {
		context.JSON(200, gin.H{
			"data": tag,
		})
	}

}
