package service

import (
	req "IoT-GPS-Backend/dto/request"
	res "IoT-GPS-Backend/dto/response"
	"IoT-GPS-Backend/entity"
	"IoT-GPS-Backend/repository"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type IoTDeviceService interface {
	RegisterNewDevice(request req.CreateDeviceRequest) (res.CreateDeviceResponse, error)
}

type ioTDeviceService struct {
	ioTDeviceRepository repository.IoTDeviceRepository
	apiKeyService       ApiKeyService
}

func NewIoTDeviceService(ioTDeviceRepository repository.IoTDeviceRepository, apiKeyService ApiKeyService) IoTDeviceService {
	return &ioTDeviceService{ioTDeviceRepository: ioTDeviceRepository, apiKeyService: apiKeyService}
}

func (i *ioTDeviceService) RegisterNewDevice(request req.CreateDeviceRequest) (res.CreateDeviceResponse, error) {
	device := entity.IoTDevice{
		Id:        uuid.New().String(),
		DeviceId:  request.DeviceId,
		CreatedAt: time.Now(),
	}

	fmt.Println(device)

	newDevice, err := i.ioTDeviceRepository.Save(device)
	if err != nil {
		return res.CreateDeviceResponse{}, err
	}

	fmt.Println(newDevice)

	newKey, err := i.apiKeyService.CreateApiKey(newDevice.Id)
	if err != nil {
		return res.CreateDeviceResponse{}, err
	}

	return res.CreateDeviceResponse{
		Id:     newDevice.Id,
		ApiKey: newKey,
	}, nil
}
