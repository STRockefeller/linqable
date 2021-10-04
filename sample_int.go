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

func (si LinqableInt) ToSlice() []int {
	return si
}
