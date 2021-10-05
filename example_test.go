package linqable

import "reflect"

func ExampleLinqablize() {
	var s string
	Linqablize(reflect.TypeOf(s), "main")

	// Output:
	// new file linqable_string.go
}

func ExampleLinqablize_second() {
	// for imported types
	var it reflect.Method
	Linqablize(reflect.TypeOf(it), "main", IsImportedType())

	// Output:
	// new file linqable_Method.go
}
