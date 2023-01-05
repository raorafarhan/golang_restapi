package factory

import (
	activityData "skyshi/features/activity/data"
	activityDelivery "skyshi/features/activity/delivery"
	activityUsecase "skyshi/features/activity/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	activityDataFactory := activityData.New(db)
	activityUsecaseFactory := activityUsecase.NewActivityUsecase(activityDataFactory)
	activityDelivery.New(e, activityUsecaseFactory)

}
