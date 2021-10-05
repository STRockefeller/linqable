package linqable

type LinqableInt []int

func NewLinqableInt(si []int) LinqableInt {
	return si
}

func (si LinqableInt) Where(predicate func(int) bool) LinqableInt {
	res := []int{}
	for _, i := range si {
		if predicate(i) {
			res = append(res, i)
		}
	}
	return res
}

func (si LinqableInt) Take(n int) LinqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: Take() out of index")
	}
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, si[i])
	}
	return res
}

func (si LinqableInt) TakeWhile(predicate func(int) bool) LinqableInt {
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

func (si LinqableInt) Skip(n int) LinqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: Skip() out of index")
	}
	return si[n:]
}

func (si LinqableInt) SkipWhile(predicate func(int) bool) LinqableInt {
	for i := 0; i < len(si); i++ {
		if predicate(si[i]) {
			continue
		} else {
			return si[i:]
		}
	}
	return LinqableInt{}
}

func (si LinqableInt) ToSlice() []int {
	return si
}
