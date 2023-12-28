package main

import (
	"fmt"
	"log"
	"os"
	"github.com/KisuraWSP/belle/repl"
	"github.com/KisuraWSP/belle/repl/doc"
	"github.com/KisuraWSP/belle/ast"
)

// version 
var VERSION_NO string = fmt.Sprintf("%v",0.1)
var VERSION string = "Belle Version " + VERSION_NO

// help <- guides the users to use the program
var HELP string =`
Belle is a tool for managing belle source code.
        	
Usage:
        belle <command> [arguments]
	
	The commands are:
     	  	run		compile and run the program via the executable 
        	version		displays the version of the language 
        	doc		displays the REPL for the language documentation
		repl 		loads the READ-EVAL-PRINT-LOOP
		fread 		reads the .belle file if given as an argument
		values 		displays the values for the enums in the language
			
`

func main() {
	args := os.Args
	for i := 1; i < len(args); i++{
		switch args[i]{
			case "run":
				fmt.Println("TO BE IMPLEMENTED")		
			case "version":
				fmt.Printf("%s\n",VERSION)
			case "doc":
				doc.Load()
			case "repl":
				repl.Load()
			case "values":
				ast.Print_reserved_values()
				ast.Print_primitive_values()
				ast.Print_symbol_values()
			case "fread":
				fmt.Println("TO BE IMPLEMENTED")
			case "help":
				fmt.Printf("%s\n",HELP)		
			default:
				log.Println("Not an valid  argument")	
			
		}
	}
}
