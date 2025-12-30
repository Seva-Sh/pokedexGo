package pokeapi

import (
	"encoding/json"
	"fmt"
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

func GetResponseArea(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("unsuccessful GET request")
	}

	var locationAreaResponse LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	return locationAreaResponse, nil
}
