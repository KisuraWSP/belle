package ast

import(
	"fmt"
)

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

func Print_reserved_values(){
	var length int = len(reserved)
	fmt.Println("VALUES {RESERVED}")
	for i:=0; i<length; i++ {
		fmt.Println(reserved[i])
	}
}

func Print_symbol_values(){
	var length int = len(symbols)
	fmt.Println("VALUES {SYMBOLS}")
	for i:=0; i<length; i++ {
		fmt.Println(symbols[i])
	}
}

func Print_primitive_values(){
	var length int = len(primitives)
	fmt.Println("VALUES {PRIMITIVES}")
	for i:=0; i<length; i++ {
		fmt.Println(primitives[i])
	}
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

func (t *tree) Walk_left(){
	// traverse the left side of the tree
}

func (t *tree) Walk_right(){
	// traverse the right side of the tree
}

func (t *tree) Check_root(){
	// check if the node is the root
}

func (t *tree) Append_node(n *node){
	// append node to the tree
}

func (t *tree) Check_node_type(n *node){
	// check the type of the node with the tree node
}
