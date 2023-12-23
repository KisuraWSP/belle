package main

import (
	"fmt"
	"log"
	"os"
)

/*
# -- help guideline
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

*/

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		if args[1] == "version" {
			fmt.Println("simp_ version 0")
		}
		if args[i] == "help" {
			fmt.Println("simp_ is a tool for managing simp_ source code.\n")
			fmt.Println("Usage:\n")
			fmt.Println("simp <command> [arguments]\n")
			fmt.Println("The commands are:\n")
			fmt.Println("run\t\tcompile and run the program via the executable")
			fmt.Println("version\t\tdisplays the version of the language")
			fmt.Println("doc\t\tdisplays the REPL for the language documentation")
		} else if args[i] == "run" {
			// run the source code
		} else if args[i] == "doc" {
			//d := &doc{}
			//d.repl_load()
			test3()
		} else {
			log.Println("Not an valid  argument")
		}
	}
}
