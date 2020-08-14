package interpret

import "fmt"

// A Parser evaluates a collection of tokens and constructs abstract syntax trees (ASTs) out of
// the resulting expressions and statements. It is also responsible for consolidating parse errors
// and providing useful feedback to the user.
type Parser struct {
	tokens   []Token
	current  int
	hadError bool
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens, 0, false}
}

// Parse triggers the recursive descent parsing, checking the given tokens and their state against
// the grammar rules of the language. Upon reaching a terminal, the function adds an instance of
// the appropriate statement to a resulting statement slice.
//   Args: nil
//   Returns: Slice of statements to interpret
//            Bool that tracks whether or not a parse error occurred
func (p *Parser) Parse() ([]Stmt, bool) {
	stmts := []Stmt{}

	for !p.isAtEOF() {
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
	if p.isAtEOF() {
		return false
	} else {
		return p.peek().tokenType == t
	}
}

func (p *Parser) advance() Token {
	if !p.isAtEOF() {
		p.current++
	}

	return p.prevToken()
}

func (p *Parser) isAtEOF() bool {
	return p.peek().tokenType == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) currToken() Token {
	return p.tokens[p.current]
}

func (p *Parser) prevToken() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(t TokenType, msg string) Token {
	if !p.check(t) {
		curr := p.currToken()
		ReportErr(curr.line, curr.col, msg)
	}
	return p.advance()
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEOF() {
		if p.prevToken().tokenType == Semicolon {
			return
		}

		switch p.peek().tokenType {
		case Class,
			Func,
			Variable,
			For,
			If,
			Print,
			Return:
			return
		}

		p.advance()
	}
}
