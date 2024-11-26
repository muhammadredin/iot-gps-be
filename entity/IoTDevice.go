package entity

import "time"

type IoTDevice struct {
	Id        string
	DeviceId  string
	CreatedAt time.Time
}
