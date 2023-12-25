package main

import (
	"fmt"
	"log"
	"os"
)

/*
# help guideline
("how we going to structure the `help` flag when called in the program")
("based on go help")

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

// version 
var VERSION_NO string = fmt.Sprintf("%v",0.1)
var VERSION string = "simp_ version " + VERSION_NO

// help <- guides the users to use the program
var HELP string =`
simp_ is a tool for managing simp_ source code.
        	
Usage:
        simp_ <command> [arguments]
	
	The commands are:
        	run		compile and run the program via the executable 
        	version		displays the version of the language 
        	doc		displays the REPL for the language documentation
		repl 		loads the READ-EVAL-PRINT-LOOP
`

// reserved keywords
const(
	simp int = iota
	start_simp
	end_simp
	global
	impl
	comma
	simp_proc
	simp_func
	simp_class

	yes_mommy
	yes_daddy

	// types
	simp_int8
	simp_int16
	simp_int32
	simp_float
	simp_double
	simp_long
	simp_string
	simp_char
)

// documentation
func docrepl_load() {
	var input string 
	fmt.Print("DOC=>")
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit" {
			log.Println("Exiting DOC REPL")
			break
		} else if input == "std" {
			content, err := os.ReadFile("std.txt")
			if err != nil {
				log.Fatal(err)
			}
			str := string(content)
			fmt.Println(str)
		} else {
			fmt.Println("Invalid Argument Entered!")
			docrepl_load()
		}

	}
}

func  doc_handler(file string) {
	// reads file and returns related documentation
	// used for the case of multiple arguments
}

// ast (abstract syntax tree)
type node struct{
	operation_type string
	content string
}

type tree struct{
	root *node 
	left *node  
	right *node 
}

func (t *tree) walk_left(){
	// traverse the left side of the tree
}

func (t *tree) walk_right(){
	// traverse the right side of the tree
}

func (t *tree) check_root(){
	// check if the node is the root
}

func (t *tree) append_node(n *node){
	// append node to the tree
}

func (t *tree) check_node_type(n *node){
	// check the type of the node with the tree node
}

// lexical analysis (generate tokens for each keyword or expression evaluated)
type token struct{
        // used for the lexical analyzer
}


// syntax analysis (expression parser)


// semantic analysis (syntax verification)


// assembler (generate assembly for the language)


// linker (link the assembly to create an executable file)


// REPL (READ-EVALUATE-PRINT-LOOP)
func repl_load(){
	var input string
	fmt.Print("RUN=>")
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit"{
			log.Print("Exiting Out of REPL")
			break
		}else if input == "version"{
			fmt.Printf("%s\n",VERSION)
			repl_load()
		}
	}

	
}

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {				               
		if args[i] == "help" {
			fmt.Printf("%s",HELP)
		}else if args[i] == "run" {
			// run the source code
		}else if args[i] == "version" {
			fmt.Printf("%s\n",VERSION)
			break
		}else if args[i] == "doc" {
			docrepl_load()
		}else if args[i] == "repl" {
			repl_load()
		}else {
			log.Println("Not an valid  argument")
		}
	}
}
