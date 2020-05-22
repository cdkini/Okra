package main

type TokenType int

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   string
	line      int
}

const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Start
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	Identifier
	String
	Numeric
	And
	Class
	Else
	False
	Func
	For
	If
	Null
	Or
	Log
	Return
	Super
	This
	True
	Variable
	EndOfFile
)
