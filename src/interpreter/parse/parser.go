package parse

import (
	"github.com/cdkini/Okra/src/okraerr"

	"github.com/cdkini/Okra/src/interpreter/ast"
)

// A Parser evaluates a collection of tokens and constructs abstract syntax trees (ASTs) out of
// the resulting expressions and statements. It is also responsible for consolidating parse errors
// and providing useful feedback to the user.
type Parser struct {
	tokens   []ast.Token // As created by the scanner
	current  int
	hadError bool
}

func NewParser(tokens []ast.Token) *Parser {
	return &Parser{tokens, 0, false}
}

// Parse triggers the recursive descent parsing, checking the given tokens and their state against
// Okra's EBNF context free grammar. Upon reaching a terminal, the function adds an instance of
// the appropriate statement to a resulting statement slice.
//   Args: nil
//   Returns: Slice of statements to interpret
//            Bool that tracks whether or not a parse error occurred
func (p *Parser) Parse() []ast.Stmt {
	stmts := []ast.Stmt{}

	for !p.isAtEOF() {
		p.parse(&stmts)
	}

	return stmts
}

// parse is a helper method used in Parse to actually kick off the recursive descent parsing.
// It is responsible for syncronizing the parser in the case a parser error is met.
//   Args: stmts [*[]Stmt] - The resulting slice to be returned and interpreted upon completion of parsing.
//   Returns: nil
func (p *Parser) parse(stmts *[]ast.Stmt) {
	stmt := p.declaration()
	*stmts = append(*stmts, stmt)
}

// consume is a helper method used throughout the parser to iterate over and utilize the input token slice.
// If the token type argument does not match expectations, an error is raised. Otherwise, it is evaluated
// and the parser moves along to the next item.
//   Args: t   (ast.TokenType) - The token type the parser is expecting to evaluate
//         msg (string)        - The error message to be used in the raised error, if applicable
//   Returns: The next token in the sequence
//   Raises: OkraError if the token type does not match the parser's expectation
func (p *Parser) consume(t ast.TokenType, msg string) ast.Token {
	if !p.check(t) {
		curr := p.currToken()
		okraerr.ReportErr(curr.Line, curr.Col, msg)
	}
	return p.advance()
}

func (p *Parser) match(types ...ast.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t ast.TokenType) bool {
	if p.isAtEOF() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() ast.Token {
	if !p.isAtEOF() {
		p.current++
	}
	return p.prevToken()
}

func (p *Parser) isAtEOF() bool {
	return p.peek().Type == ast.EOF
}

func (p *Parser) peek() ast.Token {
	return p.tokens[p.current]
}

func (p *Parser) currToken() ast.Token {
	return p.tokens[p.current]
}

func (p *Parser) prevToken() ast.Token {
	return p.tokens[p.current-1]
}
