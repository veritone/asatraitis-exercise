package appdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"translate/utils"
)

const DATA_FILENAME = "app_data.json"
var data AppData
var dataLoaded = false

func loadAppData() AppData {
	byteSlice, err := ioutil.ReadFile(DATA_FILENAME)

	if (err != nil && err.Error() == "open app_data.json: no such file or directory") {
		fmt.Println("AppData not yet initialized. Initializing..")
		jsonData, err := json.Marshal(data)
		utils.HandleErr(err)
		err = ioutil.WriteFile(DATA_FILENAME, []byte(jsonData), 0666)
		utils.HandleErr(err)
		dataLoaded = true
		return data
	}
	jsonAppData := AppData{}
	json.Unmarshal(byteSlice, &jsonAppData)
	dataLoaded = true
	return jsonAppData
}
func Commit() error {
	dataToCommit := GetAppData()
	jsonData, err := json.Marshal(dataToCommit)
	utils.HandleErr(err)
	return ioutil.WriteFile(DATA_FILENAME, []byte(jsonData), 0666)
}

func GetAppData() AppData {
if (!dataLoaded) {
		// Data is not initialized; Load it from file
		// if no file, return empty AppData
		data = loadAppData()
		return data;
	} else {
		// Return Data
		return data;
	}
}

func SetRegistryId(id string) {
	data.RegistryId = id
}
func SetSchemaId(id string) {
	data.SchemaId = id
}
func SetSchemaStatus(status string) {
	data.SchemaStatus = status
}
func SetCLIFolderId(id string) {
	data.FolderId = id
}
func AddJob(job Job) {
	data.Jobs = append(data.Jobs, job)
}
func GetJobs() []Job {
	GetAppData()
	return data.Jobs
}