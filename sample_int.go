package linqable

type LinqableInt []int

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

func (si LinqableInt) Skip(n int) LinqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: Skip() out of index")
	}
	return si[n:]
}

func (si LinqableInt) ToSlice() []int {
	return si
}
