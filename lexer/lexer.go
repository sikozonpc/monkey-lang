package lexer

// The Lexer is the first to read the source code and transform it into Token types

// Lexer :  Lexer data structure
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New : Creates a new lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
