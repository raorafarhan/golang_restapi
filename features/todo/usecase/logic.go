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

func (usecase *todoUsecase) CreateTodo(data todo.TodoCore) (row int, err error) {
	if data.Title == "" || data.Activity_Group_Id == 0 {
		return -1, errors.New("title cannot be null")
	}
	row, err = usecase.todoData.PostTodo(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *todoUsecase) GetAllTodo(activity_group_id int) (data []todo.TodoCore, err error) {
	data, err = usecase.todoData.SelectAllTodo(activity_group_id)
	if err != nil {
		return nil, err
	}
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