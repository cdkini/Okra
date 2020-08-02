package interpret

import "fmt"

type Parser struct {
	tokens   []Token
	current  int
	hadError bool
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens, 0, false}
}

func (p *Parser) Parse() ([]Stmt, bool) {
	stmts := []Stmt{}

	for !p.end() {
		p.parse(&stmts)
	}

	return stmts, p.hadError
}

func (p *Parser) parse(stmts *[]Stmt) {
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

// match returns true if the current token is one of types
func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.end() {
		return false
	} else {
		return p.peek().tokenType == t
	}
}

func (p *Parser) advance() Token {
	if !p.end() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) end() bool {
	return p.peek().tokenType == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(t TokenType, msg string) Token {
	if p.check(t) {
		return p.advance()
		// FIXME: Make OkraError instance
		// } else {
		// 	panic(&ParseError{p.peek(), msg})
	}
	return Token{} // TODO: Remove upon fixing method
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.end() {
		if p.previous().tokenType == Semicolon {
			return
		}

		switch p.peek().tokenType {
		case Class,
			Func,
			Variable,
			For,
			If,
			While,
			Print,
			Return:
			return
		}

		p.advance()
	}
}
