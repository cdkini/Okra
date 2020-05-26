package main

type TokenType int

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
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

var keywordDict = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"func":   Func,
	"if":     If,
	"null":   Null,
	"or":     Or,
	"log":    Log,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Variable,
}
