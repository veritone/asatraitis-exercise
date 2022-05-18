package query

import (
	"encoding/json"
	"translate/graphql"
	"translate/query/login"
	"translate/query/me"
	"translate/query/status"
	translateJob "translate/query/translate"
	"translate/utils"
)

func Me() me.MeResponse {
	data := graphql.GqlRequest(me.ME_QUERY, "")
	jsonRes := me.MeResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func Login(userName string, password string) login.LoginResponse {
	inputArgs, err := json.Marshal(login.LoginInput{UserName: userName, Password: password})

	utils.HandleErr(err)

	data := graphql.GqlRequest(login.LOGIN_QUERY, string(inputArgs))
	jsonRes := login.LoginResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func CreateTranslateJob(clusterId string, url string, language string) translateJob.CreateTranslateJobResponse {
	inputArgs, err := json.Marshal(translateJob.CreateTranslateJobInput{
		ClusterId: utils.GetClusterId(clusterId), 
		Url: url, 
		Language: language,
	})
	utils.HandleErr(err)

	data := graphql.GqlRequest(translateJob.TRANSLATE_JOB_QUERY, string(inputArgs))
	jsonRes := translateJob.CreateTranslateJobResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func JobStatus(id string) status.StatusResponse {
	inputArgs, err := json.Marshal(status.StatusInput{
		Id: id,
	})
	utils.HandleErr(err)
	data := graphql.GqlRequest(status.JOB_STATUS_QUERY, string(inputArgs))
	jsonRes := status.StatusResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}