package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func CleanInput(test string) []string {
	out := strings.Fields(strings.ToLower(test))

	return out
}

func processCommand(input string) {
	userInput := CleanInput(input)
	cmdMap := getCmdMap()
	if len(userInput) > 0 {
		if val, ok := cmdMap[userInput[0]]; ok {
			val.call()
		} else {
			fmt.Println("Unknown command")
		}
	} 	
}

func REPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for ; ; {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Could not scan user input properly")
		}
		processCommand(scanner.Text())
	}
}
