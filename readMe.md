# Linqable



## Abstract

一個自我滿足的golang linq實現。

在golang支援泛型之前，很難有好用的linq，先前都是使用[ahmetb的版本](https://github.com/ahmetb/go-linq)，完成度算是很高了，但依然有許多不便之處，像是`Query`型別使型別判定很不方便(因為回傳的也是`Query`)，另外還有委派的型別判定讓執行效率變得奇差。

於是就起了念頭把自己常用的linq拿出來寫，不常用的就跳過，於是就有了這個專案。



## Feature

**Advantage**

* source code 生成，使用完後可以把import移除
* 不在function中進行型別判定，堪比for loop的執行速度
* 強型別，不會看到像是`Query`這種裡面不知道裝甚麼的`struct`，委派需要的方法簽章也完全固定，無需顧慮該傳入何種方法。

**Disadvantage**

* 目前僅支援Slice結構
* 目前不打算收入所有的linq方法
* 內部錯誤會以panic呈現





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



### Delete the imports and start linq

~~import "reflect"~~

~~import "github.com/STRockefeller/linqable"~~



```go
	nums:=NewLinqableInt([]int{1,2,3,4,5,6,7,8,9,0})
	firstThreeOddNums := nums.Where(func(i int) bool {return i%2==1}).Take(3).ToSlice()
```



## Methods

現階段支援的方法

* Where
* Take
* TakeWhile
* Skip
* SkipWhile
* ToSlice

