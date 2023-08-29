package main

import (
	"fmt"
	"nasa/api"
)

const BaseUrl = "https://api.nasa.gov"
const ApiKey = "DEMO_KEY"

func main() {

	client := api.NewApiClient(BaseUrl, ApiKey)
	var count int
	date := "2023-08-28"
	// count := 3

	if count != nil {
		apod, err := api.GetApod[[]api.ApodResponse](client, nil, &count)
		if err != nil {
			panic(err)
		}

		for _, pic := range *apod {
			str := fmt.Sprintf("URL: %s", pic.Url)
			println(str)
		}

	} else {
		apod, err := api.GetApod[api.ApodResponse](client, &date, nil)
		if err != nil {
			panic(err)
		}

		str := fmt.Sprintf("URL: %s", apod.Url)
		println(str)
	}
}
