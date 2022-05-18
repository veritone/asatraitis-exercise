package login

import (
	"translate/types"
)

// Login response
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type UserLogin struct {
	Token        string       `json:"token"`
	Organization Organization `json:"organization"`
}
type LoginData struct {
	UserLogin UserLogin `json: "userLogin"`
}
type LoginResponse struct {
	Errors types.Errors    `json:"errors"`
	Data   LoginData `json:"data"`
}

// Login input Type
type LoginInput struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}