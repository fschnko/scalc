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

type parser struct {
	runes         []rune
	expr          *expression
	prevTokenType tokenType
}

func parseInput(s string) (*calc, error) {
	p := &parser{
		runes:         []rune(s),
		prevTokenType: rootToken,
		expr:          newExpression(),
	}

	for i, next := 0, 0; len(p.runes) > next; i = next {
		switch p.runes[i] {
		case '[':
			p.up()
		case ']':
			p.down()
		case ' ':
			// skip whitespaces
		default:
			next = firstIndexAfterN(p.runes, ' ', i)
			if next < 0 {
				return nil, ErrExpressionSyntax
			}
			err := p.processToken(i, next)
			if err != nil {
				return nil, err
			}
		}
		next++
	}

	return &calc{Expr: p.expr}, nil
}

func (p *parser) up() {
	if p.prevTokenType != rootToken {
		p.expr = p.expr.NewExpression()
	}
	p.prevTokenType = expressionToken
}

func (p *parser) down() {
	if !p.expr.IsRoot() {
		p.expr = p.expr.Parent()
	}
	p.prevTokenType = operandToken
}

func (p *parser) processToken(at, to int) error {

	token, err := cut(p.runes, at, to)
	if err != nil {
		return err
	}

	switch p.prevTokenType {
	case expressionToken, rootToken:
		p.prevTokenType = operatorToken

		p.expr.SetOperator(token)
	case operatorToken, operandToken:
		p.prevTokenType = operandToken

		p.expr.AddFile(token)
	default:
		return ErrExpressionSyntax
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

func firstIndexAfterN(runes []rune, r rune, n int) int {
	if n < 0 {
		n = 0
	}
	for i := n; len(runes) > i; i++ {
		if runes[i] == r {
			return i
		}
	}
	return -1
}
