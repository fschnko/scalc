package scalc

import (
	"fmt"
)

const (
	expressionToken tokenType = iota
	operatorToken
	fileToken
)

type tokenType int

type parser struct {
	runes     []rune
	level     int
	expr      *expression
	tokenType tokenType
}

func parseInput(s string) (*calc, error) {
	p := &parser{
		runes: []rune(s),
	}

	for i, next := 0, 0; len(p.runes) > i; i = next {
		if p.level < 0 {
			return nil, ErrExpressionSyntax
		}

		switch p.runes[i] {
		case '[':
			p.levelUp()
		case ']':
			p.levelDown()
		case ' ':
			// skip whitespaces
		default:
			next = firstIndexAfterN(p.runes, ' ', i)
			if next < 0 {
				return nil, ErrExpressionSyntax
			}
			err := p.parseToken(i, next)
			if err != nil {
				return nil, err
			}
		}
		next++
	}

	return &calc{Expr: p.expr}, nil
}

func (p *parser) levelUp() {
	p.level++
	p.expr = &expression{parent: p.expr}
	if p.expr.parent != nil {
		p.expr.parent.Add(&set{Expr: p.expr})
	}

	p.tokenType = expressionToken
}

func (p *parser) levelDown() {
	p.level--
	if p.level != 0 {
		p.expr = p.expr.parent
	}
	p.tokenType = operatorToken
}

func (p *parser) parseToken(at, to int) error {

	token, err := cut(p.runes, at, to)
	if err != nil {
		return err
	}

	switch p.tokenType {
	case expressionToken:
		p.tokenType = operatorToken
		p.expr.Operator = token
	case operatorToken, fileToken:
		p.tokenType = fileToken
		p.expr.Add(&set{File: token})
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
