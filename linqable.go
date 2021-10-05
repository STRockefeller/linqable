package linqable

import (
	"fmt"
	"os"
	"reflect"

	"github.com/dave/jennifer/jen"
)

// Linqablize : generate a new `.go` file with specified type
//
// Parameters
//
// t reflect.Type               => the specified type
//
// packageName string           => package name of the `.go` file
//
// opts ...LinqablizeOptionFunc => optional parameters
//
// Optional Parameters
//
// IsImportedType() => output file will import the package of the specified type
//
// IsNumeric() => output type will contains the method Sum() Max() Min()
func Linqablize(t reflect.Type, packageName string, opts ...LinqablizeOptionFunc) {
	var opt linqablizeOption
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	linqableTypeName := "Linqable" + t.Name()
	typeName := t.Name()

	jenFile := jen.NewFile(packageName)
	jenFile.HeaderComment(`Code generated by linqable.go Do NOT EDIT.`)
	// #region imported type
	if opt.isImportedType {
		typeName = t.String()
		jenFile.Id("import").Id("\"" + t.PkgPath() + "\"").Line()
	}
	// #endregion imported type

	predicateCode := jen.Id("predicate").Func().Call(jen.Id(typeName)).Id("bool")

	jenFile.Line()
	jenFile.Id("type").Id(linqableTypeName).Op("[]").Id(typeName)
	jenFile.Line()

	// #region constructor
	jenFile.Func().Id(fmt.Sprintf("NewLinqable%s", t.Name())).Call(jen.Id("si").Op("[]").Id(typeName)).Id(linqableTypeName).Block(jen.Return(jen.Id("si")))
	// #endregion constructor
	jenFile.Line()

	// #region Where
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Where").Call(predicateCode).Id(linqableTypeName).
		Block(jen.Id("res").Op(":=").Op("[]").Id(typeName).Op("{}").Line().For(jen.Id("_").Op(",").Id("i").Op(":=").Id("range").Id("si").
			Block(jen.If(jen.Id("predicate(i)")).
				Block(jen.Id("res").Op("=").Id("append(res, i)"))).Line().Return(jen.Id("res"))))

	// #endregion Where
	jenFile.Line()

	// #region Contains
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Contains").Call(jen.Id("target").Id(typeName)).Id("bool").
		Block(jen.For(jen.Id("_").Op(",").Id("i").Op(":=").Id("range").Id("si").
			Block(jen.If(jen.Id("i == target")).
				Block(jen.Return(jen.Id("true")))).Line().Return(jen.Id("false"))))

	// #endregion Contains
	jenFile.Line()

	// #region Count
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Count").Call(predicateCode).Id("int").
		Block(jen.Id("var count int").Line().For(jen.Id("_").Op(",").Id("i").Op(":=").Id("range").Id("si").
			Block(jen.If(jen.Id("predicate(i)")).
				Block(jen.Id("count++"))).Line().Return(jen.Id("count"))))

	// #endregion Count
	jenFile.Line()

	// #region Any
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Any").Call(predicateCode).Id("bool").
		Block(jen.For(jen.Id("_").Op(",").Id("i").Op(":=").Id("range").Id("si").
			Block(jen.If(jen.Id("predicate(i)")).
				Block(jen.Return(jen.Id("true")))).Line().Return(jen.Id("false"))))

	// #endregion Any
	jenFile.Line()

	// #region All
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("All").Call(predicateCode).Id("bool").
		Block(jen.For(jen.Id("_").Op(",").Id("i").Op(":=").Id("range").Id("si").
			Block(jen.If(jen.Id("predicate(i)")).
				Block(jen.Continue()).Else().
				Block(jen.Return(jen.Id("false")))).Line().Return(jen.Id("true"))))

	// #endregion All
	jenFile.Line()

	// #region Take
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Take").Call(jen.Id("n int")).Id(linqableTypeName).
		Block(jen.If(jen.Id("n < 0 || n > len(si)")).
			Block(jen.Panic(jen.Id(`"Linq: Take() out of index"`))).Line().Id("res").Op(":=").Op("[]").Id(typeName).Op("{}").Line().For(jen.Id("i := 0; i < n; i++")).
			Block(jen.Id("res = append(res, si[i])")).Line().Return(jen.Id("res")))
	// #endregion Take
	jenFile.Line()

	// #region TakeWhile
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("TakeWhile").Call(predicateCode).Id(linqableTypeName).
		Block(jen.Id("res").Op(":=").Op("[]").Id(typeName).Op("{}").Line().For(jen.Id("i := 0; i < len(si); i++")).
			Block(jen.If(jen.Id("predicate(si[i])").
				Block(jen.Id("res = append(res, si[i])")).Else().
				Block(jen.Return(jen.Id("res"))))).Line().Return(jen.Id("res")))
	// #endregion TakeWhile
	jenFile.Line()

	// #region TakeLast
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("TakeLast").Call(jen.Id("n int")).Id(linqableTypeName).
		Block(jen.If(jen.Id("n < 0 || n > len(si)")).
			Block(jen.Panic(jen.Id(`"Linq: TakeLast() out of index"`))).Line().Return(jen.Id("si.Skip(len(si) - n)")))
	// #endregion TakeLast
	jenFile.Line()

	// #region Skip
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("Skip").Call(jen.Id("n int")).Id(linqableTypeName).
		Block(jen.If(jen.Id("n < 0 || n > len(si)")).
			Block(jen.Panic(jen.Id(`"Linq: Skip() out of index"`))).Line().Return(jen.Id("si[n:]")))
	// #endregion Skip
	jenFile.Line()

	// #region SkipWhile
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("SkipWhile").Call(predicateCode).Id(linqableTypeName).
		Block(jen.For(jen.Id("i := 0; i < len(si); i++")).
			Block(jen.If(jen.Id("predicate(si[i])").
				Block(jen.Op("continue")).Else().
				Block(jen.Return(jen.Id("si[i:]"))))).Line().Return(jen.Id(fmt.Sprintf("%s{}", linqableTypeName))))
	// #endregion SkipWhile
	jenFile.Line()

	// #region SkipLast
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("SkipLast").Call(jen.Id("n int")).Id(linqableTypeName).
		Block(jen.If(jen.Id("n < 0 || n > len(si)")).
			Block(jen.Panic(jen.Id(`"Linq: SkipLast() out of index"`))).Line().Return(jen.Id("si.Take(len(si) - n)")))
	// #endregion SkipLast
	jenFile.Line()

	// #region ToSlice
	jenFile.Func().Call(jen.Id("si").Id(linqableTypeName)).Id("ToSlice").Call().Op("[]").Id(typeName).Block(jen.Return(jen.Id("si")))
	// #endregion ToSlice

	file, err := os.Create(fmt.Sprintf("linqable_%s.go", t.Name()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%#v", jenFile))
}

// LinqablizeOptionFunc : optional parameters of Linqablize()
type LinqablizeOptionFunc func(*linqablizeOption)
type linqablizeOption struct {
	isImportedType bool
	isNumeric      bool
}

// IsImportedType : optional parameter for the Linqablize()
//
// output file will import the package of the specified type
func IsImportedType() LinqablizeOptionFunc {
	return func(lo *linqablizeOption) {
		lo.isImportedType = true
	}
}

// IsNumeric : optional parameter for the Linqablize()
//
// output type will contains the method Sum() Max() Min()
func IsNumeric() LinqablizeOptionFunc {
	return func(lo *linqablizeOption) {
		lo.isNumeric = true
	}
}
