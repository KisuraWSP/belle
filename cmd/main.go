package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		if args[1] == "version" {
			version()
		}
		if args[i] == "help" {
			help()
		} else if args[i] == "run" {
			// run the source code
		} else if args[i] == "doc" {
			//d := &doc{}
			//d.repl_load()
		} else {
			log.Println("Not an valid  argument")
		}
	}
}
