package usecase

import (
	"errors"
	"skyshi/features/activity"
)

type activityUsecase struct {
	activityData activity.DataInterface
}

func NewActivityUsecase(data activity.DataInterface) activity.UsecaseInterface {
	return &activityUsecase{
		data,
	}
}

func (usecase *activityUsecase) CreateActivity(data activity.ActivityCore) (id int, data1 activity.ActivityCore, err error) {
	if data.Title == "" || data.Email == "" {
		return 0, activity.ActivityCore{}, errors.New("title cannot be null")
	}
	id, _, err = usecase.activityData.PostActivity(data)
	if err != nil {
		return 0, activity.ActivityCore{}, err
	}
	return id, activity.ActivityCore{}, err
}

func (usecase *activityUsecase) GetAllActivity() (data []activity.ActivityCore, err error) {
	data, err = usecase.activityData.SelectAllActivity()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (usecase *activityUsecase) GetOneActivity(id int) (data activity.ActivityCore, err error) {
	data, err = usecase.activityData.SelectOneActivity(id)
	if err != nil {
		return activity.ActivityCore{}, err
	} else if data.ID == 0 {
		return activity.ActivityCore{}, errors.New("user not found")
	} else { // data.ID != 0
		return data, err
	}
}

func (usecase *activityUsecase) UpdateActivity(data activity.ActivityCore) (row int, err error) {

	row, err = usecase.activityData.PatchActivity(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *activityUsecase) DeleteActivity(id int) (row int, err error) {
	row, err = usecase.activityData.DeleteActivity(id)
	if err != nil {
		return -1, err
	}
	return row, err
}
