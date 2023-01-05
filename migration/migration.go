package migration

import (
	activityModel "skyshi/features/activity/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&activityModel.Activity{})

}
