package status

import (
	"time"
	"translate/types"
)

	
type StatusResponse struct {
	Data Data `json:"data"`
	Errors types.Errors `json:"errors"`
}
type Job struct {
	ID              string    `json:"id"`
	TargetID        string    `json:"targetId"`
	ClusterID       string    `json:"clusterId"`
	Status          string    `json:"status"`
	CreatedDateTime time.Time `json:"createdDateTime"`
	Tasks Task `json:"tasks"`
}
type Task struct {
	Records []Record `json:"records"`
}
type Record struct {
	ID     string `json:"id"`
	Engine Engine `json:"engine"`
	Status string `json:"status"`
}
type Engine struct {
	Name string `json:"name"`
}
type Data struct {
	Job Job `json:"job"`
}

// Input args
type StatusInput struct {
	Id string `json:"id"`
}