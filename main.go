package main

import (
	"fmt";
	"bufio";
	"os"; 
	"strings";
)

type cliCommand struct{
	name			string
	description		string
	callback		func() error
}

func main() {
	
	cmds := map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	var usrInput string

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			usrInput = strings.ToLower(scanner.Text())
			usrInput = strings.TrimSpace(usrInput)
		}

		_, ok := cmds[usrInput]
		if !ok {
			fmt.Printf("\nCommand '%s' does not exist. Enter 'man ' for a list of possible commands\n", usrInput)
		}else{
			if err :=  cmds[usrInput].callback(); err != nil{
				fmt.Printf("%s", err)
				break
			}
		}
	}
}
