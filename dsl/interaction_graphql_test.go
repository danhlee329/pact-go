package dsl

import (
	"testing"
)

func TestGraphQLInteraction_NewGraphQLInteraction(t *testing.T) {
	given := "testGiven"
	uponReceiving := "testUponReceiving"
	operationName := "testOperation"
	query := "testQuery"

	i := (&GraphQLInteraction{}).
		Given(given).
		UponReceiving(uponReceiving).
		WithOperation(operationName).
		WithQuery(query).
		WillRespondWith(Response{})

	if i.State != given {
		t.Fatalf("Expected '%s' but got '%s'", given, i.State)
	}
	if i.Description != uponReceiving {
		t.Fatalf("Expected '%s' but got '%s'", uponReceiving, i.Description)
	}
	if i.Operation != operationName {
		t.Fatalf("Expected '%s' but got '%s'", operationName, i.Operation)
	}
	if i.Query != query {
		t.Fatalf("Expected '%s' but got '%s'", query, i.Query)
	}
}
