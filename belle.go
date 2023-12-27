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

// reserved keywords and symbols
const (
	belle int = iota 
	start
	end 
	global_declare
	impl
	belle_proc
	belle_func
	bPrint
	bPrintln
	bFprint
	bFprintln
	bEprint
	bEprintln
			
	// types
	bInt4 // signed 4-bit integer
	bInt8 // signed 8-bit integer
	bInt16 // signed 16-bit integer
	bInt32 // signed 32-bit integer
	bUint4 // unsigned 4-bit integer
	bUint8 // unsigned 8-bit integer
	bUint16 // unsigned 16-bit integer
	bUin32 // unsigned 32-bit integer
	bFloat // signed float
	bUfloat // unsigned float 
	bDouble // signed double
	bUdouble // unsigned double 
	bLong // signed long (64-bit)
	bUong // unsigned long (64-bit)
	bString // string
	bChar // char
	bAny // Any type
	bError // Error type
	

	// Symbols
	PLUS // +
	MINUS // -
	SLASH_L // /
	SLASH_R // \ //
	STAR // *
	SINGLEQUOTE // '
	DOUBLEQUOTE // "
	BRACKET_L // (
	BRACKET_R // )
	EQUAL // =
	NOT // !
	AND // &
	OR // |
	SQUARE_BRACKET_L // [
	SQUARE_BRACKET_R // ]
	COLON // :
	COMMA // ,
	DOT // .
)

// print enums
var reserved = []int{
	// basic keywords
	belle,
	start,
	end,
	global_declare,
	impl,
	belle_proc,
	belle_func,
	bPrint,
	bPrintln,
	bFprint,
	bFprintln,
	bEprint,
	bEprintln,
}

var primitives = []int{
	// types
	bInt4, 
	bInt8, 
	bInt16, 
	bInt32, 
	bUint4, 
	bUint8, 
	bUint16, 
	bUin32, 
	bFloat, 
	bUfloat, 
	bDouble, 
	bUdouble, 
	bLong, 
	bUong, 
	bString,
	bChar, 
	bAny,
	bError,
}

var symbols = []int{
	// Symbols
	PLUS, 
	MINUS, 
	SLASH_L, 
	SLASH_R,
	STAR, 
	SINGLEQUOTE,
	DOUBLEQUOTE, 
	BRACKET_L, 
	BRACKET_R, 
	EQUAL, 
	NOT, 
	AND, 
	OR, 
	SQUARE_BRACKET_L, 
	SQUARE_BRACKET_R, 
	COLON, 
	COMMA, 
	DOT,
}

func print_reserved_values(){
	var length int = len(reserved)
	fmt.Println("VALUES {RESERVED}")
	for i:=0; i<length; i++ {
		fmt.Println(i)
	}
}

func print_symbol_values(){
	var length int = len(symbols)
	fmt.Println("VALUES {SYMBOLS}")
	for i:=0; i<length; i++ {
		fmt.Println(i)
	}
}

func print_primitive_values(){
	var length int = len(primitives)
	fmt.Println("VALUES {PRIMITIVES}")
	for i:=0; i<length; i++ {
		fmt.Println(i)
	}
}

// documentation
func docrepl_load() int {
	var input string 
	fmt.Print("DOC=>")
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit" {
			log.Println("Exiting DOC REPL")
			break
		} else if input == "test" {
			content, err := os.ReadFile("test.txt")
			if err != nil {
				log.Fatal(err)
			}
			str := string(content)
			fmt.Println(str)
			break
		} else {
			fmt.Println("Invalid Argument Entered!")
			docrepl_load()
		}

	}
	return 1
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
	order int // used for the expression generation to determine sequence order
	content *tree // not sure need to do more research
	type_ int // used to identify whether if it is a reserved keyword or not
}

func (t *token) check_if_reserved(){
	// if the keyword is reserved then do some operation with it
	// if not its a value by default to that operations type 
}


func generate_expression(tokenlist []*token){
	//generate a valid expression and make sure it works with the programming languages logic 	
}

// syntax analysis (expression parser)


// semantic analysis (syntax verification)


// REPL (READ-EVALUATE-PRINT-LOOP)
func repl_load() int {
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
	return 1
}

func main() {
	var run_flag int
	args := os.Args
	for i := 1; i < len(args); i++{
		if args[i] == "run" {
			// run the source code
		}else if args[i] == "version" {
			fmt.Printf("%s\n",VERSION)
			break
		}else if args[i] == "doc" {
			ret := docrepl_load()
			run_flag = ret
			break
		}else if args[i] == "repl" {
			ret := repl_load()
			run_flag = ret
			break
		}else if args[i] == "fread" {
			// read the .belle file
			break
		}else if args[i] == "values" {
			print_reserved_values()
			print_primitive_values()
			print_symbol_values()
			break
		}else {
			log.Println("Not an valid  argument")
		}
	}
	if run_flag != 1 {
		fmt.Printf("%s",HELP)
	}
}
