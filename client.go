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

func NewApiClient(baseUrl, apiKey string) *NasaApiClient {
	return &NasaApiClient{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
		Client:  &http.Client{},
	}
}

// Returns n random items from NASA's APOD collection where n is equal to the count provided.
func (c *NasaApiClient) GetApodRandom(count int) (*[]ApodResponse, error) {
	url := fmt.Sprintf("%s/planetary/apod?api_key=%s&count=%d", c.BaseUrl, c.ApiKey, count)

	resp, err := c.Client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var apodResp []ApodResponse
	err = json.NewDecoder(resp.Body).Decode(&apodResp)

	if err != nil {
		return nil, err
	}

	return &apodResp, nil
}

// Returns today's Astronomy Picture of the Day, or optionally the APOD for the date provided.
func (c *NasaApiClient) GetApod(date *string) (*ApodResponse, error) {
	url := fmt.Sprintf("%s/planetary/apod?api_key=%s", c.BaseUrl, c.ApiKey)

	if date != nil {
		url = url + fmt.Sprintf("&date=%s", *date)
	}

	resp, err := c.Client.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apodResp ApodResponse

	err = json.NewDecoder(resp.Body).Decode(&apodResp)

	if err != nil {
		return nil, err
	}

	return &apodResp, nil
}
