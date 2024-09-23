package main

import (
	"fmt";
	"bufio";
	"os"; 
	"strings";
)



/*
*	structure: 	cliCmd
*
*	purpose:	stores information about each command and the 
*				function to be called for each command
*/
type cliCmd struct{
	name			string
	description		string
	callback		func() error
}


/*
* 	funciton: 	getCommands()
*
* 	inputs: 	None
*
* 	outputs: 	Map containing string : cliCmd (struct) key value pairs
*
*	purpose:	Returns the list of possible commands and important
*				information on how to use each command
*/
func getCommands() map[string]cliCmd {
	return map[string]cliCmd{
		"help": {
			name:			"help",
			description:	"Provides helpful information on how to use each command.",
			callback: 		commandHelp,
		},
		"exit": {
			name:			"exit",
			description:	"Exits the Pokedex command line interface and returns you to your terminal.",
			callback: 		commandExit,
		},
	}
}


/*
*	function: 	REPL()
*
*	inputs: 	None
*
*	outputs: 	None
*
*	purpose:	Runs the main interface between the user and the
* 				command line interface.
*/
func REPL() {

	scanner := bufio.NewScanner(os.Stdin)

	var usrInput string

	cmds := getCommands()

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
