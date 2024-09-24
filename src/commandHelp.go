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
func commandHelp(urls * config) error{
	note := `
Enter in any of the following commands to the Pokedex commandline.
Commands are NOT case sensitive.

Map: 	Prints out 20 map locations	
		Each subsequent Map call will print out the next 20 map locations
Mapb: 	Prints out the previous 20 map locations

Help: 	Prints a help message
Exit: 	Exits Pokedex and returns you to the terminal

`
	fmt.Print(note)
	return nil
}


