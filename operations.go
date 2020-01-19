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

// union returns a union of sets.
func union(sets ...[]int) []int {
	result := []int{}
	if len(sets) > 0 {
		result = sets[0]
		sets = sets[1:]
	}

	for i := range sets {
		result = append(result, sets[i]...)
	}

	return normalize(result)
}

// intersection returns an intersection of sets.
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
func difference(sets ...[]int) []int {
	switch len(sets) {
	case 0:
		return nil
	case 1:
		return normalize(sets[0])
	default:
		result := []int{}
		first := sets[0]
		sets = sets[1:]
		for i := range sets {
			result = append(result, diff(first, sets[i])...)
		}
		return normalize(result)
	}
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

// intersect returns an intersection of two sets.
func intersect(a, b []int) []int {
	result := make([]int, 0)
	for i := 0; len(a) > i; i++ {
		for j := 0; len(b) > j; j++ {
			if a[i] == b[j] {
				result = append(result, a[i])
			}
		}
	}
	return normalize(result)
}

// diff returns a difference of two sets.
func diff(a, b []int) []int {
	shadow := make([]int, len(a))
	copy(shadow, a)

	for i := 0; len(shadow) > i && i >= 0; i++ {
		for j := 0; len(b) > j && i >= 0; j++ {
			if shadow[i] == b[j] {
				shadow = append(shadow[:i], shadow[i+1:]...)
				b = append(b[:j], b[j+1:]...)
				j--
				i--
			}
		}
	}
	return append(shadow, b...)
}
