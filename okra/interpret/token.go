package interpret

import "fmt"

// A Token is a substring of source text given context or meaning by the scanner
type Token struct {
	tokenType TokenType
	lexeme    string // Exact substring lexed by scanner
	literal   interface{}
	line      int
	col       int
}

// Only used for testing purposes within interpret_test package
func NewToken(tokenType TokenType, literal interface{}) *Token {
	return &Token{tokenType: tokenType, literal: literal}
}

func (t *Token) TokenType() TokenType {
	return t.tokenType
}

type TokenType int

func (t Token) String() string {
	return fmt.Sprintf("%v", keywordSlice[t.tokenType])
}

const (
	And TokenType = iota
	Bang
	BangEqual
	Class
	Comma
	Construct
	Dot
	Else
	EOF
	Equal
	EqualEqual
	False
	For
	Func
	Greater
	GreaterEqual
	Identifier
	If
	Invalid
	LeftBrace
	LeftBracket
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
	RightBracket
	RightParen
	Semicolon
	Slash
	Star
	String
	Struct
	Super
	This
	True
	Variable
	While
)

// Used for debugging purposes (returns the str representation as opposed to the enum val)
var keywordSlice = []string{
	"And",
	"Bang",
	"BangEqual",
	"Class",
	"Comma",
	"Construct",
	"Dot",
	"Else",
	"EOF",
	"Equal",
	"EqualEqual",
	"False",
	"For",
	"Func",
	"Greater",
	"GreaterEqual",
	"Identifier",
	"If",
	"Invalid",
	"LeftBrace",
	"LeftBracket",
	"LeftParen",
	"Less",
	"LessEqual",
	"Minus",
	"Null",
	"Numeric",
	"Or",
	"Plus",
	"Print",
	"Return",
	"RightBrace",
	"RightBracket",
	"RightParen",
	"Semicolon",
	"Slash",
	"Star",
	"String",
	"Struct",
	"Super",
	"This",
	"True",
	"Variable",
	"While",
}

var keywordDict = map[string]TokenType{
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"func":   Func,
	"if":     If,
	"null":   Null,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Variable,
	"while":  While,
}
