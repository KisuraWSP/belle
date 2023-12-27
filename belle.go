package main

import (
	"fmt"
	"log"
	"os"
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
		fmt.Println(reserved[i])
	}
}

func print_symbol_values(){
	var length int = len(symbols)
	fmt.Println("VALUES {SYMBOLS}")
	for i:=0; i<length; i++ {
		fmt.Println(symbols[i])
	}
}

func print_primitive_values(){
	var length int = len(primitives)
	fmt.Println("VALUES {PRIMITIVES}")
	for i:=0; i<length; i++ {
		fmt.Println(primitives[i])
	}
}

// documentation
func docrepl_load() {
	var state int
	var input string 
	fmt.Print("DOC=>")
	state = 1
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit" {
			log.Println("Exiting DOC REPL")
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
			state = 0
			break
		}
	}
	if state != 1 {
		fmt.Print("Do You want to enter again (Y/n): ")
		fmt.Scan(&input)
		if input == "y" || input == "Y"{
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
func repl_load() {
	var state int
	var input string 
	fmt.Print("RUN=>")
	state = 1
	fmt.Scan(&input)
	for input != "exit" {
		if input == "exit" {
			log.Println("Exiting REPL")
			break
		} else if input == "version" {
			fmt.Printf("%s\n",VERSION)
			break
		} else {
			fmt.Println("Invalid Argument Entered!")
			state = 0
			break
		}
		repl_load()
	}
	if state != 1 {
		fmt.Print("Do You want to enter again (Y/n): ")
		fmt.Scan(&input)
		if input == "y" || input == "Y"{
			repl_load()
		}
	}
}

func main() {
	args := os.Args
	for i := 1; i < len(args); i++{
		switch args[i]{
			case "run":
				fmt.Println("TO BE IMPLEMENTED")		
			case "version":
				fmt.Printf("%s\n",VERSION)
			case "doc":
				docrepl_load()
			case "repl":
				repl_load()
			case "values":
				print_reserved_values()
				print_primitive_values()
				print_symbol_values()
			case "fread":
				fmt.Println("TO BE IMPLEMENTED")
			case "help":
				fmt.Printf("%s\n",HELP)		
			default:
				log.Println("Not an valid  argument")	
			
		}
	}
}
