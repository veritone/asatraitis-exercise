package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"translate/utils"
)

const API_ENDPOINT = "https://api.dev.us-1.veritone.com/v3/graphql"

var client = &http.Client{}

func getNewRequest(gql string, variables string) *http.Request {
	jsonValue, err := json.Marshal(GraphQLRequest{Query: gql, Variables: getGqlVariables(variables)})
	utils.HandleErr(err)
	req, err := http.NewRequest("POST", API_ENDPOINT, bytes.NewBuffer(jsonValue))
	utils.HandleErr(err)
	req.Header.Add("content-type", "application/json")
	token, err := readToken()
	if err != nil {
		fmt.Println("No Token file. Login to access authorized actions.")
		token = ""
	}
	req.Header.Add("authorization", "Bearer "+token)
	return req
}

func GqlRequest(gql string, variables string) []byte {
	req := getNewRequest(gql, variables)
	res, err := client.Do(req)
	utils.HandleErr(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	utils.HandleErr(err)
	return data
}