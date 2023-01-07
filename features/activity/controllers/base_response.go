package controllers

import (
	"net/http"
	"skyshi/features/todo"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  string      `json:"status"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data,omitempty"`
}

type DeleteResponse struct {
	Status  string        `json:"status"`
	Massage string        `json:"massage"`
	Data    todo.TodoCore `json:"data"`
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Status = "Success"
	response.Massage = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}
func SuccessCreatedResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Status = "Success"
	response.Massage = "Success"
	response.Data = data

	return c.JSON(http.StatusCreated, response)
}

func FailedResponseNotFound(c echo.Context, data string) error {
	response := BaseResponse{}
	response.Status = "Not Found"
	response.Massage = "Activity with ID " + data + " Not Found"
	response.Data = nil

	return c.JSON(http.StatusNotFound, response)
}

func SuccessDeleteResponse(c echo.Context, data todo.TodoCore) error {
	response := DeleteResponse{}
	response.Status = "Success"
	response.Massage = "Success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func FailedResponseBadRequest(c echo.Context) error {
	response := BaseResponse{}
	response.Status = "Bad Request"
	response.Massage = "title cannot be null"
	response.Data = nil

	return c.JSON(http.StatusBadRequest, response)
}
