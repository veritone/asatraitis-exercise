package folder

import "translate/types"

type GetRootFolderResponse struct {
	Data struct {
		RootFolders []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"rootFolders"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}

type CreateFolderResponse struct {
	Data struct {
		CreateFolder struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"createFolder"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type CreateFolderInput struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ParentId string `json:"parentId"`
}

type MoveTdoToFolderResponse struct {
	Data struct {
		FileTemporalDataObject struct {
			ID string `json:"id"`
		} `json:"fileTemporalDataObject"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type MoveTdoToFolderInput struct {
	TdoId string `json:"tdoId"`
	FolderId string `json:"folderId"`
}