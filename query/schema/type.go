package schema

import "translate/types"

type CreateSchemaDraftResponse struct {	
	Data struct {
		UpsertSchemaDraft struct {
			ID string `json:"id"`
		} `json:"upsertSchemaDraft"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type CreateSchemaDraftInput struct {
	DataRegistryId string `json:"dataRegistryId"`
}

type PublishDraftSchemaResponse struct {
	Data struct {
		UpdateSchemaState struct {
			ID           string `json:"id"`
			Status       string `json:"status"`
		} `json:"updateSchemaState"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type PublishDraftSchemaInput struct {
	Id string `json:"id"`
}