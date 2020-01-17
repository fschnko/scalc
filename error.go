package scalc

// Package constant errors.
const (
	ErrExpressionSyntax Error = "expression syntax error"
)

// Error represents constant error.
type Error string

func (e Error) Error() string {
	return string(e)
}
