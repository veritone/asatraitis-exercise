package sdo

const UPDATE_TDO_WITH_CONTENT = `
	mutation updateTdoWithSdo($tdoId: ID!, $schemaId: ID!, $email: String!, $jobTitle: String, $translatedTo: String!, $url: String!) {
		updateTDO(
			input: {
				id: $tdoId			
				contentTemplates: [
					{
						schemaId: $schemaId
						data: {							
							email: $email
							jobTitle: $jobTitle
							translatedTo: $translatedTo
							originalFileUrl: $url
						}
					}
				]
			}
		)
		{
			id
			status
		}
	}
`