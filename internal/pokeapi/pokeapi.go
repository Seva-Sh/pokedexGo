package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var LocationAreaURL = "https://pokeapi.co/api/v2/location-area"

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func UnmarshalLocationAreaResponse(data []byte, err error) (LocationAreaResponse, error) {
	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}

func ExtractURL(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("unsuccessful GET request")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, err
}

func GetLocationAreaResponse(url string) (LocationAreaResponse, error, []byte) {
	data, err := ExtractURL(url)
	if err != nil {
		return LocationAreaResponse{}, err, []byte{}
	}

	locationArea, err := UnmarshalLocationAreaResponse(data, err)
	return locationArea, err, data
}
