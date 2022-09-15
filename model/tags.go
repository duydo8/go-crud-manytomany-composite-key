package model

type Tag struct {
	ID              int
	TagName         string           `gorm:"column:tag_name;unique"`
	IsActive        bool             `gorm:"column:is_active;default:true"`
	IsDeleted       bool             `gorm:"column:is_deleted;default:false"`
	TagProducts     []TagProduct     `json:"-",gorm:"foreignKey:TagId;references:ID"`
	TagProductTests []TagProductTest `json:"-",gorm:"foreignKey:TagId;references:ID"`
}
