package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Token Types
	IDENT      = "INDENTIFIER"
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
	LITERAL    = "LITERAL"
	COMMENT    = "COMMENT"
	WHITESPACE = "WHITESPACE"

	// reserved keywords
	LetToken    = "let"
	MutToken    = "mut"
	ImportToken = "import"
	LoadToken   = "load"
	MatchToken  = "match"
	CaseToken   = "case"
	ElseToken   = "else"
	OrElseToken = "orelse"
	ReturnToken = "return"

	// types
	Int      = "Int"      // signed integer (64 bit)
	Flt      = "Flt"      // Float (64 bit)
	String   = "String"   // string
	Bool     = "Bool"     // Boolean
	Chr      = "Chr"      // char
	UInt     = "UInt"     // Unsigned integer (64 bit) type
	Any      = "Any"      // Any type
	Function = "Function" // Function type
	Record   = "Record"   // Record type
	Byte     = "Byte"     // Byte type
	Array    = "Array"    // Array type

	// Operators
	PLUS               = "+"
	MINUS              = "-"
	STAR               = "*"
	LSLASH             = "/"
	ASSIGN             = "="
	EQUAL              = "equal"
	NOT                = "not"
	OR                 = "or"
	AND                = "and"
	GREATER_THAN       = "greaterthan"
	LESS_THAN          = "lessthan"
	GREATER_THAN_EQUAL = "greaterthan_orequal"
	LESS_THAN_EQUAL    = "lessthan_orequal"
	WALRUS             = ":="

	// Symbols
	RSLASH      = `\`
	SINGLEQUOTE = `'`
	DOUBLEQUOTE = `"`
	ARROW       = `->`
	COLON_DUB   = `::`
	HASH        = `#`
	AT          = `@`
	BACK_TICK   = "`"

	// Seperator
	LBRACE    = "("
	RBRACE    = ")"
	SEMICOLON = ";"
	COLON     = ":"
	COMMA     = ","
	DOT       = "."
	LSQUIRLY  = "{"
	RSQUIRLY  = "}"
)

var keywords = map[string]TokenType{
	"let":      LetToken,
	"mut":      MutToken,
	"import":   ImportToken,
	"load":     LoadToken,
	"match":    MatchToken,
	"case":     CaseToken,
	"else":     ElseToken,
	"orelse":   OrElseToken,
	"return":   ReturnToken,
	"Int":      Int,      // signed integer (64 bit)
	"Flt":      Flt,      // Float (64 bit)
	"String":   String,   // string
	"Bool":     Bool,     // Boolean
	"Chr":      Chr,      // char
	"UInt":     UInt,     // Unsigned integer (64 bit) type
	"Any":      Any,      // Any type
	"Function": Function, // Function type
	"Record":   Record,   // Record type
	"Byte":     Byte,     // Byte type
	"Array":    Array,    // Array type
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
