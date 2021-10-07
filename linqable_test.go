package linqable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinqablize_string(t *testing.T) {
	var i string
	ti := reflect.TypeOf(i)
	Linqablize(ti, "linqable")
}

func TestLinqablize_long_withDefault(t *testing.T) {
	var i int64
	ti := reflect.TypeOf(i)
	Linqablize(ti, "linqable", HasDefaultValue("int64(88888)"))
}
func TestStruct(t *testing.T) {
	var ms MyStruct
	Linqablize(reflect.TypeOf(ms), "linqable", IsImportedType())
}

type MyStruct struct {
	a int
	b string
}

func TestSample(t *testing.T) {
	si := (linqableInt)([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	{ // case ToSlice
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, si.ToSlice())
	}
	{ // case Where
		actual := si.Where(func(i int) bool { return i%2 == 0 }).ToSlice()
		assert.Equal(t, []int{0, 2, 4, 6, 8}, actual)
	}
	{ // case Take
		actual := si.Take(3).ToSlice()
		assert.Equal(t, []int{0, 1, 2}, actual)
	}
	{ // case Skip
		actual := si.Skip(5).ToSlice()
		assert.Equal(t, []int{5, 6, 7, 8, 9}, actual)
	}
	{ // case TakeWhile
		actual := si.TakeWhile(func(i int) bool { return i < 5 }).ToSlice()
		assert.Equal(t, []int{0, 1, 2, 3, 4}, actual)
	}
	{ // case SkipWhile
		actual := si.SkipWhile(func(i int) bool { return i < 8 }).ToSlice()
		assert.Equal(t, []int{8, 9}, actual)
	}
	{ // case Contains
		actual := si.Contains(3)
		assert.Equal(t, true, actual)
	}
	{ // case Contains
		actual := si.Contains(10)
		assert.Equal(t, false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i > 10 })
		assert.Equal(t, false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i < 2 })
		assert.Equal(t, true, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i < 3 })
		assert.Equal(t, false, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i >= 0 })
		assert.Equal(t, true, actual)
	}
	{ // case TakeLast
		actual := si.TakeLast(3).ToSlice()
		assert.Equal(t, []int{7, 8, 9}, actual)
	}
	{ // case SkipLast
		actual := si.SkipLast(7).ToSlice()
		assert.Equal(t, []int{0, 1, 2}, actual)
	}
	{ // case Count
		actual := si.Count(func(i int) bool { return i%2 == 1 })
		assert.Equal(t, 5, actual)
	}
	{ // case Append
		actual := si.Take(2).Append(3).ToSlice()
		assert.Equal(t, []int{0, 1, 3}, actual)
	}
	{ // case ElementAt
		actual := si.ElementAt(3)
		assert.Equal(t, 3, actual)
	}
	{ // case First
		actual := si.First(func(i int) bool { return i > 2 })
		assert.Equal(t, 3, actual)
	}
	{ // case FirstOrDefault
		actual := si.FirstOrDefault(func(i int) bool { return i > 100 })
		assert.Equal(t, 0, actual)
	}
	{ // case Last
		actual := si.Last(func(i int) bool { return i < 8 })
		assert.Equal(t, 7, actual)
	}
	{ // case LastOrDefault
		actual := si.LastOrDefault(func(i int) bool { return i < 8 })
		assert.Equal(t, 7, actual)
	}
}
