package dsl

// TODO: add test

// GraphQLInteraction extends the Interaction struct by allowing GraphQL
// requests to be served over HTTP.
type GraphQLInteraction struct {
	Operation string `json:"operation"`

	Variables map[string]interface `json:"variables"`

	Query string `json:"query"`

	Interaction
}

// WithOperation sets the Operations field
func (g *GraphQLInteraction) WithOperation(operation string) (*GraphQLInteraction){
	g.Operation = operation
	return g
}

// TODO: complete/iron out
// WithVariables set the Variables field
// func (g *GraphQLInteraction) WithVariables(variables string) (*GraphQLInteraction){
// 	g.Variables = variables
// 	return g
// }

// WithQuery sets the Query field
func (g *GraphQLInteraction) WithQuery(query string) *GraphQLInteraction {
	// TODO: checks will be ideal
	g.Query = query
	return g
}

// ReturnInteraction returns a new Interaction instance which will be needed
// when verifying Interactions
func (g *GraphQLInteraction) ReturnInteraction Interaction {
	return Interaction{
		Request{
			Body: {
				Operation: g.Operation,
				Variables: g.Variables,
				// TODO: follow jq example and check to see what is available
				Query: q.Query
			},
			Method: "POST",
			Headers: MapMatcher{
				"content-type": "application/json"
			}
		},
		Response: g.Response,
		Description: g.Description,
		State: g.State
	}
}
