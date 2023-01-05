package data

import (
	"skyshi/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
}

func fromCore(data activity.ActivityCore) Activity {
	return Activity{
		Title: data.Title,
		Email: data.Email,
	}
}

func (data *Activity) toCore() activity.ActivityCore {
	return activity.ActivityCore{
		ID:    data.ID,
		Title: data.Title,
		Email: data.Email,
	}
}

func toCoreList(data []Activity) []activity.ActivityCore {
	var list []activity.ActivityCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
