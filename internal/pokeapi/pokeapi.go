package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

// Location Area Methods
func UnmarshalLocationAreaResponse(data []byte) (LocationAreaResponse, error) {
	var locationAreaResponse LocationAreaResponse
	err := json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}

func GetLocationAreaResponse(url string) (LocationAreaResponse, error, []byte) {
	data, err := ExtractURL(url)
	if err != nil {
		return LocationAreaResponse{}, err, []byte{}
	}

	locationArea, err := UnmarshalLocationAreaResponse(data)
	return locationArea, err, data
}

// Location Area with Name Methods
func UnmarshalLocationAreaNamedResponse(data []byte) (LocationAreaNamedResponse, error) {
	var locationAreaNamedResponse LocationAreaNamedResponse
	err := json.Unmarshal(data, &locationAreaNamedResponse)
	if err != nil {
		return LocationAreaNamedResponse{}, err
	}

	return locationAreaNamedResponse, nil
}

func GetLocationAreaNamedResponse(url string) (LocationAreaNamedResponse, error, []byte) {
	data, err := ExtractURL(url)
	if err != nil {
		return LocationAreaNamedResponse{}, err, []byte{}
	}

	locationAreaNamed, err := UnmarshalLocationAreaNamedResponse(data)
	return locationAreaNamed, err, data
}

// Extracting Pokemon Methods
func UnmarshalPokemonResponse(data []byte) (Pokemon, error) {
	var pokemonResponse Pokemon
	err := json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResponse, nil
}

func GetPokemon(name string) (Pokemon, error) {
	url := PokemonURL + "/" + name
	data, err := ExtractURL(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon, err := UnmarshalPokemonResponse(data)
	return pokemon, err
}
