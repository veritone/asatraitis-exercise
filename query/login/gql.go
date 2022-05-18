package login

const LOGIN_QUERY = `
	mutation UserLogin($userName: String!, $password: String!) {
		userLogin(
		input: { userName: $userName, password: $password }
		) {
		token
		organization {
			id
			name
		}
		}
	}
`