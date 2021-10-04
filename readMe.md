# Linqable



## Abstract

一個自我滿足的golang linq實現。

在golang支援泛型之前，很難有好用的linq，先前都是使用ahmetb的版本，完成度算是很高了，但依然有許多不便之處，像是`Query`型別使型別判定很不方便(因為回傳的也是`Query`)，另外還有委派的型別判定讓執行效率變得奇差。

於是就起了念頭把自己常用的linq拿出來寫，不常用的就跳過，於是就有了這個專案。

## Installation

```powershell
go get github.com/STRockefeller/linqable
```



## Methods

現階段支援的方法

* Where
* Take
* Skip
* ToSlice