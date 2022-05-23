package sdo

import "translate/types"

type UpdateTdoWithContentResponse struct {
	Data struct {
		UpdateTDO struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"updateTDO"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type UpdateTdoWithContentInput struct {
	TdoId string `json:"tdoId"`
	SchemaId string `json:"schemaId"`
	Email string `json:"email"`
	JobTitle string `json:"jobTitle"`
	TranslatedTo string `json:"translatedTo"`
	Url string `json:"url"`
}