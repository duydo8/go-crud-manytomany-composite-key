package model

import (
	"time"
)

type Tag struct {
	ID          int
	TagName     string       `gorm:"column:tag_name;unique"`
	IsActive    bool         `gorm:"column:is_active;default:true"`
	IsDeleted   bool         `gorm:"column:is_deleted;default:false"`
	TagProducts []TagProduct `gorm:"foreignKey:tag_id"`
}
type Product struct {
	ID          int
	ProductName string    `gorm:"column:product_name"`
	CreateDate  time.Time `gorm:"column:create_date"`
	Description string
	Image       string
	VoteAverage float64
	IsDeleted   bool
	IsActive    bool
	TagProducts []TagProduct `gorm:"foreignKey:product_id"`
}
type TagProduct struct {
	TagId      int `gorm:"column:tag_id;PrimaryKey;autoIncrement:false"`
	ProductId  int `gorm:"column:product_id;PrimaryKey;autoIncrement:false"`
	Tag        Tag
	Product    Product
	CreateTime time.Time
}
