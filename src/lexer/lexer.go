package lexer

import (
	"strings"
)

/*
	ENUMS
*/

type Keyword int
type Type int
type Operator int
type Seperator int
type Literal int
type Identifier int
type Comment int
type Whitespace int
type Symbols int

const (
	// reserved keywords
	LetToken Keyword = iota
	MutToken
	ImportToken
	LoadToken
	MatchToken

	// types
	Int      Type = iota // signed integer (64 bit)
	Flt                  // Float (64 bit)
	String               // string
	Bool                 // Boolean
	Chr                  // char
	UInt                 // Unsigned integer (64 bit) type
	Any                  // Any type
	Function             // Function type
	Record               // Record type

	// Operators
	PLUS               Operator = iota // +
	MINUS                              // -
	STAR                               // *
	SLASH_L                            // /
	EQUAL                              // =
	NOT                                // !
	OR                                 // ||
	AND                                // &&
	GREATER_THAN                       // >
	LESS_THAN                          // <
	GREATER_THAN_EQUAL                 // >=
	LESS_THAN_EQUAL                    // <=
	WALRUS                             // :=

	// Symbols
	SLASH_R     Symbols = iota // \ //
	SINGLEQUOTE                // '
	DOUBLEQUOTE                // "
	ARROW                      // ->
	COLON_DUB                  // ::
	HASH                       // #
	AT                         // @
	BACK_TICK                  // `

	// Seperator
	BRACKET_L Seperator = iota // (
	BRACKET_R                  // )
	SEMICOLON                  // ;
	COLON                      // :
	COMMA                      // ,
	DOT                        // .
	SQUIRLY_L                  // {
	SQUIRLY_R                  // }
)

// lexical analysis (generate tokens for each keyword or expression evaluated)
type TokenPair struct {
	Kind int // checks the type of the token
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

func Lexer(str string) []TokenPair {
	var tokenList []TokenPair
	for len(str) > 0 {
		if strings.Index(str, "(") == 0 {
			tokenList = append(tokenList, TokenPair{int(BRACKET_L), "("})
		} else if strings.Index(str, "+") == 0 {
			tokenList = append(tokenList, TokenPair{int(BRACKET_L), "+"})
		} else if strings.Index(str, "-") == 0 {
			tokenList = append(tokenList, TokenPair{int(BRACKET_L), "-"})
		} else if strings.Index(str, "/") == 0 {
			tokenList = append(tokenList, TokenPair{int(BRACKET_L), "/"})
		} else if strings.Index(str, "*") == 0 {
			tokenList = append(tokenList, TokenPair{int(BRACKET_L), "*"})
		} else {
		}
	}
	return tokenList
}
