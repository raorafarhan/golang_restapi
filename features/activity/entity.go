package activity

import "time"

type ActivityCore struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type UsecaseInterface interface {
	GetAllActivity() (data []ActivityCore, err error)
	GetOneActivity(id int) (data ActivityCore, err error)
	CreateActivity(data ActivityCore) (row int, err error)
	UpdateActivity(data ActivityCore) (row int, err error)
	DeleteActivity(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllAcivity() (data []ActivityCore, err error)
	SelectOneActivity(id int) (data ActivityCore, err error)
	PostActivity(data ActivityCore) (row int, err error)
	PatchActivity(data ActivityCore) (row int, err error)
	DeleteActivity(id int) (row int, err error)
}
