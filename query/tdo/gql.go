package tdo

const GET_TDO = `
	query getTdo($id: ID!){
		temporalDataObject(
			id: $id
		) {
			id
			name
			createdDateTime
			organization{ id, name,  createdDateTime}
		}
	}
`