package lexer

/*
	ENUMS
*/
const (
	// reserved keywords
	Declare int = iota
	Implement
	Proc
	Return
	Var
	If
	Then
	Else
	For
	Call

	// types
	Int    // signed integer (64 bit)
	Flt    // signed float
	String // string
	Bool   // Boolean
	Chr    // char
	Any    // Any type
	Error  // Error type

	// Operators
	PLUS               // +
	MINUS              // -
	STAR               // *
	SLASH_L            // /
	EQUAL              // =
	NOT                // !
	OR                 // |
	AND                // &
	TRIANGLE_BRACKET_L // <
	TRIANGLE_BRACKET_R // >

	// Symbols
	SLASH_R     // \ //
	SINGLEQUOTE // '
	DOUBLEQUOTE // "

	// Seperator
	BRACKET_L // (
	BRACKET_R // )
	SEMICOLON // ;
	COLON     // :
	COMMA     // ,
	DOT       // .
)

// lexical analysis (generate tokens for each keyword or expression evaluated)
type Token struct {
	Type int // checks the type of the token
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
	Value string // represents what the value that the token is evaluating
}
