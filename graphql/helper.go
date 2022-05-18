package graphql

import (
	"io/ioutil"
)

func readToken() (string, error) {
	byteSlice, err := ioutil.ReadFile(".token")
	return string(byteSlice), err
}

func getGqlVariables(varArg string) string {
	if varArg == "" {
		return "{}"
	}
	return varArg
}