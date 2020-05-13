package token

// Type : Represents the type of a token literal
//        for performance take it could be an int
//        but for debuging reasons is betters
type Type string

// Token : Represents the structure of a token by the lexer
type Token struct {
	Type    Type
	Literal string
	// Line int
	// Column int
	// FileName string
}

const (
	ILLEGAL = "ILLEGAL" // ilegal character
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y...
	INT   = "INT"   // 42

	// Identifiers + literals
	ASSIGN = "="
	PLUS   = "+"

	// Operators
	COMMA     = ","
	SEMICOLON = ";"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
