package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"translate/config"
	"translate/utils"
)

var client = &http.Client{}

func getNewRequest(gql string, variables string) *http.Request {
	jsonValue, err := json.Marshal(GraphQLRequest{Query: gql, Variables: getGqlVariables(variables)})
	utils.HandleErr(err)
	req, err := http.NewRequest("POST", config.API_ENDPOINT, bytes.NewBuffer(jsonValue))
	utils.HandleErr(err)
	req.Header.Add("content-type", "application/json")
	auth, err := readAuthToken()
	if err != nil {
		fmt.Println("No Token file. Login to access authorized actions.")
		auth.Token = ""
	}
	req.Header.Add("authorization", "Bearer "+auth.Token)
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