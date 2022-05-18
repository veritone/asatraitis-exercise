package me
import (
	"translate/types"
)
// Me response struct
type Me struct {
	Name string `json: "name"`
}
type Data struct {
	Me Me `json:"me"`
}
type MeResponse struct {
	Data   Data   `json:"data"`
	Errors types.Errors `json: "errors"`
}