package scalc

import (
	"fmt"
)

const (
	rootToken tokenType = iota
	expressionToken
	operatorToken
	operandToken
)

type tokenType int

// Parser a expression parser.
type Parser struct {
	runes         []rune
	expr          *Expression
	prevTokenType tokenType
}

// New creates new parser for a string.
func New(s string) *Parser {
	return &Parser{
		runes:         []rune(s),
		prevTokenType: rootToken,
		expr:          NewExpression(),
	}
}

// Process parsess expression.
func (p *Parser) Process() (*Expression, error) {
	if len(p.runes) == 0 {
		return nil, ednOfLineSyntaxError(0)
	}

	brackets := make(rstack, 0)

	for i, next := 0, 0; next < len(p.runes); i = next {
		switch p.runes[i] {
		case '[':
			brackets.Push(p.runes[i])

			if p.prevTokenType == expressionToken {
				return nil, tokenSyntaxError(string(p.runes[i]), i)
			}

			p.up()
		case ']':
			r := brackets.Pop()
			if p.prevTokenType != operandToken || !isPair(r, p.runes[i]) {
				return nil, tokenSyntaxError(string(p.runes[i]), i)
			}

			p.down()
		case ' ', '\t', '\n':
			// skip whitespaces
		default:
			next = endOfToken(p.runes, i)
			if next < 0 {
				return nil, ednOfLineSyntaxError(len(p.runes))
			}

			err := p.processToken(i, next)
			if err != nil {
				return nil, err
			}
		}
		next++
	}

	if len(brackets) > 0 {
		return nil, ednOfLineSyntaxError(len(p.runes))
	}

	return p.expr, nil
}

func (p *Parser) up() {
	if p.prevTokenType != rootToken {
		p.expr = p.expr.NewExpression()
	}

	p.prevTokenType = expressionToken
}

func (p *Parser) down() {
	if !p.expr.IsRoot() {
		p.expr = p.expr.Parent()
	}

	p.prevTokenType = operandToken
}

func (p *Parser) processToken(at, to int) error {
	token, err := cut(p.runes, at, to+1)
	if err != nil {
		return err
	}

	switch p.prevTokenType {
	case expressionToken:
		if !isOperator(token) {
			return tokenSyntaxError(token, at)
		}

		p.prevTokenType = operatorToken

		p.expr.SetOperator(token)
	case operatorToken, operandToken:
		p.prevTokenType = operandToken

		p.expr.AddFile(token)
	default:
		return tokenSyntaxError(token, at)
	}

	return nil
}

func cut(runes []rune, at, to int) (string, error) {
	if at >= to || len(runes) < to {
		return "", fmt.Errorf("out of bounds: offset %d, len %d", to, len(runes))
	}

	result := make([]rune, to-at)
	copy(result, runes[at:to])

	return string(result), nil
}

func endOfToken(runes []rune, n int) int {
	if n < 0 {
		n = 0
	}

	for i := n; i < len(runes); i++ {
		if isCloserToken(runes[i]) {
			return i - 1
		}
	}

	return -1
}

func isCloserToken(r rune) bool {
	switch r {
	case ' ', '\n', '\t', ']':
		return true
	default:
		return false
	}
}

func isPair(open, close rune) bool {
	switch open {
	case '[':
		return close == ']'
	default:
		return false
	}
}

func isOperator(op string) bool {
	switch op {
	case Sum, Int, Dif:
		return true
	default:
		return false
	}
}
