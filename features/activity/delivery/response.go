package delivery

import (
	"skyshi/features/activity"
	"time"
)

type ActivityResponse struct {
	ID         uint      `json:"id" form:"id"`
	Title      string    `json:"title" form:"title"`
	Email      string    `json:"email" form:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func FromCore(data activity.ActivityCore) ActivityResponse {
	return ActivityResponse{
		ID:         data.ID,
		Title:      data.Title,
		Email:      data.Email,
		Created_At: data.Created_At,
		Updated_At: data.Updated_At,
	}

}

func FromCoreList(data []activity.ActivityCore) []ActivityResponse {
	var list []ActivityResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
