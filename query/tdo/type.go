package tdo

import "time"

type TdoResponse struct {
	Data struct {
		TemporalDataObject struct {
			ID              string    `json:"id"`
			Name            string    `json:"name"`
			CreatedDateTime time.Time `json:"createdDateTime"`
			Organization    struct {
				ID              string    `json:"id"`
				Name            string    `json:"name"`
				CreatedDateTime time.Time `json:"createdDateTime"`
			} `json:"organization"`
		} `json:"temporalDataObject"`
	} `json:"data"`
}

type TdoInput struct {
	Id string `json:"id"`
}