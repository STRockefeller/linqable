package linqable

type linqableInt []int

func newLinqableInt(si []int) linqableInt {
	return si
}

func (si linqableInt) Contains(target int) bool {
	for _, i := range si {
		if i == target {
			return true
		}
	}
	return false
}

func (si linqableInt) Any(predicate func(int) bool) bool {
	for _, i := range si {
		if predicate(i) {
			return true
		}
	}
	return false
}

func (si linqableInt) All(predicate func(int) bool) bool {
	for _, i := range si {
		if predicate(i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func (si linqableInt) Where(predicate func(int) bool) linqableInt {
	res := []int{}
	for _, i := range si {
		if predicate(i) {
			res = append(res, i)
		}
	}
	return res
}

func (si linqableInt) Take(n int) linqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: Take() out of index")
	}
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, si[i])
	}
	return res
}

func (si linqableInt) TakeWhile(predicate func(int) bool) linqableInt {
	res := []int{}
	for i := 0; i < len(si); i++ {
		if predicate(si[i]) {
			res = append(res, si[i])
		} else {
			return res
		}
	}
	return res
}

func (si linqableInt) Skip(n int) linqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: Skip() out of index")
	}
	return si[n:]
}

func (si linqableInt) SkipWhile(predicate func(int) bool) linqableInt {
	for i := 0; i < len(si); i++ {
		if predicate(si[i]) {
			continue
		} else {
			return si[i:]
		}
	}
	return linqableInt{}
}

func (si linqableInt) ToSlice() []int {
	return si
}
