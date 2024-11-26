package entity

import "time"

type ApiKey struct {
	Id          string
	IoTDeviceId string
	Key         string
	CreatedAt   time.Time
}
