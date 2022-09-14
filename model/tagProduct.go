package model

import "time"

type TagProduct struct {
	TagId      int `gorm:"column:tag_id;PrimaryKey;autoIncrement:false"`
	ProductId  int `gorm:"column:product_id;PrimaryKey;autoIncrement:false"`
	Tag        Tag
	Product    Product
	CreateTime time.Time
}
type TagProductCreate struct {
	TagId     int
	ProductId int
}
type TagProductTest struct {
	ID         int `gorm:"primaryKey`
	TagId      int `gorm:"column:tag_id"`
	ProductId  int `gorm:"column:product_id"`
	CreateTime time.Time
}
