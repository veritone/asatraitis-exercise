package folder

const GET_ROOT_FOLDER = `
	query listRootFolders {
		rootFolders {
			id
			name
		}
	}
`

const CREATE_CLI_FOLDER = `
	mutation createFolder($name: String!, $parentId: ID!, $description: String!) {
		createFolder(input: {
			name: $name
			description: $description
			parentId: $parentId
			rootFolderType: collection
		}) {
			id
			name
		}
	}
`

const MOVE_TDO_TO_FOLDER = `
	mutation fileTdo($tdoId: ID!, $folderId: ID!) {
		fileTemporalDataObject (input: {
			tdoId: $tdoId
			folderId: $folderId
		}) {
			id
		}
	}
`