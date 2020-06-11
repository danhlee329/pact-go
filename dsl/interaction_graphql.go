package dsl

// TODO: find better way to extend the Interaction struct (copy-pasta...)

// GraphQLInteraction extends the Interaction struct by allowing GraphQL
// requests to be served over HTTP.
type GraphQLInteraction struct {
	Operation   string                 `json:"operation"`
	Variables   map[string]interface{} `json:"variables"`
	Query       string                 `json:"query"`
	State       string                 `json:"providerState,omitempty"`
	Description string                 `json:"description"`
	Response    Response               `json:"response"`
}

type GraphQLHttpBody struct {
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
	Query     string                 `json:"query"`
}

// Given specifies a provider state. Optional.
func (i *GraphQLInteraction) Given(state string) *GraphQLInteraction {
	i.State = state

	return i
}

// UponReceiving specifies the name of the test case. This becomes the name of
// the consumer/provider pair in the Pact file. Mandatory.
func (i *GraphQLInteraction) UponReceiving(description string) *GraphQLInteraction {
	i.Description = description

	return i
}

// WithOperation sets the Operations field
func (g *GraphQLInteraction) WithOperation(operation string) *GraphQLInteraction {
	g.Operation = operation
	return g
}

// TODO: complete/iron out
// WithVariables set the Variables field
// func (g *GraphQLInteraction) WithVariables(variables string) *GraphQLInteraction {
// 	g.Variables = variables
// 	return g
// }

// WithQuery sets the Query field
func (g *GraphQLInteraction) WithQuery(query string) *GraphQLInteraction {
	// TODO: checks will be ideal
	g.Query = query
	return g
}

// WillRespondWith specifies the details of the HTTP response that will be used to
// confirm that the Provider must satisfy. Mandatory.
func (i *GraphQLInteraction) WillRespondWith(response Response) *GraphQLInteraction {
	i.Response = response

	return i
}

// ReturnInteraction returns a new Interaction instance which will be needed
// when verifying Interactions
// Reference https://graphql.org/learn/serving-over-http/#post-request for POST
// requests over HTTP
func (g *GraphQLInteraction) ReturnInteraction() *Interaction {
	// TODO: error handling

	request := Request{
		Body: GraphQLHttpBody{
			Operation: g.Operation,
			Variables: g.Variables,
			// TODO: follow js example and check to see what is available
			Query: g.Query,
		},
		Method: "POST",
		Headers: MapMatcher{
			"content-type": String("application/json"),
		},
	}

	return &Interaction{
		Request:     request,
		Response:    g.Response,
		Description: g.Description,
		State:       g.State,
	}
}
