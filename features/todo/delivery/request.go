package delivery

import (
	"skyshi/features/todo"
)

type TodoRequest struct {
	Activity_Group_Id uint   `json:"activity_group_id" form:"activity_group_id"`
	Title             string `json:"title" form:"title"`
	Is_Active         bool   `json:"is_active" form:"is_active"`
	Priority          string `json:"priority" form:"priority"`
}

func ToCore(data TodoRequest) todo.TodoCore {
	return todo.TodoCore{
		Activity_Group_Id: data.Activity_Group_Id,
		Title:             data.Title,
		Is_Active:         data.Is_Active,
		Priority:          data.Priority,
	}
}
