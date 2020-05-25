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
