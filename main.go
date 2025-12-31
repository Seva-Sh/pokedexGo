package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Seva-Sh/pokedexgo/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	interval := 5 * time.Second
	cfg := &config{}
	cfg.Cache = pokecache.NewCache(interval)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()

		cleanedUserInput := cleanInput((userInput))
		if len(cleanedUserInput) == 0 {
			continue
		}

		commands := getCommands()

		if value, ok := commands[cleanedUserInput[0]]; ok {
			err := value.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		// fmt.Printf("Your command was: %v\n", cleanedUserInput[0])
	}
}
