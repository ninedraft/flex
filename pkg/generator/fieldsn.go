package generator

import (
	"strings"

	. "github.com/dave/jennifer/jen"
)

func (gen Generator) GenerateFieldsN() *Statement {
	var methodName = "fieldsN"
	if !gen.ForcePrivateMethods {
		methodName = strings.Title(methodName)
	}
	var N = Lit(gen.Spec.FieldsN())
	var method = Func().Params(gen.recv()).Id(methodName).Params().Int().Block(Return(N))
	return method
}
