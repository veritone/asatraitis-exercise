package graphql

type GraphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}