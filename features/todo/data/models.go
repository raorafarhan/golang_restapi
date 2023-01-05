package data

import (
	"skyshi/features/todo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity_Group_Id uint
	Title             string
	Is_Active         bool
	Priority          string `gorm:"default:very-high"`
}

func fromCore(data todo.TodoCore) Todo {
	return Todo{
		Activity_Group_Id: data.Activity_Group_Id,
		Title:             data.Title,
		Is_Active:         data.Is_Active,
		Priority:          data.Priority,
	}
}

func (data *Todo) toCore() todo.TodoCore {
	return todo.TodoCore{
		ID:                data.ID,
		Title:             data.Title,
		Activity_Group_Id: data.Activity_Group_Id,
		Is_Active:         data.Is_Active,
		Priority:          data.Priority,
	}
}

func toCoreList(data []Todo) []todo.TodoCore {
	var list []todo.TodoCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
