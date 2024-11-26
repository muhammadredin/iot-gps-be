package service

import (
	"IoT-GPS-Backend/constants"
	req "IoT-GPS-Backend/dto/request"
	res "IoT-GPS-Backend/dto/response"
	"IoT-GPS-Backend/entity"
	"IoT-GPS-Backend/repository"
	"errors"
	"github.com/google/uuid"
	"time"
)

type GeoLogService interface {
	CreateGeoLog(request req.CreateGeoLogRequest) (res.GeoLogResponse, error)
}

type geoLogService struct {
	geoLogRepository repository.GeoLogRepository
	uptimeService    UptimeService
}

func NewGeoLogService(geoLogRepository repository.GeoLogRepository, uptimeService UptimeService) GeoLogService {
	return &geoLogService{geoLogRepository, uptimeService}
}

func (g *geoLogService) CreateGeoLog(request req.CreateGeoLogRequest) (res.GeoLogResponse, error) {
	currentUptime, err := g.uptimeService.GetUptimeById(request.UptimeId)
	if err != nil {
		return res.GeoLogResponse{}, err
	}

	if !currentUptime.EndAt.IsZero() {
		return res.GeoLogResponse{}, errors.New(constants.UptimeUnauthorizedError)
	}

	geoLog := entity.GeoLog{
		Id:        uuid.New().String(),
		UptimeId:  request.UptimeId,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		Timestamp: time.Now(),
	}

	savedGeoLog, err := g.geoLogRepository.Create(geoLog)
	if err != nil {
		return res.GeoLogResponse{}, err
	}

	return res.GeoLogResponse{
		Id:        savedGeoLog.Id,
		UptimeId:  savedGeoLog.UptimeId,
		Longitude: savedGeoLog.Longitude,
		Latitude:  savedGeoLog.Latitude,
		Timestamp: savedGeoLog.Timestamp,
	}, nil
}
