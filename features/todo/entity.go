package todo

import "time"

type TodoCore struct {
	ID                uint      `json:"id"`
	Activity_Group_Id uint      `json:"activity_group_id"`
	Title             string    `json:"title"`
	Is_Active         bool      `json:"is_active"`
	Priority          string    `json:"priority"`
	Created_At        time.Time `json:"createdAt"`
	Updated_At        time.Time `json:"updatedAt"`
}

type UsecaseInterface interface {
	GetAllTodo(activity_group_id int) (data []TodoCore, err error)
	GetOneTodo(id int) (data TodoCore, err error)
	CreateTodo(data TodoCore) (id int, data1 TodoCore, err error)
	UpdateTodo(data TodoCore) (row int, err error)
	DeleteTodo(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllTodo(activity_group_id int) (data []TodoCore, err error)
	SelectOneTodo(id int) (data TodoCore, err error)
	PostTodo(data TodoCore) (id int, data1 TodoCore, err error)
	PatchTodo(data TodoCore) (row int, err error)
	DeleteTodo(id int) (row int, err error)
}
