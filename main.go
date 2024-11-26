package main

import (
	"IoT-GPS-Backend/handler"
	"IoT-GPS-Backend/repository"
	"IoT-GPS-Backend/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=Enigma dbname=iot_gps host=localhost port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	ioTDeviceRepository := repository.NewIotDeviceRepository(db)
	apiKeyRepository := repository.NewApiKeyRepository(db)
	uptimeRepository := repository.NewUptimeRepository(db)
	geoLogRepository := repository.NewGeoLogRepository(db)

	apiKeyService := service.NewApiKeyService(apiKeyRepository)
	ioTDeviceService := service.NewIoTDeviceService(ioTDeviceRepository, apiKeyService)
	uptimeService := service.NewUptimeService(uptimeRepository)
	geoLogService := service.NewGeoLogService(geoLogRepository, uptimeService)

	iotDeviceHandler := handler.NewIoTDeviceHandler(ioTDeviceService, uptimeService, geoLogService)

	r := gin.Default()

	iot := r.Group("/iot")
	{
		iot.POST("/register", iotDeviceHandler.HandleRegisterDevice)
		iot.POST("/uptime/:deviceId", iotDeviceHandler.HandleCreateUptime)
		iot.PUT("/uptime/:deviceId/:uptimeId", iotDeviceHandler.HandleEndUptime)
		iot.POST("/geolog", iotDeviceHandler.HandleCreateGeoLog)
	}

	err = r.Run(":8081")
	if err != nil {
		return
	}
}
