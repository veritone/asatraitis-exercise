package translateJob

import (
	"time"
	"translate/types"
)

// translateJob response
type CreateTranslateJobResponse struct {
	Data   CreateTranslateJobData `json:"data"`
	Errors types.Errors                 `json:"errors"`
}
type CreateJob struct {
	ID              string    `json:"id"`
	TargetID        string    `json:"targetId"`
	CreatedDateTime time.Time `json:"createdDateTime"`
}
type CreateTranslateJobData struct {
	CreateJob CreateJob `json:"createJob"`
}

// translateJob input type
type CreateTranslateJobInput struct {
	Url string `json:"url"`
	ClusterId string `json:"clusterId"`
	Language string `json:"language"` 
}