package main

import (
	"fmt"
	"os"
	"github.com/sireyeque/Pokedex/internal/apireq"
	"github.com/sireyeque/Pokedex/internal/pokecache"
)

type navURL struct{
		Next string
		Prev any
	}

type cliCmd struct {
	name string
	desc string
	call func(c *pokecache.Cache, nav *navURL) error
}

func getCmdMap() map[string]cliCmd {
	return map[string]cliCmd {
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex",
			call: commandExit,
		},
		"help": {
			name: "help",
			desc: "Displays a help message",
			call: commandHelp,
		},
		"map": {
			name: "map",
			desc: "Returns the next 20 map locations",
			call: commandMap,
		},
		"mapb": {
			name: "mapb",
			desc: "Returns the previous 20 map locations",
			call: commandMapB,
		},
	}
}

func commandMapB(c *pokecache.Cache, nav *navURL) error {
	if nav.Prev == nil {
		fmt.Printf("You're on the first page\n")
		return fmt.Errorf("You're on the first page\n")
	}
	data := apireq.RecieveLocArea(nav.Prev.(string))

	nav.Next = data.Next
	nav.Prev = data.Previous

	// set data from Get request
	for i := 0; i < len(data.Results); i++ {
		fmt.Printf("%s\n", data.Results[i].Name)
	}
	return nil
}

func commandMap(c *pokecache.Cache, nav *navURL) error {
	data := apireq.RecieveLocArea(nav.Next)

	nav.Next = data.Next
	nav.Prev = data.Previous

	// set data from Get request
	for i := 0; i < len(data.Results); i++ {
		fmt.Printf("%s\n", data.Results[i].Name)
	}
	return nil
}

func commandHelp(c *pokecache.Cache, nav *navURL) error {
	cmdMap := getCmdMap()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, c := range cmdMap {
		fmt.Printf("%v: %v\n", c.name, c.desc)
	}
	return nil
}

func commandExit(c *pokecache.Cache, nav *navURL) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
