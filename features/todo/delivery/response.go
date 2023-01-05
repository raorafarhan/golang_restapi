package delivery

import (
	"skyshi/features/todo"
)

type TodoResponse struct {
	ID                uint   `json:"id"`
	Activity_Group_Id uint   `json:"activity_group_id"`
	Title             string `json:"email"`
	Is_Active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
}

func FromCore(data todo.TodoCore) TodoResponse {
	return TodoResponse{
		ID:                data.ID,
		Activity_Group_Id: data.Activity_Group_Id,
		Title:             data.Title,
		Is_Active:         data.Is_Active,
		Priority:          data.Priority,
	}

}

func FromCoreList(data []todo.TodoCore) []TodoResponse {
	var list []TodoResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
