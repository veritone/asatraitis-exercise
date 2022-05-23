package registry

const GET_DATA_REGISTRY = `
	query dataRegistries($name: String!) {
		dataRegistries(
			name: $name
		) {
			records {
				id
				name
			}
		}
	}
`

const CREATE_DATA_REGISTRY = `
	mutation createDataRegistry($name: String!, $description: String!) {
		createDataRegistry(
			input: {
				name: $name
				description: $description
				source: "N/A"
			}
		) {
			id
		}
	}
`