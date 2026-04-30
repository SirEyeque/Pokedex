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
				val.call(c, &nav)
			} else {
				fmt.Println("Unknown command")
			}
		} 	
	}
}
