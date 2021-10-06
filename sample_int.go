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

func (si linqableInt) Count(predicate func(int) bool) int {
	var count int
	for _, i := range si {
		if predicate(i) {
			count++
		}
	}
	return count
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

func (si linqableInt) Append(i int) linqableInt {
	return append(si, i)
}

func (si linqableInt) ElementAt(index int) int {
	if index >= len(si) {
		panic("linq: ElementAt() out of index")
	}
	return si[index]
}

func (si linqableInt) ElementAtOrDefault(index int) int {
	var defaultValue int
	if index >= len(si) {
		return defaultValue
	}
	return si[index]
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
	if n < 0 || n >= len(si) {
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

func (si linqableInt) TakeLast(n int) linqableInt {
	if n < 0 || n >= len(si) {
		panic("Linq: TakeLast() out of index")
	}
	return si.Skip(len(si) - n)
}

func (si linqableInt) Skip(n int) linqableInt {
	if n < 0 || n >= len(si) {
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

func (si linqableInt) SkipLast(n int) linqableInt {
	if n < 0 || n > len(si) {
		panic("Linq: SkipLast() out of index")
	}
	return si.Take(len(si) - n)
}

func (si linqableInt) ToSlice() []int {
	return si
}
