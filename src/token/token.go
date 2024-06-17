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
	UNKNOWN    = "UNKNOWN"
	LITERAL    = "LITERAL"
	COMMENT    = "COMMENT"
	WHITESPACE = "WHITESPACE"

	// reserved keywords
	LetToken       = "let"
	ImportToken    = "import"
	LoadToken      = "load"
	MatchToken     = "match"
	CaseToken      = "case"
	ElseToken      = "else"
	OrElseToken    = "orelse"
	ReturnToken    = "return"
	PrintToken     = "print"
	ProcedureToken = "proc"
	EndToken       = "end"

	// Operators
	PLUS               = "+"
	MINUS              = "-"
	STAR               = "*"
	LSLASH             = "/"
	ASSIGN             = "="
	EQUAL              = "=="
	NOT                = "!"
	OR                 = "||"
	AND                = "&&"
	GREATER_THAN       = ">"
	LESS_THAN          = "<"
	GREATER_THAN_EQUAL = ">="
	LESS_THAN_EQUAL    = "<="

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
)

var keywords = map[string]TokenType{
	"let":    LetToken,
	"proc":   ProcedureToken,
	"import": ImportToken,
	"load":   LoadToken,
	"match":  MatchToken,
	"case":   CaseToken,
	"else":   ElseToken,
	"orelse": OrElseToken,
	"return": ReturnToken,
	"end":    EndToken,
	"print":  PrintToken,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
