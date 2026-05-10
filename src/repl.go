package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/sireyeque/Pokedex/internal/pokecache"
)

func CleanInput(test string) []string {
	out := strings.Fields(strings.ToLower(test))

	return out
}

func REPL() {
	var nav navURL
	loc_endpoint := "https://pokeapi.co/api/v2/location-area/" 
	pokemon_endpoint := "https://pokeapi.co/api/v2/pokemon/" 
	nav.Next = "https://pokeapi.co/api/v2/location-area/"
	nav.Prev = ""
	scanner := bufio.NewScanner(os.Stdin)
	num_seconds := 10
	dur := time.Duration(num_seconds) * time.Second
	c := pokecache.NewCache(dur)
	for ; ; {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Could not scan user input properly")
		}
		userInput := CleanInput(scanner.Text())
		cmdMap := getCmdMap()
		if len(userInput) > 0 {
			if val, ok := cmdMap[userInput[0]]; ok {
				if userInput[0] == "explore" && len(userInput) > 1 {
					ex_nav := navURL{}
					ex_nav.Next = loc_endpoint + userInput[1]
					ex_nav.Prev = nil
					val.call(c, &ex_nav)
				} else if userInput[0] == "catch" && len(userInput) > 1 {
					pok_nav := navURL{}
					pok_nav.Next = pokemon_endpoint + userInput[1]
					pok_nav.Prev = nil
					val.call(c, &pok_nav)
				} else {
					val.call(c, &nav)
				}
			} else {
				fmt.Println("Unknown command")
			}
		} 	
	}
}
