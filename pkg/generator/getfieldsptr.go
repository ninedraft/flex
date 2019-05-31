package generator

import (
	"fmt"
	"reflect"

	. "github.com/dave/jennifer/jen"
)

func (gen Generator) GenerateGetFieldPtr() *Statement {
	var methodName = gen.methodName("getFieldPtr")
	var doc = Commentf("%s returns field pointer by name. Panics if field not found.\n"+
		"If field is nil and not an interface, then allocates value of corresponding type, assignes it to field and returns pointer to it.", methodName).Line()
	// (name string)
	var args = Id("name").String()
	// func(reciever Type) GetFieldPtr() interface{}
	var methodSign = doc.Func().Params(gen.recvArgs()).Id(methodName).Params(args).Interface()

	var cases []Code
	for _, field := range gen.Spec.Fields {
		field := field
		// reciever.FieldName
		var fieldLiteral = func() *Statement { return gen.recv().Dot(field.Name) }
		var getField = Do(func(st *Statement) {
			switch field.Kind {
			case reflect.Ptr:
				st.If(fieldLiteral().Op("==").Nil()).Block(
					Add(fieldLiteral().Op("=").New(Id(field.Type))),
				).Line().Return(fieldLiteral())
			case reflect.Interface:
				st.Return(Op("&").Add(fieldLiteral()))
			default:
				st.Return(fieldLiteral())
			}
		})
		cases = append(cases, Case(Lit(field.Name)).Block(getField))
	}

	/*
		default:
			panic(field not found!)
		}
	*/
	{
		var prefix = fmt.Sprintf("[%s.%s]: field ", gen.Spec.Name, methodName)
		var panicMsg = Lit(prefix).Op("+").Qual("strconv", "Quote").Call(Id("name")).Op("+").Lit(" not found")
		cases = append(cases, Default().Block(Panic(panicMsg)))
	}
	return methodSign.Block(Switch(Id("name")).Block(cases...))
}
