package initializers

import "go-crud-manytomany-composite-key/model"

func SyncDatabase() {
	DB.AutoMigrate(&model.Tag{}, &model.TagProduct{},
		&model.Product{})
}
