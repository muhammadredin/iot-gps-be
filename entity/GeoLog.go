package entity

import "time"

type GeoLog struct {
	Id        string
	UptimeId  string
	Longitude float64
	Latitude  float64
	Timestamp time.Time
}
