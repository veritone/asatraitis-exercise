package appdata

type AppData struct {
	RegistryId string `json:"registryId"`
	SchemaId string `json:"schemaId"`
	SchemaStatus string `json:"schemaStatus"`
	FolderId string `json:"folderId"`
	Jobs []Job `json:"jobs"`
}
type Job struct {
	Id string `json:"id"`
	TargetId string `json:"targetId"`
	Status string `json:"status"`
	JobTitle string `json:"jobTitle"`
	Tasks []Task `json:"tasks"`
	Results []Result `json:"results"`
}
type Task struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}
type Result struct {
	Id string `json:"id"`
	AssetType string `json:"assetType"`
	SignedUri string `json:"signedUri"`
}