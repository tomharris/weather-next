package graphql

import (
	"github.com/graphql-go/graphql"
)

var locationWeatherType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LocationWeather",
		Fields: graphql.Fields{
			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
			"longitude": &graphql.Field{
				Type: graphql.Float,
			},
			"temperature": &graphql.Field{
				Type: graphql.Float,
			},
			"feelsLike": &graphql.Field{
				Type: graphql.Float,
			},
			"weatherType": &graphql.Field{
				Type: graphql.String,
			},
			"weatherDescription": &graphql.Field{
				Type: graphql.String,
			},
			"alerts": &graphql.Field{
				Type: graphql.NewList(weatherAlertType),
			},
		},
	},
)
