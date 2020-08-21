package parse

import (
	"Okra/okra/interpreter/ast"
	"Okra/okra/okraerr"
	"fmt"
)

// A Parser evaluates a collection of tokens and constructs abstract syntax trees (ASTs) out of
// the resulting expressions and statements. It is also responsible for consolidating parse errors
// and providing useful feedback to the user.
type Parser struct {
	tokens   []ast.Token
	current  int
	hadError bool
}

func NewParser(tokens []ast.Token) *Parser {
	return &Parser{tokens, 0, false}
}

// Parse triggers the recursive descent parsing, checking the given tokens and their state against
// the grammar rules of the language. Upon reaching a terminal, the function adds an instance of
// the appropriate statement to a resulting statement slice.
//   Args: nil
//   Returns: Slice of statements to interpret
//            Bool that tracks whether or not a parse error occurred
func (p *Parser) Parse() ([]ast.Stmt, bool) {
	stmts := []ast.Stmt{}

	for !p.isAtEOF() {
		p.parse(&stmts)
	}

	return stmts, p.hadError
}

func (p *Parser) parse(stmts *[]ast.Stmt) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Parse Error:", r.(error))
			p.hadError = true
			p.synchronize()
		}
	}()

	stmt := p.declaration()
	*stmts = append(*stmts, stmt)
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

func (p *Parser) consume(t ast.TokenType, msg string) ast.Token {
	if !p.check(t) {
		curr := p.currToken()
		okraerr.ReportErr(curr.Line, curr.Col, msg)
	}
	return p.advance()
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEOF() {
		if p.prevToken().Type == ast.Semicolon {
			return
		}

		switch p.peek().Type {
		case ast.Class,
			ast.Func,
			ast.Variable,
			ast.For,
			ast.If,
			ast.Print,
			ast.Return:
			return
		}

		p.advance()
	}
}
