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
	callback    func(*config) error
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
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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
		locationAreaResponse, err = pokeapi.UnmarshalLocationAreaResponse(value, nil)
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

func commandMapb(cfg *config) error {
	var url string
	var locationAreaResponse pokeapi.LocationAreaResponse
	var err error
	var data []byte
	// determine the url we are working with
	if cfg.Previous == nil {
		fmt.Print("you're on the first page\n")
		return nil
	} else {
		url = *cfg.Previous
	}

	// check if the current url is in the cache
	value, ok := cfg.Cache.Get(url)
	if ok {
		locationAreaResponse, err = pokeapi.UnmarshalLocationAreaResponse(value, nil)
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
