package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"translate/config"
	"translate/query/status"
	"translate/types"
)

func HandleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func HandleGraphQLError (errors types.Errors) {
	if (errors != nil) {
		fmt.Printf("ERROR: %v \n", errors[0].Message)
		os.Exit(1)
	}
}

func SaveToken(token string) error {
	return ioutil.WriteFile(".token", []byte(token), 0666)
}
func SaveAuthInfo(token string, email string, id string) error {
	authJson, err := json.Marshal(types.Auth{Token: token, Email: email, Id: id})
	if (err != nil) {
		return err
	}
	return ioutil.WriteFile("auth.json", []byte(authJson), 0666)
}
func GetAuthInfo() types.Auth {
	bSlice, err := ioutil.ReadFile("auth.json")
	HandleErr(err)
	jsonData := types.Auth{}
	json.Unmarshal(bSlice, &jsonData)
	return jsonData
}

func GetClusterId(clusterId string) string {
	if clusterId == "" {
		return config.CLUSTER_ID
	}
	return clusterId
}

func GetJobs() []string {
	bSlice, err := ioutil.ReadFile("jobs")
	if (err != nil && err.Error() == "open jobs: no such file or directory") {
		return []string{}
	}	
	return strings.Split(string(bSlice), ",")
}

func SaveJob(jobId string) {
	bSlice, err := ioutil.ReadFile("jobs")
	if (err != nil && err.Error() == "open jobs: no such file or directory") {
		ioutil.WriteFile("jobs", []byte(jobId), 0666)
		return
	}	
	sSlice := strings.Split(string(bSlice), ",")
	sSlice = append(sSlice, jobId)
	ioutil.WriteFile("jobs", []byte(strings.Join(sSlice, ",")), 0666)
}

func GetCompletedTasks(records []status.Record) string {
	tasksCompleted := 0
	for _, task := range records {
		if (task.Status == "complete") {
			tasksCompleted++
		}
	}
	return strconv.Itoa(tasksCompleted)+"/"+strconv.Itoa(len(records))
}