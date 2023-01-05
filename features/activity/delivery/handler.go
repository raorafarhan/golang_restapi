package delivery

import (
	"net/http"
	"skyshi/features/activity"

	"skyshi/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type activityDelivery struct {
	activityUsecase activity.UsecaseInterface
}

func New(e *echo.Echo, usecase activity.UsecaseInterface) {
	handler := &activityDelivery{
		activityUsecase: usecase,
	}
	e.POST("/activity-groups", handler.CreateActivity)
	e.GET("/activity-groups", handler.GetAllActivity)
	e.GET("/activity-groups/:id", handler.GetOneActivity)
	e.PATCH("/activity-groups/:id", handler.UpdateActivity)
	e.DELETE("/activity-groups/:id", handler.DeleteActivity)
}

func (handler *activityDelivery) CreateActivity(c echo.Context) error {
	var data ActivityRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return helper.FailedResponseBadRequest(c)
	}

	row, err := handler.activityUsecase.CreateActivity(ToCore(data))
	if err != nil {
		return helper.FailedResponseBadRequest(c)
	}

	if row != 1 {
		return helper.FailedResponseBadRequest(c)
	}

	return helper.SuccessCreatedResponse(c, data)
}

func (handler *activityDelivery) GetAllActivity(c echo.Context) error {
	data, err := handler.activityUsecase.GetAllActivity()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	return helper.NewSuccesResponse(c, data)
}

func (handler *activityDelivery) GetOneActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	data, err := handler.activityUsecase.GetOneActivity(idConv)
	if err != nil {
		return helper.FailedResponseNotFound(c, id)
	}
	return helper.NewSuccesResponse(c, data)
}

func (handler *activityDelivery) UpdateActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	_, err := handler.activityUsecase.GetOneActivity(idConv)
	if err != nil {
		return helper.FailedResponseNotFound(c, id)
	}
	var data ActivityRequest
	errBind := c.Bind(&data)

	if data.Title == "" {
		return helper.FailedResponseBadRequest(c)
	}

	if errBind != nil {
		return helper.FailedResponseBadRequest(c)
	}
	updateCore := ToCore(data)
	updateCore.ID = uint(idConv)

	row, err := handler.activityUsecase.UpdateActivity(updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, (err.Error()))
	}
	dataUpdate, _ := handler.activityUsecase.GetOneActivity(idConv)
	return helper.NewSuccesResponse(c, dataUpdate)
}

func (handler *activityDelivery) DeleteActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, _ := strconv.Atoi(id)
	row, err := handler.activityUsecase.DeleteActivity(idConv)
	if err != nil {
		return helper.FailedResponseNotFound(c, id)
	}

	if row != 1 {
		return helper.FailedResponseNotFound(c, id)
	}
	return helper.SuccessDeleteResponse(c, nil)
}
