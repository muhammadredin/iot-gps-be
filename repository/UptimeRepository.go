package repository

import (
	"IoT-GPS-Backend/entity"
	"database/sql"
	"log"
)

type UptimeRepository interface {
	GetById(id string) (entity.Uptime, error)
	Create(uptime entity.Uptime) (entity.Uptime, error)
	Update(uptime entity.Uptime) (entity.Uptime, error)
}

type uptimeRepository struct {
	db *sql.DB
}

func NewUptimeRepository(db *sql.DB) UptimeRepository {
	return &uptimeRepository{db}
}

func (u *uptimeRepository) GetById(id string) (entity.Uptime, error) {
	var uptime entity.Uptime
	err := u.db.QueryRow("SELECT * FROM t_uptime WHERE id = $1", id).
		Scan(&uptime.Id, &uptime.IoTDeviceId, &uptime.StartAt, &uptime.EndAt)
	if err != nil {
		log.Fatal(err)
		return uptime, err
	}

	return uptime, nil
}

func (u *uptimeRepository) Create(uptime entity.Uptime) (entity.Uptime, error) {
	_, err := u.db.Exec("INSERT INTO t_uptime (id, iot_device_id, start_at, end_at) VALUES ($1, $2, $3, $4)",
		uptime.Id, uptime.IoTDeviceId, uptime.StartAt, uptime.EndAt)
	if err != nil {
		return entity.Uptime{}, err
	}
	return uptime, nil
}

func (u *uptimeRepository) Update(uptime entity.Uptime) (entity.Uptime, error) {
	_, err := u.db.Exec("UPDATE t_uptime SET end_at = $1 WHERE id = $2", uptime.EndAt, uptime.Id)
	if err != nil {
		return entity.Uptime{}, err
	}
	return uptime, nil
}
