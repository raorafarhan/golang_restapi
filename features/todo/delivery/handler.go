package delivery

import (
	"net/http"
	"skyshi/features/todo"

	controllers "skyshi/features/todo/controllers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type todoDelivery struct {
	todoUsecase todo.UsecaseInterface
}

func New(e *echo.Echo, usecase todo.UsecaseInterface) {
	handler := &todoDelivery{
		todoUsecase: usecase,
	}
	e.POST("/todo-items", handler.CreateTodo)
	e.GET("/todo-items", handler.GetAllTodo)
	e.GET("/todo-items/:id", handler.GetOneTodo)
	e.PATCH("/todo-items/:id", handler.UpdateTodo)
	e.DELETE("/todo-items/:id", handler.DeleteTodo)
}

func (handler *todoDelivery) CreateTodo(c echo.Context) error {
	var data TodoRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return controllers.FailedResponseBadRequest(c)
	}

	id, _, err := handler.todoUsecase.CreateTodo(ToCore(data))
	if err != nil {
		return controllers.FailedResponseBadRequest(c)
	}

	data1, _ := handler.todoUsecase.GetOneTodo(id)
	return controllers.SuccessCreatedResponse(c, data1)
}

func (handler *todoDelivery) GetAllTodo(c echo.Context) error {

	activity_group_id, err := strconv.Atoi(c.QueryParam("activity_group_id"))
	if err != nil && activity_group_id != 0 {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := handler.todoUsecase.GetAllTodo(activity_group_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	return controllers.NewSuccesResponse(c, data)
}

func (handler *todoDelivery) GetOneTodo(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	data, err := handler.todoUsecase.GetOneTodo(idConv)
	if err != nil {
		return controllers.FailedResponseNotFound(c, id)
	}
	return controllers.NewSuccesResponse(c, data)
}

func (handler *todoDelivery) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	_, err := handler.todoUsecase.GetOneTodo(idConv)
	if err != nil {
		return controllers.FailedResponseNotFound(c, id)
	}
	var data TodoRequest
	errBind := c.Bind(&data)
	if errBind != nil {
		return controllers.FailedResponseBadRequest(c)
	}
	updateCore := ToCore(data)
	updateCore.ID = uint(idConv)

	row, err := handler.todoUsecase.UpdateTodo(updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	dataUpdate, _ := handler.todoUsecase.GetOneTodo(idConv)
	return controllers.NewSuccesResponse(c, dataUpdate)
}

func (handler *todoDelivery) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	row, err := handler.todoUsecase.DeleteTodo(idConv)
	if err != nil {
		return controllers.FailedResponseNotFound(c, id)
	}

	if row != 1 {
		return controllers.FailedResponseNotFound(c, id)
	}
	return controllers.SuccessDeleteResponse(c, nil)
}
