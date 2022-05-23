package graphql

import (
	"encoding/json"
	"io/ioutil"
	"translate/types"
)

func readToken() (string, error) {
	byteSlice, err := ioutil.ReadFile(".token")
	return string(byteSlice), err
}

func readAuthToken() (types.Auth, error) {
	byteSlice, err := ioutil.ReadFile("auth.json")
	jsonAuth := types.Auth{}
	json.Unmarshal(byteSlice, &jsonAuth)
	return jsonAuth, err
}

func getGqlVariables(varArg string) string {
	if varArg == "" {
		return "{}"
	}
	return varArg
}