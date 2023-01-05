package migration

import (
	activityModel "skyshi/features/activity/data"
	todoModel "skyshi/features/todo/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&activityModel.Activity{})
	db.AutoMigrate(&todoModel.Todo{})

}
