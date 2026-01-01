package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Seva-Sh/pokedexgo/internal/pokecache"
)

func main() {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	interval := 5 * time.Second
	cfg := &config{}
	cfg.Cache = pokecache.NewCache(interval)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			return
		}

		userInput := scanner.Text()

		cleanedUserInput := cleanInput((userInput))
		if len(cleanedUserInput) == 0 {
			continue
		}

		if value, ok := commands[cleanedUserInput[0]]; ok {
			if len(cleanedUserInput) == 2 {
				extraParam := &cleanedUserInput[1]
				err = value.callback(cfg, extraParam)
			} else {
				err = value.callback(cfg, nil)
			}

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		// fmt.Printf("Your command was: %v\n", cleanedUserInput[0])
	}
}
