package model

import "time"

type Product struct {
	ID             int
	ProductName    string    `gorm:"column:product_name"`
	CreateDate     time.Time `gorm:"column:create_date"`
	Description    string
	Image          string
	VoteAverage    float64
	IsDeleted      bool
	IsActive       bool
	TagProduct     []TagProduct     `gorm:"foreignKey:product_id"`
	TagProductTest []TagProductTest `gorm:"foreignKey:product_id"`
}
type ProductCreate struct {
	ProductName string
	Description string
	Image       string
	VoteAverage float64
}
