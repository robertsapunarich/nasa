package api

type ApodResponse struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	MediaType      string `json:"mediaType"`
	ServiceVersion string `json:"serviceVersion"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}
