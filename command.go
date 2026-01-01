package main

import (
	"fmt"
	"os"

	"github.com/Seva-Sh/pokedexgo/internal/pokeapi"
	"github.com/Seva-Sh/pokedexgo/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the 20 location names",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location names",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display a list of Pokemon in a given location",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, location *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, location *string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func commandMap(cfg *config, location *string) error {
	var url string
	var locationAreaResponse pokeapi.LocationAreaResponse
	var err error
	var data []byte
	// determine the url we are working with
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = pokeapi.LocationAreaURL
	}

	// check if the current url is in the cache
	value, ok := cfg.Cache.Get(url)
	if ok {
		locationAreaResponse, err = pokeapi.UnmarshalLocationAreaResponse(value)
	} else {
		locationAreaResponse, err, data = pokeapi.GetLocationAreaResponse(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, data)
	}

	if err != nil {
		return err
	}

	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	for _, area := range locationAreaResponse.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config, location *string) error {
	var url string
	var locationAreaResponse pokeapi.LocationAreaResponse
	var err error
	var data []byte
	// determine the url we are working with
	if cfg.Previous == nil {
		fmt.Print("you're on the first page\n")
		return nil
	}
	url = *cfg.Previous

	// check if the current url is in the cache
	value, ok := cfg.Cache.Get(url)
	if ok {
		locationAreaResponse, err = pokeapi.UnmarshalLocationAreaResponse(value)
	} else {
		locationAreaResponse, err, data = pokeapi.GetLocationAreaResponse(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, data)
	}

	if err != nil {
		return err
	}

	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	for _, area := range locationAreaResponse.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *config, location *string) error {
	if location == nil {
		fmt.Println("You must provide a location name. Example: explore pastoria-city-area")
		return nil
	}
	var locationAreaNamedResponse pokeapi.LocationAreaNamedResponse
	var err error
	var data []byte
	url := pokeapi.LocationAreaURL + "/" + *location

	value, ok := cfg.Cache.Get(url)
	if ok {
		locationAreaNamedResponse, err = pokeapi.UnmarshalLocationAreaNamedResponse(value)
	} else {
		locationAreaNamedResponse, err, data = pokeapi.GetLocationAreaNamedResponse(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, data)
	}

	if err != nil {
		return err
	}

	fmt.Println("Exploring " + *location + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaNamedResponse.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
