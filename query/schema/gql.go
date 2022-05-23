package schema

const CREAT_SCHEMA_DRAFT = `
mutation createSchemaDraft($dataRegistryId: ID!) {
  upsertSchemaDraft(input: {
    dataRegistryId: $dataRegistryId
    majorVersion: 1
    schema: {
      type: "object",
      title: "properties",
      required: [
        "email",
				"originalFileUrl",
				"translatedTo"
      ],
      properties: {
        email: {
          type: "string"
        },
				jobTitle: {
					type: "string"
				},
        originalFileUrl: {
          type: "string"
        },
				translatedTo: {
					type: "string"
				}
      },
      description: "Additional details for translation job."      
    }
  }) {
    id
  }
}
`

const PUBLISH_SCHEMA_DRAFT = `
mutation publishSchemaDraft($id: ID!) {
  updateSchemaState(input: {
    id: $id
    status: published
  }) {
    id
    status
  }
}
`