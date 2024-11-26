package service

import (
	"IoT-GPS-Backend/constants"
	res "IoT-GPS-Backend/dto/response"
	"IoT-GPS-Backend/entity"
	"IoT-GPS-Backend/repository"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UptimeService interface {
	CreateUptime(deviceId string) (res.CreateUptimeResponse, error)
	EndUptime(deviceId string, uptimeId string) error
}

type uptimeService struct {
	uptimeRepository repository.UptimeRepository
}

func NewUptimeService(uptimeRepository repository.UptimeRepository) UptimeService {
	return &uptimeService{uptimeRepository: uptimeRepository}
}

func (u *uptimeService) CreateUptime(deviceId string) (res.CreateUptimeResponse, error) {
	uptime := entity.Uptime{
		Id:          uuid.New().String(),
		IoTDeviceId: deviceId,
		StartAt:     time.Now(),
		EndAt:       time.Time{},
	}

	newUptime, err := u.uptimeRepository.Create(uptime)
	if err != nil {
		return res.CreateUptimeResponse{}, err
	}

	return res.CreateUptimeResponse{Id: newUptime.Id}, nil
}

func (u *uptimeService) EndUptime(deviceId string, uptimeId string) error {
	uptime, err := u.uptimeRepository.GetById(uptimeId)
	if err != nil {
		return err
	}

	if uptime.IoTDeviceId != deviceId {
		return errors.New(constants.UptimeUnauthorizedError)
	}

	if !uptime.EndAt.IsZero() {
		return errors.New(constants.BadRequestError)
	}

	uptime.EndAt = time.Now()
	_, err = u.uptimeRepository.Update(uptime)
	if err != nil {
		return err
	}
	return nil
}
