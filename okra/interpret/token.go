package interpret

import "fmt"

type TokenType int

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t Token) String() string {
	return fmt.Sprintf("%v: %v |", keywordSlice[t.tokenType], t.lexeme)
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
	Log
	Minus
	Null
	Numeric
	Or
	Plus
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
)

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
	"Log",
	"Minus",
	"Null",
	"Numeric",
	"Or",
	"Plus",
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
}

var keywordDict = map[string]TokenType{
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"func":   Func,
	"if":     If,
	"null":   Null,
	"log":    Log,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Variable,
}
