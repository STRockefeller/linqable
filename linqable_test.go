package linqable

import (
	"fmt"
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
	Linqablize(ti, "linqable", HasDefaultValue("int64(88888)"), IsNumericType())
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
	assert := assert.New(t)
	si := (linqableInt)([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	{ // case ToSlice
		assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, si.ToSlice())
	}
	{ // case Where
		actual := si.Where(func(i int) bool { return i%2 == 0 }).ToSlice()
		assert.Equal([]int{0, 2, 4, 6, 8}, actual)
	}
	{ // case Take
		actual := si.Take(3).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // case Skip
		actual := si.Skip(5).ToSlice()
		assert.Equal([]int{5, 6, 7, 8, 9}, actual)
	}
	{ // case TakeWhile
		actual := si.TakeWhile(func(i int) bool { return i < 5 }).ToSlice()
		assert.Equal([]int{0, 1, 2, 3, 4}, actual)
	}
	{ // case SkipWhile
		actual := si.SkipWhile(func(i int) bool { return i < 8 }).ToSlice()
		assert.Equal([]int{8, 9}, actual)
	}
	{ // case Contains
		actual := si.Contains(3)
		assert.Equal(true, actual)
	}
	{ // case Contains
		actual := si.Contains(10)
		assert.Equal(false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i > 10 })
		assert.Equal(false, actual)
	}
	{ // case Any
		actual := si.Any(func(i int) bool { return i < 2 })
		assert.Equal(true, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i < 3 })
		assert.Equal(false, actual)
	}
	{ // case All
		actual := si.All(func(i int) bool { return i >= 0 })
		assert.Equal(true, actual)
	}
	{ // case TakeLast
		actual := si.TakeLast(3).ToSlice()
		assert.Equal([]int{7, 8, 9}, actual)
	}
	{ // case SkipLast
		actual := si.SkipLast(7).ToSlice()
		assert.Equal([]int{0, 1, 2}, actual)
	}
	{ // case Count
		actual := si.Count(func(i int) bool { return i%2 == 1 })
		assert.Equal(5, actual)
	}
	{ // case Append
		actual := si.Take(2).Append(3).ToSlice()
		assert.Equal([]int{0, 1, 3}, actual)
	}
	{ // case ElementAt
		actual := si.ElementAt(3)
		assert.Equal(3, actual)
	}
	{ // case First
		actual := si.First(func(i int) bool { return i > 2 })
		assert.Equal(3, actual)
	}
	{ // case FirstOrDefault
		actual := si.FirstOrDefault(func(i int) bool { return i > 100 })
		assert.Equal(0, actual)
	}
	{ // case Last
		actual := si.Last(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // case LastOrDefault
		actual := si.LastOrDefault(func(i int) bool { return i < 8 })
		assert.Equal(7, actual)
	}
	{ // case Prepend
		actual := si.Preppend(999).First(func(i int) bool { return true })
		assert.Equal(999, actual)
	}
	{ // case Reverse
		actual := si.Reverse().ToSlice()
		assert.Equal([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, actual)
	}
	{ // case Single
		actual := si.Single(func(i int) bool { return i < 1 })
		assert.Equal(0, actual)
	}
	{ // case SingleOeDefault
		actual := si.SingleOrDefault(func(i int) bool { return i > 3 })
		assert.Equal(0, actual)
	}
	{ // case SumInt
		actual := si.SumInt32(func(i int) int32 { return int32(i) })
		assert.Equal(int32(45), actual)
	}
	{ // case SumFloat
		actual := si.SumFloat32(func(i int) float32 { return float32(i) })
		assert.Equal(float32(45.0), actual)
	}
	{ // case Max
		actual := si.Max()
		assert.Equal(9, actual)
	}
	{ // case Min
		actual := si.Min()
		assert.Equal(0, actual)
	}
	{ // case Repeat
		actual := repeatInt(123, 3).ToSlice()
		assert.Equal([]int{123, 123, 123}, actual)
	}
	{ // case ForEach
		si.ForEach(func(i int) { fmt.Println("Foreach test ", i) })
	}
	{ // case Remove
		actual := newLinqableInt([]int{1, 2, 3, 4})
		actual2 := actual.Remove(3)
		assert.True(actual2)
		assert.Equal(newLinqableInt([]int{1, 2, 4}), actual)
	}
	{ // case RemoveAll
		actual := newLinqableInt([]int{1, 2, 3, 4, 5, 6, 7})
		actual2 := actual.RemoveAll(func(i int) bool { return i%2 == 1 })
		assert.Equal(4, actual2)
		assert.Equal(newLinqableInt([]int{2, 4, 6}), actual)
	}
	{ // case RemoveAt
		actual := newLinqableInt([]int{1, 2, 3, 4, 5})
		actual.RemoveAt(3)
		assert.Equal(newLinqableInt([]int{1, 2, 3, 5}), actual)
	}
	{ // case RemoveRange
		actual := newLinqableInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(-1, 3)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // case RemoveRange
		actual := newLinqableInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(1, 50)
		assert.Equal(fmt.Errorf("argument out of range"), err)
	}
	{ // case RemoveRange
		actual := newLinqableInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		err := actual.RemoveRange(2, 2)
		assert.NoError(err)
		assert.Equal(newLinqableInt([]int{1, 2, 5, 6, 7, 8, 9}), actual)
	}
	{ // case Distinct
		actual := newLinqableInt([]int{1, 2, 3, 1, 5, 5, 2, 3, 8}).Distinct().ToSlice()
		assert.Equal([]int{1, 2, 3, 5, 8}, actual)
	}
	{ // OrderBy
		si := newLinqableInt([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderBy(func(i int) int { return i })
		assert.Equal(newLinqableInt([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), orderedSi)
	}
	{ // OrderByDescending
		si := newLinqableInt([]int{5, 8, 2, 3, 6, 9, 4, 1, 7, 0})
		orderedSi := si.OrderByDescending(func(i int) int { return i })
		assert.Equal(newLinqableInt([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), orderedSi)
	}
	{ // Add
		si := newLinqableInt([]int{1, 2, 3})
		si.Add(4)
		assert.Equal(newLinqableInt([]int{1, 2, 3, 4}), si)
	}
	{ // Add Range
		si := newLinqableInt([]int{1, 2, 3})
		si.AddRange([]int{4, 5, 6})
		assert.Equal(si, newLinqableInt([]int{1, 2, 3, 4, 5, 6}))
	}
	{ // Clear
		si := newLinqableInt([]int{1, 2, 3})
		capacity := cap(si.ToSlice())
		si.Clear()
		assert.Equal(si, newLinqableInt(make([]int, capacity)))
	}
	{ // Exists
		si := newLinqableInt([]int{1, 2, 3})
		assert.True(si.Exists(func(i int) bool { return i == 2 }))
		assert.False(si.Exists(func(i int) bool { return i-10 > 0 }))
	}
	{ // Find
		si := newLinqableInt([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(6, si.Find(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindAll
		si := newLinqableInt([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(newLinqableInt([]int{6, 8}), si.FindAll(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindIndex
		si := newLinqableInt([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(3, si.FindIndex(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLast
		si := newLinqableInt([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(8, si.FindLast(func(i int) bool { return i%2 == 0 }))
	}
	{ // FindLastIndex
		si := newLinqableInt([]int{1, 3, 5, 6, 7, 8, 9})
		assert.Equal(5, si.FindLastIndex(func(i int) bool { return i%2 == 0 }))
	}
}
