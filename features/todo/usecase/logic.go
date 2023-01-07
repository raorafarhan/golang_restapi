package usecase

import (
	"errors"
	"skyshi/features/todo"
)

type todoUsecase struct {
	todoData todo.DataInterface
}

func NewTodoUsecase(data todo.DataInterface) todo.UsecaseInterface {
	return &todoUsecase{
		data,
	}
}

func (usecase *todoUsecase) CreateTodo(data todo.TodoCore) (id int, data1 todo.TodoCore, err error) {
	if data.Title == "" {
		return 0, todo.TodoCore{}, errors.New("title cannot be null")
	}

	if data.Activity_Group_Id < 1 {
		return 0, todo.TodoCore{}, errors.New("activity_group_id cannot be null")
	}

	id, _, err = usecase.todoData.PostTodo(data)
	if err != nil {
		return 0, todo.TodoCore{}, err
	}
	return id, todo.TodoCore{}, err
}

func (usecase *todoUsecase) GetAllTodo(activity_group_id int) (data []todo.TodoCore, err error) {
	data, err = usecase.todoData.SelectAllTodo(activity_group_id)
	return data, err
}

func (usecase *todoUsecase) GetOneTodo(id int) (data todo.TodoCore, err error) {
	data, err = usecase.todoData.SelectOneTodo(id)
	if err != nil {
		return todo.TodoCore{}, err
	} else if data.ID == 0 {
		return todo.TodoCore{}, errors.New(err.Error())
	} else { // data.ID != 0
		return data, err
	}
}

func (usecase *todoUsecase) UpdateTodo(data todo.TodoCore) (row int, err error) {

	row, err = usecase.todoData.PatchTodo(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *todoUsecase) DeleteTodo(id int) (row int, err error) {
	row, err = usecase.todoData.DeleteTodo(id)
	if err != nil {
		return -1, err
	}
	return row, err
}
