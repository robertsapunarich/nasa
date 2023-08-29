package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NasaApiClient struct {
	BaseUrl string
	ApiKey  string
	Client  *http.Client
}

type Resp interface {
	ApodResponse | []ApodResponse
}

func NewApiClient(baseUrl, apiKey string) *NasaApiClient {
	return &NasaApiClient{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
		Client:  &http.Client{},
	}
}

func GetApod[R Resp](nasaApiClient *NasaApiClient) (*R, error) {
	url := fmt.Sprintf("%s/planetary/apod?api_key=%s", nasaApiClient.BaseUrl, nasaApiClient.ApiKey)
	resp, err := nasaApiClient.Client.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apodResp R

	err = json.NewDecoder(resp.Body).Decode(&apodResp)

	if err != nil {
		return nil, err
	}

	return &apodResp, nil
}
