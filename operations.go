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
// Algorithm has side effects for input data.
func union(sets ...[]int) []int {
	result := []int{}

	for i := range sets {
		result = sum(result, sets[i])
	}

	return result
}

// intersection returns an intersection of sets.
// Algorithm has side effects for input data.
func intersection(sets ...[]int) []int {
	result := []int{}
	if len(sets) > 0 {
		result = sets[0]
		sets = sets[1:]
	}

	for i := range sets {
		result = intersect(result, sets[i])
	}
	return result
}

// difference returns a difference of the first set and the rest ones.
// Algorithm has side effects for input data.
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
	for i, j := 0, 0; len(a) > i && len(b) > j; {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			a[i], b[j] = b[j], a[i]
			if len(b) > j+1 && b[j] == b[j+1] {
				b = append(b[:j], b[j+1:]...)
				break
			}
			j++
		case a[i] == b[j]:
			b = append(b[:j], b[j+1:]...)
		}
	}

	return append(a, b...)
}

// intersect returns an intersection of two sorted sets.
func intersect(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	result := make([]int, 0)
	for i, j := 0, 0; len(a) > i && len(b) > j; {
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

	shadow := make([]int, len(a))
	copy(shadow, a)

	for i, j := 0, 0; len(shadow) > i && len(b) > j; {
		switch {
		case shadow[i] < b[j]:
			i++
		case shadow[i] > b[j]:
			shadow[i], b[j] = b[j], shadow[i]
			j++
		case shadow[i] == b[j]:
			shadow = append(shadow[:i], shadow[i+1:]...)
			b = append(b[:j], b[j+1:]...)
		}
	}

	return append(shadow, b...)
}

// normalize removes duplicates from a set.
func normalize(a []int) []int {
	for i := 0; len(a) > i; i++ {
		for j := i + 1; len(a) > j; j++ {
			if a[i] == a[j] {
				a = append(a[:j], a[j+1:]...)
				j--
			}
		}
	}
	return a
}
