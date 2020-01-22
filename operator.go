package scalc

import "fmt"

// Operator names.
const (
	Sum = "SUM"
	Int = "INT"
	Dif = "DIF"
)

// Operator represents an expression operator.
type Operator interface {
	// Calculate returns a result of operation with sets.
	Calculate(sets ...[]int) []int
	// String returns a short name of operator.
	fmt.Stringer
}

// NewOperator returns an operator for a given name.
func NewOperator(name string) Operator {
	switch name {
	case Sum:
		return unionOperator{}
	case Int:
		return intersectionOperator{}
	case Dif:
		return differenceOperator{}
	default:
		return nil
	}
}

// IsOperator returns true if the given name is operator.
func IsOperator(name string) bool {
	switch name {
	case Sum, Int, Dif:
		return true
	default:
		return false
	}
}

// unionOperator calculates an unionOperator join of sets.
type unionOperator struct{}

func (unionOperator) Calculate(sets ...[]int) []int {
	result := []int{}

	for i := range sets {
		result = union(result, sets[i])
	}

	return result
}

func (unionOperator) String() string {
	return Sum
}

// intersectionOperator calculates an intersectionOperator join of sets.
type intersectionOperator struct{}

func (intersectionOperator) Calculate(sets ...[]int) []int {
	switch len(sets) {
	case 0, 1:
		return nil
	default:
		result := []int{}
		result = sets[0]
		sets = sets[1:]
		for i := range sets {
			result = intersect(result, sets[i])
		}

		return result
	}
}

func (intersectionOperator) String() string {
	return Int
}

// differenceOperator returns a differenceOperator of the first set and the rest ones.
type differenceOperator struct{}

func (differenceOperator) Calculate(sets ...[]int) []int {
	switch len(sets) {
	case 0:
		return nil
	case 1:
		return sets[0]
	default:
		result := sets[0]
		sets = sets[1:]
		for i := range sets {
			result = diff(result, sets[i])
		}

		return result
	}
}

func (differenceOperator) String() string {
	return Dif
}

// union returns an union of two sorted sets.
func union(a, b []int) []int {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	result := make([]int, 0, len(a))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			result = append(result, a[i])
			i++
		case a[i] > b[j]:
			result = append(result, b[j])
			j++
		case a[i] == b[j]:
			result = append(result, a[i])
			i++
			j++
		}
	}

	var tail []int
	if len(a[i:]) > len(b[j:]) {
		tail = a[i:]
	} else {
		tail = b[j:]
	}

	return append(result, tail...)
}

// intersect returns an intersection of two sorted sets.
func intersect(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	result := make([]int, 0)

	for i, j := 0, 0; i < len(a) && j < len(b); {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		case a[i] == b[j]:
			result = append(result, a[i])
			i++
			j++
		}
	}

	return result
}

// diff returns a difference betwean sorted A and B.
func diff(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	result := make([]int, 0)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			result = append(result, a[i])
			i++
		case a[i] > b[j]:
			j++
		case a[i] == b[j]:
			i++
			j++
		}
	}

	return append(result, a[i:]...)
}
