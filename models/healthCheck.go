package models

// HealthCheck ...
type HealthCheck struct {
	DbUP   string `json:"dbUP"`
	Status string `json:"status"`
}
