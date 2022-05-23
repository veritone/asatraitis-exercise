package registry

import (
	"time"
	"translate/types"
)

type GetRegistryResponse struct {
		Data struct {
			DataRegistries struct {
				Records []struct {
					ID              string    `json:"id"`
					Name            string    `json:"name"`
					CreatedDateTime time.Time `json:"createdDateTime"`
					Description     string    `json:"description"`
					Source          string    `json:"source"`
				} `json:"records"`
			} `json:"dataRegistries"`
		} `json:"data"`
		Errors types.Errors `json:"errors"`
}

type GetRegistryInput struct {
	Name string `json:"name"`
}

type CreateDataRegistryResponse struct {
	Data struct {
		CreateDataRegistry struct {
			ID string `json:"id"`
		} `json:"createDataRegistry"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type CreateDataRegistryInput struct {
	Name string `json:"name"`
	Description string `json:"description"`
}