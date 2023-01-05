package delivery

import (
	"skyshi/features/activity"
)

type ActivityResponse struct {
	ID    uint   `json:"id" form:"id"`
	Title string `json:"tittle" form:"tittle"`
	Email string `json:"email" form:"email"`
}

func FromCore(data activity.ActivityCore) ActivityResponse {
	return ActivityResponse{
		ID:    data.ID,
		Title: data.Title,
		Email: data.Email,
	}

}

func FromCoreList(data []activity.ActivityCore) []ActivityResponse {
	var list []ActivityResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
