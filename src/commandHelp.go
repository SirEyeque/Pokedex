package main

import(
	"fmt"
)


/*
*	function:	commandHelp()
*
*	output:		error
*
*	purpose:	Prints a statement to the terminal on how to use
*				the Pokedex commands
*/
func commandHelp(urls * config) error{
	
	cmds := getCommands()

	for _, item := range cmds{
		fmt.Printf("%s:\t\t%s\n", item.name, item.description)
	}

	return nil
}


