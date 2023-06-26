package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tomharris/weather-next/data"
	"github.com/tomharris/weather-next/graphql"
)

func main() {
	weatherService := data.WeatherService{
		ApiKey:      os.Getenv("WEATHER_API_KEY"),
		ApiEndpoint: os.Getenv("WEATHER_API_ENDPOINT"),
	}
	schema, err := graphql.NewSchema(weatherService)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var data graphql.QueryData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(400)
			return
		}

		result := schema.ExecuteQuery(r.Context(), data)
		json.NewEncoder(w).Encode(result)
	})

	port := os.Getenv("PORT")
	fmt.Println("Now server is running on port " + port)
	http.ListenAndServe(":8080", nil)
}
