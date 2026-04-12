package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCmd struct {
	name string
	desc string
	call func() error
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
	}
}

func commandMap() error {
	type locArea struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous any    `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err == nil{
		fmt.Errorf("%v", err)
	}

	body, err := io.ReadAll(resp.Body)
	data := locArea{}
	json.Unmarshal(body, &data)
	if err == nil {
		fmt.Errorf("%v", err)
	}
	for i := 0; i < len(data.Results); i++ {
		fmt.Printf("%s\n", data.Results[i].Name)
	}
	return nil
}

func commandHelp() error {
	cmdMap := getCmdMap()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, c := range cmdMap {
		fmt.Printf("%v: %v\n", c.name, c.desc)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
