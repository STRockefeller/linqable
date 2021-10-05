package linqable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinqablize(t *testing.T) {
	var i string
	ti := reflect.TypeOf(i)
	Linqablize(ti, "linqable")
}

func TestSample(t *testing.T) {
	si := (LinqableInt)([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
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
}
