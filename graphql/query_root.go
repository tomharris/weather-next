package graphql

import (
	"github.com/graphql-go/graphql"

	weathernext "github.com/tomharris/weather-next"
)

func newQueryRoot(locationWeatherService weathernext.LocationWeatherService) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"locationWeather": &graphql.Field{
					Type:        locationWeatherType,
					Description: "Get weather for a location",
					Args: graphql.FieldConfigArgument{
						"latitude": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Float),
						},
						"longitude": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Float),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return locationWeatherResolver(locationWeatherService, p)
					},
				},
			},
		},
	)
}

func locationWeatherResolver(lws weathernext.LocationWeatherService, p graphql.ResolveParams) (interface{}, error) {
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	return lws.GetWeatherForLocation(latitude, longitude)
}
