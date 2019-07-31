package service

import "universalsdk/models"

type UsdkService interface {
	DeviceCheck(deviceCheckCollection models.DeviceCheckDetailsObjectCollection) (*models.PuppyObject, error)
}
