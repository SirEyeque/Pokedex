package main

import(
	"fmt";
)


/*
*	function:	commandHelp()
*
*	output:		error
*
*	purpose:	Prints a statement to the terminal on how to use
*				the Pokedex commands
*/
func commandHelp() error{
	note := `

Enter in any of the following commands to the Pokedex commandline.
Commands are NOT case sensitive.

Help: Prints a help message

Exit: Exits Pokedex and returns you to the terminal
`
	fmt.Print(note)
	return nil
}


