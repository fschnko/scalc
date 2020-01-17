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
	result, err := c.Expr.calculate()
	if err != nil {
		return nil, err
	}

	sort.Ints(result)

	return result, nil

}

type expression struct {
	parent   *expression
	Operator string
	Sets     []*set
}

func (e *expression) Add(s *set) {
	e.Sets = append(e.Sets, s)
}

func (e *expression) calculate() ([]int, error) {
	if e == nil || e.Operator == "" {
		return nil, nil
	}

	sets := make([][]int, 0, len(e.Sets))

	for _, s := range e.Sets {
		set, err := s.Set()
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}

	return calculate(e.Operator, sets...), nil
}

func (e *expression) Equal(x *expression) bool {
	if e == nil && x == nil {
		return true
	}

	if e.Operator != x.Operator || len(e.Sets) != len(x.Sets) {
		return false
	}

	for i := range e.Sets {
		if !e.Sets[i].Equal(x.Sets[i]) {
			return false
		}
	}
	return true
}

type set struct {
	File string
	Expr *expression
}

func (s *set) Set() ([]int, error) {
	var (
		set []int
		err error
	)

	if s.File != "" {
		set, err = readFromFile(s.File)
		if err != nil {
			return nil, fmt.Errorf("read from file: %w", err)
		}
	} else {
		set, err = s.Expr.calculate()
	}
	return set, err
}

func (s *set) Equal(x *set) bool {
	return s.File == x.File && s.Expr.Equal(x.Expr)
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
