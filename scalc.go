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
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Calculator represents a calculator interface.
type Calculator interface {
	Calculate() ([]int, error)
}

// New parses input string for expression
// and returns calculator for a given one.
func New(s string) (Calculator, error) {
	return parseInput(s)
}

type calc struct {
	Expr *expression
}

func (c *calc) Calculate() ([]int, error) {
	result, err := c.Expr.Calculate()
	if err != nil {
		return nil, err
	}

	sort.Ints(result)

	return result, nil

}

type expression struct {
	parent   *expression
	operator string
	operands []*operand
}

func newExpression() *expression {
	return &expression{}
}

func (e *expression) IsRoot() bool {
	return e != nil && e.parent == nil
}

func (e *expression) Parent() *expression {
	return e.parent
}

func (e *expression) NewExpression() *expression {
	child := &expression{parent: e}

	e.add(&operand{expr: child})

	return child
}

func (e *expression) SetOperator(op string) {
	e.operator = op
}

func (e *expression) AddFile(file string) {
	e.add(&operand{file: file})
}

func (e *expression) add(s *operand) {
	e.operands = append(e.operands, s)
}

func (e *expression) Calculate() ([]int, error) {
	if e == nil || e.operator == "" {
		return nil, nil
	}

	operands := make([][]int, 0, len(e.operands))

	for _, o := range e.operands {
		set, err := o.Value()
		if err != nil {
			return nil, err
		}

		operands = append(operands, set)
	}

	return calculate(e.operator, operands...), nil
}

// Equal checks expressions for equality.
func (e *expression) Equal(x *expression) bool {
	if e == nil && x == nil {
		return true
	}

	if e.operator != x.operator || len(e.operands) != len(x.operands) {
		return false
	}

	for i := range e.operands {
		if !e.operands[i].Equal(x.operands[i]) {
			return false
		}
	}
	return true
}

func (e *expression) String() string {
	if e == nil || e.operator == "" {
		return "[ EMPTY ]"
	}

	operands := make([]string, 0, len(e.operands))

	for _, o := range e.operands {
		operands = append(operands, o.String())
	}

	return fmt.Sprintf("[ %s %s ]", e.operator, strings.Join(operands, " "))
}

type operand struct {
	file string
	expr *expression
}

func (o *operand) Value() ([]int, error) {
	var (
		val []int
		err error
	)

	switch {
	case o.file != "":
		val, err = readFromFile(o.file)
		if err != nil {
			err = fmt.Errorf("read from file: %w", err)
		}
	case o.expr != nil:
		val, err = o.expr.Calculate()
	default:
		err = ErrEmptyOperand
	}

	return val, err
}

// Equal checks operands for equality.
func (o *operand) Equal(x *operand) bool {
	return o.file == x.file && o.expr.Equal(x.expr)
}

func (o *operand) String() string {
	switch {
	case o.file != "":
		return o.file
	case o.expr != nil:
		return o.expr.String()
	default:
		return "<EMPTY>"
	}
}

func readFromFile(name string) ([]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dig, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, dig)
	}

	return result, scanner.Err()
}
