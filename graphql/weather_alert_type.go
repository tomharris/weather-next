package graphql

import (
	"github.com/graphql-go/graphql"
)

var weatherAlertType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WeatherAlert",
		Fields: graphql.Fields{
			"sender": &graphql.Field{
				Type: graphql.String,
			},
			"event": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"start": &graphql.Field{
				Type: graphql.Int,
			},
			"end": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
