package delivery

import (
	"skyshi/features/activity"
)

type ActivityRequest struct {
	Title string `json:"title" form:"title"`
	Email string `json:"email" form:"email"`
}

func ToCore(data ActivityRequest) activity.ActivityCore {
	return activity.ActivityCore{
		Title: data.Title,
		Email: data.Email,
	}
}
