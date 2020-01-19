package scalc

type rstack []rune

func (rs *rstack) Push(r rune) {
	*rs = append(*rs, r)
}

func (rs *rstack) Pop() rune {
	if len(*rs) == 0 {
		return 0
	}
	r := (*rs)[len(*rs)-1]
	*rs = (*rs)[:len(*rs)-1]

	return r
}
