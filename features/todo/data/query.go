package data

import (
	"skyshi/features/todo"

	"gorm.io/gorm"
)

type dataTodo struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.DataInterface {
	return &dataTodo{
		db,
	}
}

func (repo *dataTodo) PostTodo(data todo.TodoCore) (int, todo.TodoCore, error) {
	var todoPost Todo
	todoPost.Title = data.Title
	todoPost.Activity_Group_Id = data.Activity_Group_Id
	todoPost.Is_Active = data.Is_Active

	tx := repo.db.Create(&todoPost)
	if tx.Error != nil {
		return -1, todo.TodoCore{}, tx.Error
	}
	tx1 := repo.db.Where("title = ? AND activity_group_id= ? AND is_active= ?", todoPost.Title, todoPost.Activity_Group_Id, todoPost.Is_Active).First(&todoPost)
	if tx1.Error != nil {
		return 0, todo.TodoCore{}, tx.Error
	}
	todoData := todoPost.toCore()
	return int(todoData.ID), todoData, nil
}

func (repo *dataTodo) SelectAllTodo(activity_group_id int) ([]todo.TodoCore, error) {
	var todos []Todo

	if activity_group_id != 0 {
		tx := repo.db.Where("activity_group_id = ?", activity_group_id).Find(&todos)
		if tx.Error != nil {
			return []todo.TodoCore{}, tx.Error
		}
	} else {
		tx := repo.db.Find(&todos)
		if tx.Error != nil {
			return []todo.TodoCore{}, tx.Error
		}
	}
	todoCore := toCoreList(todos)
	return todoCore, nil
}

func (repo *dataTodo) SelectOneTodo(id int) (todo.TodoCore, error) {
	var todoList Todo
	todoList.ID = uint(id)
	tx := repo.db.Where("id = ?", id).First(&todoList)
	if tx.Error != nil {
		return todo.TodoCore{}, tx.Error
	}
	todoData := todoList.toCore()
	return todoData, nil
}

func (repo *dataTodo) DeleteTodo(id int) (row int, err error) {
	var todo Todo
	todo.ID = uint(id)
	tx := repo.db.Unscoped().Delete(&todo)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataTodo) PatchTodo(data todo.TodoCore) (int, error) {
	var todoUpdate Todo
	txDataOld := repo.db.First(&todoUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Title != "" {
		todoUpdate.Title = data.Title
	}

	if data.Priority != "" {
		todoUpdate.Priority = data.Priority
	}

	tx := repo.db.Save(&todoUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
