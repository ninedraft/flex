package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func (gen Generator) GenerateGetField() *Statement {
	var methodName = gen.methodName("getField")
	var doc = Commentf("%s returns field value by name. Panics if field not found", methodName).Line()
	var prefix = fmt.Sprintf("[%s.%s]: field ", gen.Spec.Name, methodName)
	var panicMsg = Lit(prefix).Op("+").Qual("strconv", "Quote").Call(Id("name")).Op("+").Lit(" not found")
	var cases []Code
	for _, field := range gen.Spec.FieldNames() {
		cases = append(cases, Case(Lit(field)).Block(Return(gen.recv().Dot(field))))
	}
	var switchBlock = Switch(Id("name")).Block(append(cases, Default().Block(Panic(panicMsg)))...)
	var args = Id("name").String()
	var method = doc.Func().Params(gen.recvArgs()).Id(methodName).Params(args).Interface().Block(switchBlock)
	return method
}
