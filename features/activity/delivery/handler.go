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
		return c.JSON(http.StatusBadRequest, helper.FailResponseCreate())
	}

	row, err := handler.activityUsecase.CreateActivity(ToCore(data))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailResponseCreate())
	}

	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.FailResponseCreate())
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseGetCreatePatch(data))
}

func (handler *activityDelivery) GetAllActivity(c echo.Context) error {
	data, err := handler.activityUsecase.GetAllActivity()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseGetCreatePatch(data))
}

func (handler *activityDelivery) GetOneActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(errConv.Error()))
	}
	data, err := handler.activityUsecase.GetOneActivity(idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailRespGetOne("Activity with ID "+id+" Not Found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseGetCreatePatch(data))
}

func (handler *activityDelivery) UpdateActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(errConv.Error()))
	}
	_, err := handler.activityUsecase.GetOneActivity(idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailRespGetOne("Activity with ID "+id+" Not Found"))
	}
	var data ActivityRequest
	errBind := c.Bind(&data)

	if data.Title == "" {
		return c.JSON(http.StatusBadRequest, helper.FailResponseCreate())
	}

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailResponseCreate())
	}
	updateCore := ToCore(data)
	updateCore.ID = uint(idConv)

	row, err := handler.activityUsecase.UpdateActivity(updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(err.Error()))
	}
	dataUpdate, _ := handler.activityUsecase.GetOneActivity(idConv)
	return c.JSON(http.StatusOK, helper.SuccessResponseGetCreatePatch(dataUpdate))
}

func (handler *activityDelivery) DeleteActivity(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailResp(errConv.Error()))
	}
	row, err := handler.activityUsecase.DeleteActivity(idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailRespGetOne("Activity with ID "+id+" Not Found"))
	}

	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.FailRespGetOne("Activity with ID "+id+" Not Found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseGetCreatePatch(nil))
}
