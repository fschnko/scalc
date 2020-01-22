package scalc

// Operations
const (
	Sum = "SUM"
	Int = "INT"
	Dif = "DIF"
)

func calculate(operator string, operands ...[]int) []int {
	switch operator {
	case Sum:
		return union(operands...)
	case Int:
		return intersection(operands...)
	case Dif:
		return difference(operands...)
	default:
		return nil
	}
}

// union returns a union of sorted sets.
func union(sets ...[]int) []int {
	result := []int{}

	for i := range sets {
		result = sum(result, sets[i])
	}

	return result
}

// intersection returns an intersection of sets.
func intersection(sets ...[]int) []int {
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

// difference returns a difference of the first set and the rest ones.
func difference(sets ...[]int) []int {
	switch len(sets) {
	case 0:
		return nil
	case 1:
		return sets[0]
	default:
		result := []int{}
		first := sets[0]
		sets = sets[1:]

		for i := range sets {
			result = sum(result, diff(first, sets[i]))
		}

		return result
	}
}

// sum returns an union of two sorted sets.
func sum(a, b []int) []int {
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

func diff(a, b []int) []int {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
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
			result = append(result, b[j])
			j++
		case a[i] == b[j]:
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
