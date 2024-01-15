package lexer

import(
	"fmt"
)

/*
	ENUMS
*/
const (
	// reserved keywords
	start int = iota 
	end
	declare
	implement
	bProc
	bFunc
	bReturn
	bVar
	if_ // if
	then 
	else_ // else
	for_ // for
	call
	
	// functions (built-in)
	bPrint
	bScan
			
	// types
	bInt // signed integer (64 bit)
	bDecPt // signed float
	bString // string
	bBool // Boolean
	bChr // char
	bAny // Any type
	bError // Error type
	

	// Operators
	PLUS // +
	MINUS // -
	STAR // *
	SLASH_L // /
	EQUAL // =
	NOT // !
	OR // |
	AND // &
	TRIANGLE_BRACKET_L // <
	TRIANGLE_BRACKET_R  // >
	
	// Symbols
	SLASH_R // \ //
	SINGLEQUOTE // '
	DOUBLEQUOTE // "
	
	// Seperator
	BRACKET_L // (
	BRACKET_R // )
	SEMICOLON // ;
	COLON // :
	COMMA // ,
	DOT // .
)

// print enums
var reserved = []int{
	start,
	end,
	declare,
	implement,
	bProc,
	bFunc,
	bReturn,
	bVar,
	if_, 
	then, 
	else_, 
	for_,
	call,
}

var primitives = []int{
	bInt, 
	bDecPt, 
	bString,
	bBool,
	bChr, 
	bAny, 
	bError,
}

var symbols = []int{
	SLASH_R, 
	SINGLEQUOTE, 
	DOUBLEQUOTE, 
}

var operators = []int{
	PLUS, 
	MINUS, 
	STAR, 
	SLASH_L, 
	EQUAL, 
	NOT, 
	OR,
	AND, 
	TRIANGLE_BRACKET_L, 
	TRIANGLE_BRACKET_R,
};

var seperators = []int{
	BRACKET_L, 
	BRACKET_R, 
	SEMICOLON, 
	COLON, 
	COMMA, 
	DOT, 
};

var builtin_func = []int{
	bPrint,
	bScan,
}; 

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

func Print_operator_values(){
	var length int = len(operators)
	fmt.Println("VALUES {OPERATORS}")
	for i:=0; i<length; i++ {
		fmt.Println(operators[i])
	}
}

func Print_builtin_func_values(){
	var length int = len(builtin_func)
	fmt.Println("VALUES {BUILTIN_FUNC}")
	for i:=0; i<length; i++ {
		fmt.Println(builtin_func[i])
	}
}

func Print_seperator_values(){
	var length int = len(seperators)
	fmt.Println("VALUES {SEPERATORS}")
	for i:=0; i<length; i++ {
		fmt.Println(seperators[i])
	}
}


func Load(){
	Print_seperator_values()
	Print_builtin_func_values()
	Print_operator_values()
	Print_primitive_values()
	Print_symbol_values()
	Print_reserved_values()
}

// lexical analysis (generate tokens for each keyword or expression evaluated)
type token struct{
	type_ string // checks the type of the token 
		     /*
		     	Possible types include:
		     		identifier - defined by the programmer
		     		keyword - reserved by the langauge
		     		seperator - ; , : , " , " , { , } , ( , )
		     		operator - + , > , =
		     		literal - "test", true, 2.232
		     		comment - discarded by the program
		     		whitespace - discarded by the program
		     */
	value string // represents what the value that the token is evaluating 
}

func (t *token) Check_if_keyword(){
	// if the keyword is reserved then do some operation with it
	// if not its a value by default to that operations type 
}

/*
	 TODO
	 check_if_identifier
	 check_if_seperator
	 check_if_literal
	 check_if_comment
	 check_if_whitespace
	 check_if_operator
*/

func Generate_expression(tokenlist []*token){
	//generate a valid expression and make sure it works with the programming languages logic
	/*
	we would want to express it like this as below
	      	eg:- x = 1;
	      	 [(identifier,x),(operator,=),(literal,1)]
	*/ 	
}
