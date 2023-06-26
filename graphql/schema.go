package graphql

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"

	weathernext "github.com/tomharris/weather-next"
)

type Schema struct {
	gqlSchema graphql.Schema
}

type QueryData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func NewSchema(locationWeatherService weathernext.LocationWeatherService) (Schema, error) {
	gqlSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQueryRoot(locationWeatherService),
		},
	)

	if err != nil {
		return Schema{}, err
	}

	return Schema{gqlSchema: gqlSchema}, nil
}

func (s Schema) ExecuteQuery(ctx context.Context, query QueryData) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Context:        ctx,
		Schema:         s.gqlSchema,
		RequestString:  query.Query,
		VariableValues: query.Variables,
		OperationName:  query.Operation,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
