package repository

import (
	"IoT-GPS-Backend/entity"
	"database/sql"
	"fmt"
	"log"
)

type IoTDeviceRepository interface {
	GetAll() ([]entity.IoTDevice, error)
	GetById(id string) (entity.IoTDevice, error)
	Save(device entity.IoTDevice) (entity.IoTDevice, error)
}

type ioTDeviceRepository struct {
	db *sql.DB
}

func NewIotDeviceRepository(db *sql.DB) IoTDeviceRepository {
	return &ioTDeviceRepository{db}
}

func (i *ioTDeviceRepository) GetAll() ([]entity.IoTDevice, error) {
	rows, err := i.db.Query("SELECT id, device_id, created_at FROM m_iot_device")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	iotDevices := make([]entity.IoTDevice, 0)

	for rows.Next() {
		var device entity.IoTDevice
		if err := rows.Scan(&device.Id, &device.DeviceId, &device.CreatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		iotDevices = append(iotDevices, device)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return iotDevices, nil
}

func (i *ioTDeviceRepository) GetById(id string) (entity.IoTDevice, error) {
	var device entity.IoTDevice
	err := i.db.QueryRow("SELECT * FROM m_iot_device WHERE id = $1", id).
		Scan(&device.Id, &device.DeviceId, &device.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return device, err
	}

	return device, nil
}

func (i *ioTDeviceRepository) Save(device entity.IoTDevice) (entity.IoTDevice, error) {
	_, err := i.db.Exec("INSERT INTO m_iot_device (id, device_id, created_at) VALUES ($1, $2, $3)", device.Id, device.DeviceId, device.CreatedAt)
	if err != nil {
		return entity.IoTDevice{}, err
	}
	fmt.Println("Successfully saved device")
	return device, nil
}
