package interpret

import (
	"fmt"
	"os"
)

// A Parser evaluates a collection of tokens to constructs an abstract syntax tree of expressions.
// It is also responsible for consolidating parse errors and providing useful feedback to the user.
type Parser struct {
	tokens []*Token
	curr   int
	errors []OkraError
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{tokens, 0, make([]OkraError, 0)}
}

// Parse evaluates a series of tokens using recursive descent, checking the given values against
// the grammar rules of the language. Upon reaching a terminal, the function ceases execution.
//   Args: nil
//   Returns: An instance of Expr that best fits the token stream
func (p *Parser) Parse() []Stmt {
	var statements []Stmt
	for p.getCurrToken().tokenType != EOF {
		statements = append(statements, p.statement())
	}
	return statements
}

func (p *Parser) varDeclaration() Stmt {
	identifier := p.consume(Identifier, "Expected variable name.")

	var expr Expr
	if p.match(Equal) {
		expr = p.expression()
	}
	p.consume(Semicolon, "Expected ';' at end of line.")
	return VariableStmt{identifier, expr}
}

func (p *Parser) statement() Stmt {
	if p.match(Variable) {
		return p.varDeclaration()
	}
	if p.match(Print) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() Stmt {
	expr := p.expression()
	p.consume(Semicolon, "Expected ';' at end of line.")
	return &PrintStmt{expr}
}

func (p *Parser) expressionStatement() Stmt {
	expr := p.expression()
	p.consume(Semicolon, "Expected ';' at end of line.")
	return &ExpressionStmt{expr}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.getPrevToken()
		right := p.comparison()
		expr = BinaryExpr{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.addOrSubtract()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.getPrevToken()
		right := p.addOrSubtract()
		expr = BinaryExpr{expr, operator, right}
	}

	return expr
}

func (p *Parser) addOrSubtract() Expr {
	expr := p.multiplyOrDivide()

	for p.match(Plus, Minus) {
		operator := p.getPrevToken()
		right := p.multiplyOrDivide()
		expr = BinaryExpr{expr, operator, right}
	}

	return expr
}

func (p *Parser) multiplyOrDivide() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.getPrevToken()
		right := p.unary()
		expr = BinaryExpr{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	expr := p.primary()

	for p.match(Bang, Minus) {
		operator := p.getPrevToken()
		right := p.primary()
		expr = UnaryExpr{operator, right}
	}

	return expr
}

func (p *Parser) primary() Expr {
	if p.match(True) {
		return LiteralExpr{true}
	} else if p.match(False) {
		return LiteralExpr{false}
	} else if p.match(Null) {
		return LiteralExpr{nil}
	}

	if p.match(Numeric, String) {
		return LiteralExpr{p.getPrevToken().literal}
	}

	if p.match(Identifier) {
		return VariableExpr{p.getPrevToken()}
	}

	if p.match(LeftParen) {
		expr := p.expression()
		p.consume(RightParen, "Expected ')' after expression")
		return GroupingExpr{expr}
	}

	p.addParseError("No valid expression found for token")

	return nil
}

func (p *Parser) addParseError(msg string) {
	p.errors = append(p.errors, NewOkraError(p.getCurrToken().line, p.getCurrToken().col, msg))
	p.skipToNextValidExpr()
}

func (p *Parser) reportParseErrors() {
	for _, e := range p.errors {
		fmt.Println(e)
	}
	os.Exit(-1)
}

func (p *Parser) skipToNextValidExpr() {
	p.advance()
	for p.getCurrToken().tokenType != EOF {
		switch p.getCurrToken().tokenType {
		case Class:
		case Func:
		case Variable:
		case For:
		case If:
		case Print:
		case Return:
			return
		}
	}
	p.advance()
}

func (p *Parser) match(tokens ...TokenType) bool {
	for _, t := range tokens {
		if p.getCurrToken().tokenType == t {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() Token {
	if p.getCurrToken().tokenType != EOF {
		p.curr++
	}
	return p.getPrevToken()
}

func (p *Parser) consume(tokenType TokenType, msg string) Token {
	if p.getCurrToken().tokenType == tokenType {
		return p.advance()
	}
	p.addParseError(msg)
	return Token{}
}

func (p *Parser) getCurrToken() Token {
	return *p.tokens[p.curr]
}

func (p *Parser) getPrevToken() Token {
	return *p.tokens[p.curr-1]
}
