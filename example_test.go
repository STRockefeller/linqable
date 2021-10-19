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

func ExampleLinqablize_third() {
	// for numeric types such as int/uint/float
	var nt int32
	Linqablize(reflect.TypeOf(nt), "main", IsNumericType(), HasDefaultValue("100"))
	// Output:
	// new file linqable_int.go with methods Max, Min, ...
}
