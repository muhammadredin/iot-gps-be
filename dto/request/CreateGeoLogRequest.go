package request

type CreateGeoLogRequest struct {
	UptimeId  string  `json:"uptime_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
