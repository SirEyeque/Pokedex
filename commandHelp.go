package main

import(
	"fmt";
)


func commandHelp() error{
	note := `
Enter in any of the following commands to the Pokedex commandline

Exit
Help

`
	fmt.Print(note)
	return nil
}


