# Linqable



## Abstract

linq realization in golang



## Feature

**Advantage**

* Source code generation. You can delete the import statement after use.
* There is no type judgment in the function. The execution speed is close to the for loop
* strongly-typed: You will not see any  `interface{}` appears in the input or output of the methods. Method signatures of the delegate functions are all fixed.

**Disadvantage**

* Only slices are currently supported.
* This package will **not** contain all the methods of linq.
* A panic will show if anything wrong with this package.





## Installation

```powershell
go get github.com/STRockefeller/linqable
```



## Quickstart

### Import

Import this package and [reflect](https://pkg.go.dev/reflect)

```go
import (
	"reflect"

	"github.com/STRockefeller/linqable"
)
```



### Generate Source Code

```go
func linqable.Linqablize(t reflect.Type, packageName string)
```



```go
func main() {
	var i int
	linqable.Linqablize(reflect.TypeOf(i),"main")
}
```



Run the code then a new `.go` file will be generated with the file name `linqable_XXXXX.go` where "XXXX" is the type name.



#### if the type is imported by other package

```go
linqable.Linqablize(reflect.TypeOf(myStruct),"main", linqable.IsImportedType())
```





### Delete the imports and start linq

~~import "reflect"~~

~~import "github.com/STRockefeller/linqable"~~



```go
	nums:=NewLinqableInt([]int{1,2,3,4,5,6,7,8,9,0})
	firstThreeOddNums := nums.Where(func(i int) bool {return i%2==1}).Take(3).ToSlice()
```



## Methods

Supportable methods

* Where
* Take
* TakeWhile
* Skip
* SkipWhile
* ToSlice

