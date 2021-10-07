package linqable

type linqableInt []int

func newLinqableInt(si []int) linqableInt {
	return si
}

func (si linqableInt) Contains(target int) bool {
	for _, elem := range si {
		if elem == target {
			return true
		}
	}
	return false
}

func (si linqableInt) Count(predicate func(int) bool) int {
	var count int
	for _, elem := range si {
		if predicate(elem) {
			count++
		}
	}
	return count
}

func (si linqableInt) Any(predicate func(int) bool) bool {
	for _, elem := range si {
		if predicate(elem) {
			return true
		}
	}
	return false
}

func (si linqableInt) All(predicate func(int) bool) bool {
	for _, elem := range si {
		if predicate(elem) {
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

func (si linqableInt) Preppend(i int) linqableInt {
	return append([]int{i}, si.ToSlice()...)
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

func (si linqableInt) Empty() linqableInt {
	return newLinqableInt([]int{})
}

func (si linqableInt) First(predicate func(int) bool) int {
	if len(si) <= 0 {
		panic("linq: First() empty set")
	}
	for _, elem := range si {
		if predicate(elem) {
			return elem
		}
	}
	panic("linq: First() no match element in the slice")
}

func (si linqableInt) FirstOrDefault(predicate func(int) bool) int {
	var defaultValue int
	if len(si) <= 0 {
		return defaultValue
	}
	for _, elem := range si {
		if predicate(elem) {
			return elem
		}
	}
	return defaultValue
}

func (si linqableInt) Last(predicate func(int) bool) int {
	if len(si) <= 0 {
		panic("linq: Last() empty set")
	}
	for i := len(si) - 1; i >= 0; i-- {
		if predicate(si[i]) {
			return si[i]
		}
	}
	panic("linq: Last() no match element in the slice")
}

func (si linqableInt) LastOrDefault(predicate func(int) bool) int {
	var defaultValue int
	if len(si) <= 0 {
		return defaultValue
	}
	for i := len(si) - 1; i >= 0; i-- {
		if predicate(si[i]) {
			return si[i]
		}
	}
	return defaultValue
}

func (si linqableInt) Single(predicate func(int) bool) int {
	if si.Count(predicate) == 1 {
		return si.First(predicate)
	}
	panic("linq: Single() eligible data count is not unique")
}

func (si linqableInt) SingleOrDefault(predicate func(int) bool) int {
	var defaultValue int
	if si.Count(predicate) == 1 {
		return si.First(predicate)
	}
	return defaultValue
}

func (si linqableInt) Where(predicate func(int) bool) linqableInt {
	res := []int{}
	for _, elem := range si {
		if predicate(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func (si linqableInt) Reverse() linqableInt {
	for i, j := 0, len(si)-1; i < j; i, j = i+1, j-1 {
		si[i], si[j] = si[j], si[i]
	}
	return si
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

func (si linqableInt) SumInt(selector func(int) int) int {
	var sum int
	for _, elem := range si {
		sum += selector(elem)
	}
	return sum
}

func (si linqableInt) SumInt16(selector func(int) int16) int16 {
	var sum int16
	for _, elem := range si {
		sum += selector(elem)
	}
	return sum
}

func (si linqableInt) SumInt64(selector func(int) int64) int64 {
	var sum int64
	for _, elem := range si {
		sum += selector(elem)
	}
	return sum
}

func (si linqableInt) SumFloat32(selector func(int) float32) float32 {
	var sum float32
	for _, elem := range si {
		sum += selector(elem)
	}
	return sum
}

func (si linqableInt) SumFloat64(selector func(int) float64) float64 {
	var sum float64
	for _, elem := range si {
		sum += selector(elem)
	}
	return sum
}

func (si linqableInt) ToSlice() []int {
	return si
}
