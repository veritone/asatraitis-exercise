package types

type Errors []struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

type Auth struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Id string `json:"id"`
}

type AppData struct {

}
