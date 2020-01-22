package scalc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Expression represents an mathematic expression in the tree form.
type Expression struct {
	parent   *Expression
	operator Operator
	operands []*operand
}

// NewExpression creates a new empty expression.
func NewExpression() *Expression {
	return &Expression{}
}

// IsRoot returns true if the expression is root (have no parent).
func (e *Expression) IsRoot() bool {
	return e != nil && e.parent == nil
}

// Parent returns an expression parent.
func (e *Expression) Parent() *Expression {
	return e.parent
}

// NewExpression creates a new child expression and returns it.
func (e *Expression) NewExpression() *Expression {
	child := &Expression{parent: e}

	e.add(&operand{expr: child})

	return child
}

// SetOperator sets an operator of expression.
func (e *Expression) SetOperator(op Operator) {
	e.operator = op
}

// AddFile adds a file type operand.
func (e *Expression) AddFile(file string) {
	e.add(&operand{file: file})
}

func (e *Expression) add(s *operand) {
	e.operands = append(e.operands, s)
}

// Value returns a value of the expression.
func (e *Expression) Value() ([]int, error) {
	if e == nil || e.operator == nil {
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

	return e.operator.Calculate(operands...), nil
}

// Equal checks expressions for equality.
func (e *Expression) Equal(x *Expression) bool {
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

func (e *Expression) String() string {
	if e == nil || e.operator == nil {
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
	expr *Expression
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
		val, err = o.expr.Value()
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
