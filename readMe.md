# Linqable

## Abstract

linq realization in golang

## Feature

### Advantage

* Source code generation. You can delete the import statement after use.
* There is no type judgment in the function. The execution speed is close to the for loop
* strongly-typed: You will not see any  `interface{}` appears in the input or output of the methods. Method signatures of the delegate functions are all fixed.

### Notice

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

Linq-like

* All
* Any
* Append
* Contains
* Count
* Distinct
* ElementAt
* ElementAtOrDefault
* Empty
* First
* FirstOrDefault
* Last
* LastOrDefault
* Max
* Min
* Prepend
* Repeat
* Reverse
* Single
* SingleOrDefault
* Skip
* SkipLast
* SkipWhile
* Sum
* Take
* TakeLast
* TakeWhile
* ToSlice
* Where

not linq

* ForEach
* Remove
* RemoveAll
* RemoveAt
* RemoveRange
* ReplaceAll

## Differences from [System.Linq](https://docs.microsoft.com/en-us/dotnet/api/system.linq?view=net-5.0)

### Overload Methods

for example, the method `Count` in C# can be used as

 ```C#
 public static int Count<TSource> (this System.Collections.Generic.IEnumerable<TSource> source);
 ```

or

```C#
public static int Count<TSource> (this System.Collections.Generic.IEnumerable<TSource> source, Func<TSource,bool> predicate);
```

but in the `linqableInt.Count` generated by this package, it can only be

```go
func (linqableInt) Count(predicate func(int) bool) int
```

if you want to count all elements in the slice, ~~you can use `len()`~~ you can use the way below

```go
intSlice.Count(func(i int)bool{return true})
```

### Default value

You can set the default value of the type with the optional argument `HasDefaultValue(stringValue string)`

For example

```go
 var l int64
 Linqablize( reflect.TypeOf(l), "linqable", HasDefaultValue("int64(88888)"))
```

### Method - Sum

#### Case1

C# - System.Linq

```C#
int sum = intList.Sum();
```

```C#
int sum = intList.Sum(i=>i);
```

Go - this package

```go
sum := intSlice.SumInt(func(i int) int { return i })
```

#### Case2

C# - System.Linq

```C#
double sum = intList.Sum(i => Convert.ToDouble(i));
```

Go - this package

```go
sum := intSlice.SumFloat64(func(i int) float64 { return float64(i) })
```
