package response

import "time"

type GeoLogResponse struct {
	Id        string    `json:"id"`
	UptimeId  string    `json:"uptime_id"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Timestamp time.Time `json:"timestamp"`
}
