/*Package scalc provides a basic sets calculator that parses an expression string.

Grammar of calculator is given:
expression := “[“ operator sets “]”
sets := set | set sets
set := file | expression
operator := “SUM” | “INT” | “DIF”

Each file should contain sorted integers, one integer in a line.
Meaning of operators:
SUM - returns union of all sets
INT - returns intersection of all sets
DIF - returns difference of first set and the rest ones
*/
package scalc

import (
	"fmt"
	"sort"
)

// Calculator represents a calculator interface.
type Calculator interface {
	Calculate() ([]int, error)
}

// New parses input string for expression
// and returns calculator for a given one.
func New(s string) (Calculator, error) {
	expr, err := NewParser(s).Process()
	if err != nil {
		return nil, fmt.Errorf("process expression: %w", err)
	}

	return &calc{expr: expr}, nil
}

type calc struct {
	expr *Expression
}

func (c *calc) Calculate() ([]int, error) {
	result, err := c.expr.Value()
	if err != nil {
		return nil, err
	}

	sort.Ints(result)

	return result, nil
}
