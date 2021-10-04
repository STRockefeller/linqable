package linqable

import (
	"reflect"
	"testing"
)

func TestLinqablize(t *testing.T) {
	var i string
	ti := reflect.TypeOf(i)
	Linqablize(ti, "linqable")
}
