package repository

import (
	"IoT-GPS-Backend/entity"
	"database/sql"
)

type GeoLogRepository interface {
	Create(geoLog entity.GeoLog) (entity.GeoLog, error)
}

type geoLogRepository struct {
	db *sql.DB
}

func NewGeoLogRepository(db *sql.DB) GeoLogRepository {
	return &geoLogRepository{db}
}

func (g *geoLogRepository) Create(geoLog entity.GeoLog) (entity.GeoLog, error) {
	_, err := g.db.Exec("INSERT INTO t_geo_log (id, uptime_id, latitude, longitude, timestamp) VALUES ($1, $2, $3, $4, $5)",
		geoLog.Id, geoLog.UptimeId, geoLog.Latitude, geoLog.Longitude, geoLog.Timestamp)
	if err != nil {
		return entity.GeoLog{}, err
	}
	return geoLog, nil
}
