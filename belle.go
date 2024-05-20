package main

import (
	"fmt"
	"log"
	"os"
)

// version
var VERSION_NO string = "0.1.5 alpha build 0"
var VERSION string = "Belle Version " + VERSION_NO

// help <- guides the users to use the program
var HELP string = `
Belle is a tool for managing belle source code.
        	
Usage:
        belle <command> [arguments]
	
	The commands are:
     	  	run		compile and run the program via the executable 
        	version		displays the version of the language 
			
`

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("%s\n", HELP)
	}
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "run":
			fmt.Println("TO BE IMPLEMENTED")
		case "version":
			fmt.Printf("%s\n", VERSION)
		case "help":
			fmt.Printf("%s\n", HELP)
		default:
			log.Fatalf("Not an valid  argument")
		}
	}
}
