package entity

import "time"

type Uptime struct {
	Id          string
	IoTDeviceId string
	StartAt     time.Time
	EndAt       time.Time
}
