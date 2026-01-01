package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var LocationAreaURL = "https://pokeapi.co/api/v2/location-area"

type LocationAreaNamedResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
