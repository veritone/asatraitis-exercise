package query

import (
	"encoding/json"
	"translate/config"
	"translate/graphql"
	"translate/query/folder"
	"translate/query/login"
	"translate/query/me"
	"translate/query/registry"
	"translate/query/schema"
	"translate/query/sdo"
	"translate/query/sosearch"
	"translate/query/status"
	"translate/query/tdo"
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

func GetRegistry() registry.GetRegistryResponse {
		inputArgs, err := json.Marshal(registry.GetRegistryInput{Name: config.REGISTRY_NAME})
		utils.HandleErr(err)
		data := graphql.GqlRequest(registry.GET_DATA_REGISTRY, string(inputArgs))
		jsonRes := registry.GetRegistryResponse{}
		json.Unmarshal(data, &jsonRes)
		return jsonRes
}

func CreateRegistry() registry.CreateDataRegistryResponse {
	inputArgs, err := json.Marshal(registry.CreateDataRegistryInput{Name: config.REGISTRY_NAME, Description: config.REGISTRY_DESCRIPTION})
	utils.HandleErr(err)
	data := graphql.GqlRequest(registry.CREATE_DATA_REGISTRY, string(inputArgs))
	jsonRes := registry.CreateDataRegistryResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func CreateSchemaDraft(regId string) schema.CreateSchemaDraftResponse {
	inputArgs, err := json.Marshal(schema.CreateSchemaDraftInput{DataRegistryId: regId})
	utils.HandleErr(err)
	data := graphql.GqlRequest(schema.CREAT_SCHEMA_DRAFT, string(inputArgs))
	jsonRes := schema.CreateSchemaDraftResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func PublishSchemaDraft(schemaId string) schema.PublishDraftSchemaResponse {
	inputArgs, err := json.Marshal(schema.PublishDraftSchemaInput{Id: schemaId})
	utils.HandleErr(err)
	data := graphql.GqlRequest(schema.PUBLISH_SCHEMA_DRAFT, string(inputArgs))
	jsonRes := schema.PublishDraftSchemaResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func GetRootFolder() folder.GetRootFolderResponse {
	data := graphql.GqlRequest(folder.GET_ROOT_FOLDER, "")
	jsonRes := folder.GetRootFolderResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func CreateCLIFolder(name string, description string, parentId string) folder.CreateFolderResponse {
	inputArgs, err := json.Marshal(folder.CreateFolderInput{Name: name, Description: description, ParentId: parentId})
	utils.HandleErr(err)

	data := graphql.GqlRequest(folder.CREATE_CLI_FOLDER, string(inputArgs))
	jsonRes := folder.CreateFolderResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func UpdateTdoWithContent(tdoInput sdo.UpdateTdoWithContentInput) sdo.UpdateTdoWithContentResponse {
	inputArgs, err := json.Marshal(tdoInput)
	utils.HandleErr(err)

	data := graphql.GqlRequest(sdo.UPDATE_TDO_WITH_CONTENT, string(inputArgs))
	jsonRes := sdo.UpdateTdoWithContentResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func MoveTDOToFolder(tdoId string, folderId string) folder.MoveTdoToFolderResponse {
	inputArgs, err := json.Marshal(folder.MoveTdoToFolderInput{TdoId: tdoId, FolderId: folderId})
	utils.HandleErr(err)

	data := graphql.GqlRequest(folder.MOVE_TDO_TO_FOLDER, string(inputArgs))
	jsonRes := folder.MoveTdoToFolderResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func SearchByTitle(value string) sosearch.SearchResponse {
	inputArgs, err := json.Marshal(sosearch.SearchInput{Value: "*"+value+"*"})
	utils.HandleErr(err)

	data := graphql.GqlRequest(sosearch.SEARCH_BY_JOB_TITLE, string(inputArgs))
	jsonRes := sosearch.SearchResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}

func GetTDO(id string) tdo.TdoResponse {
	inputArgs, err := json.Marshal(tdo.TdoInput{Id: id})
	utils.HandleErr(err)

	data := graphql.GqlRequest(tdo.GET_TDO, string(inputArgs))
	jsonRes := tdo.TdoResponse{}
	json.Unmarshal(data, &jsonRes)
	return jsonRes
}