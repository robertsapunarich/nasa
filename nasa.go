package main

import (
	"fmt"
	"nasa/api"
)

const BaseUrl = "https://api.nasa.gov"
const ApiKey = "DEMO_KEY"

func main() {

	client := api.NewApiClient(BaseUrl, ApiKey)

	apod, err := api.GetApod[api.ApodResponse](client)

	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("URL: %s", apod.Url)
	println(str)
}
