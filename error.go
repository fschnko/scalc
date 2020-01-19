package scalc

import "fmt"

// Package constant errors.
const (
	ErrExpressionSyntax Error = "expression syntax error"
	ErrEmptyOperand     Error = "operand is empty"
)

// Error represents constant error.
type Error string

func (e Error) Error() string {
	return string(e)
}

func ednOfLineSyntaxError(index int) error {
	return fmt.Errorf("%w: unexpected end of the line at col %d", ErrExpressionSyntax, index)
}

func tokenSyntaxError(token string, index int) error {
	return fmt.Errorf("%w: unexpected token %q at col %d", ErrExpressionSyntax, token, index)
}
