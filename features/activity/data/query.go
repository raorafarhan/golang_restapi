package data

import (
	"skyshi/features/activity"

	"gorm.io/gorm"
)

type dataActivity struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.DataInterface {
	return &dataActivity{
		db,
	}
}

func (repo *dataActivity) PostActivity(data activity.ActivityCore) (row int, err error) {
	var activity Activity
	activity.Tittle = data.Title
	activity.Email = data.Email

	tx := repo.db.Create(&activity)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataActivity) SelectAllAcivity() ([]activity.ActivityCore, error) {
	var activites []Activity
	tx := repo.db.Find(&activites)
	if tx.Error != nil {
		return nil, tx.Error
	}
	activityCore := toCoreList(activites)
	return activityCore, nil
}

func (repo *dataActivity) SelectOneActivity(id int) (activity.ActivityCore, error) {
	var activityList Activity
	activityList.ID = uint(id)
	tx := repo.db.Where("id = ?", id).First(&activityList)
	if tx.Error != nil {
		return activity.ActivityCore{}, tx.Error
	}
	activityData := activityList.toCore()
	return activityData, nil
}

func (repo *dataActivity) DeleteActivity(id int) (row int, err error) {
	var activity Activity
	activity.ID = uint(id)
	tx := repo.db.Unscoped().Delete(&activity)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataActivity) PatchActivity(data activity.ActivityCore) (int, error) {
	var activityUpdate Activity
	txDataOld := repo.db.First(&activityUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Title != "" {
		activityUpdate.Tittle = data.Title
	}

	tx := repo.db.Save(&activityUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
