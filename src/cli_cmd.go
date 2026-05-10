package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
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
		"explore": {
			name: "explore",
			desc: "Lists all the pokemon that appear in a specific location",
			call: commandExplore,
		},
		"catch": {
			name: "catch",
			desc: "Attempts to catch the pokemon specified in the command",
			call: commandCatch,
		},
	}
}

func commandCatch(c *pokecache.Cache, nav *navURL) error {
	body, ok := c.Get(nav.Next)
	if !ok {
		body = apireq.RecieveData(nav.Next)
		c.Add(nav.Next, body)
	}
	if data, ok := (apireq.PullData(body, "catch")).(apireq.PokemonData); ok {
		fmt.Printf("Throwing a Pokeball at %v...\n", data.Name, data.BaseExperience)
		val := rand.Intn(data.BaseExperience)
		time.Sleep(1*time.Second)
		if val < 20 {
			fmt.Printf("%v has been caught!!\n", data.Name)
		} else {
			fmt.Printf("%v has escaped...\n", data.Name)
		}
		return nil
	}
	fmt.Println("Incorrect URI given for Poke API request")
	return nil
}

func commandExplore(c *pokecache.Cache, nav *navURL) error {
	body, ok := c.Get(nav.Next)
	if !ok {
		body = apireq.RecieveData(nav.Next)
		c.Add(nav.Next, body)
	}
	if data, ok := (apireq.PullData(body, "explore")).(apireq.LocDetails); ok {
		if len(data.PokemonEncounters) == 0 {
			fmt.Println("Incorrect URI given for Poke API request")
		}
		for _, item := range data.PokemonEncounters {
			fmt.Println(item.Pokemon.Name)
		}
		return nil
	}
	fmt.Println("Incorrect URI given for Poke API request")
	return nil
}

func commandMapB(c *pokecache.Cache, nav *navURL) error {
	if nav.Prev == nil {
		fmt.Printf("You're on the first page\n")
		return fmt.Errorf("You're on the first page\n")
	}
	body, ok := c.Get(nav.Prev.(string))
	if !ok {
		body = apireq.RecieveData(nav.Prev.(string))
		c.Add(nav.Prev.(string), body)
	}
	if data, ok := (apireq.PullData(body, "map")).(apireq.LocArea); ok {
		nav.Next = data.Next
		nav.Prev = data.Previous

		// set data from Get request
		for i := 0; i < len(data.Results); i++ {
			fmt.Printf("%s\n", data.Results[i].Name)
		}
		return nil
	}
	fmt.Println("Incorrect URI given for Poke API request")
	return nil
}

func commandMap(c *pokecache.Cache, nav *navURL) error {
	body, ok := c.Get(nav.Next)
	if !ok {
		body = apireq.RecieveData(nav.Next)
		c.Add(nav.Next, body)
	}
	if data, ok := (apireq.PullData(body, "map")).(apireq.LocArea); ok {
		nav.Next = data.Next
		nav.Prev = data.Previous

		// set data from Get request
		for i := 0; i < len(data.Results); i++ {
			fmt.Printf("%s\n", data.Results[i].Name)
		}
		return nil
	}
	fmt.Println("Incorrect URI given for Poke API request")
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
