package generator

import (
	"io"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type Generator struct {
	Spec                Spec
	PointerReciever     bool
	MethodPrefix        string
	ForcePrivateMethods bool
}

func (gen Generator) MethodGenerators() map[string]func() *Statement {
	return map[string]func() *Statement{
		"GetFieldNames": gen.GenerateGetFieldNames,
		"FieldsN":       gen.GenerateFieldsN,
		"GetField":      gen.GenerateGetField,
	}
}

func (gen Generator) Render(wr io.Writer, pkg string, generators ...func() *Statement) error {
	if len(generators) == 0 {
		for _, generator := range gen.MethodGenerators() {
			generators = append(generators, generator)
		}
	}
	var file = NewFile(pkg)
	for _, methodGen := range generators {
		file.Add(methodGen().Line())
	}
	return file.Render(wr)
}

func (gen Generator) recvArgs() *Statement {
	if gen.PointerReciever {
		return gen.recv().Op("*").Id(gen.Spec.Name)
	}
	return gen.recv().Id(gen.Spec.Name)
}

func (gen Generator) recv() *Statement {
	return Id("_" + gen.Spec.Name)
}

func (gen Generator) methodName(name string) string {
	if !gen.ForcePrivateMethods {
		name = strings.Title(name)
	}
	return gen.MethodPrefix + name
}
