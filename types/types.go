package types

type Errors []struct {
	Message string `json: "messgae"`
	Name    string `json: "name"`
}