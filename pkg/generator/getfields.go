package generator

import (
	. "github.com/dave/jennifer/jen"
)

func (gen Generator) GenerateGetFields() *Statement {
	var methodName = gen.methodName("getFields")
	var doc = Commentf("%s returns fields values by names. Panics if field not found", methodName).Line()
	var args = Id("names").Op("...").String()
	var valuesDecl = Var().Id("values").Op("=").Make(Index().Interface(), Lit(0), Len(Id("names")))
	var getValue = gen.recv().Dot(gen.methodName("getField"))
	var method = doc.Func().Params(gen.recvArgs()).Id(methodName).Params(args).Interface().Block(
		valuesDecl,
		For(Id("_").Op(",").Id("name").Op(":=").Range().Id("names").Block(
			Id("values").Op("=").Append(Id("values"), getValue.Call(Id("name")))),
		),
		Return(Id("values")),
	)
	return method
}
