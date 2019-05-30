package generator

import (
	"strings"

	. "github.com/dave/jennifer/jen"
)

func (gen Generator) GenerateGetFieldNames() *Statement {
	var methodName = "getFieldNames"
	if !gen.ForcePrivateMethods {
		methodName = strings.Title(methodName)
	}
	var names = func(group *Group) {
		for _, name := range gen.Spec.FieldNames() {
			group.Lit(name)
		}
	}
	var method = Func().Params(gen.recv()).Id(methodName).Params().Index().String().Block(
		Return(Index().String().ValuesFunc(names)))
	return method
}
