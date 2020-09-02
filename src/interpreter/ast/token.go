package ast

// A Token is a substring of source text given context or meaning by the scanner.
type Token struct {
	Type    TokenType   // Token enum
	Lexeme  string      // Exact piece of source code tokenized
	Literal interface{} // Value associated with source code
	Line    int
	Col     int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int, col int) *Token {
	return &Token{tokenType, lexeme, literal, line, col}
}

type TokenType int

const (
	And TokenType = iota
	Bang
	BangEqual
	Colon
	Comma
	Dot
	Else
	EOF
	Equal
	False
	For
	Func
	Greater
	GreaterEqual
	Identifier
	If
	Invalid
	LeftBrace
	LeftParen
	Less
	LessEqual
	Minus
	Null
	Numeric
	Or
	Plus
	Print
	Return
	RightBrace
	RightParen
	Semicolon
	Slash
	Star
	String
	Struct
	This
	True
	Variable
)

// Used to differentiate between user defined object declaration and built-in keywords.
var KeywordDict = map[string]TokenType{
	"struct": Struct,
	"else":   Else,
	"false":  False,
	"for":    For,
	"func":   Func,
	"if":     If,
	"null":   Null,
	"print":  Print,
	"return": Return,
	"this":   This,
	"true":   True,
	"var":    Variable,
}
