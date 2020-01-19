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
	"sort"
)

// Calculator represents a calculator interface.
type Calculator interface {
	// Calculate calculates an expression an returns sorted result.
	Calculate() ([]int, error)
}

// New returns calculator for a given expression.
func New(expr *Expression) Calculator {
	return &calc{expr: expr}
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
