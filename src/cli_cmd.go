package main

import (
	"fmt"
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
	}
}

func commandHelp() error{
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
