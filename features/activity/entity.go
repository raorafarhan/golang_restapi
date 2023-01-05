package activity

import "time"

type ActivityCore struct {
	ID         uint
	Title      string
	Email      string
	Created_At time.Time
	Updated_At time.Time
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
