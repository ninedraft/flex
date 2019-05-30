package generator

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

type Spec struct {
	Name   string
	Fields []StructField
	AST    ast.Node
}

func NewSpec(node ast.TypeSpec) Spec {
	var structSpec = node.Type.(*ast.StructType)
	var spec = Spec{
		Name:   node.Name.Name,
		Fields: getFields(structSpec),
		AST:    structSpec,
	}
	return spec
}

func getFields(structNode *ast.StructType) []StructField {
	if structNode.Fields == nil {
		return []StructField{}
	}
	var fields = make([]StructField, 0, structNode.Fields.NumFields())
	for _, field := range structNode.Fields.List {
		var fieldStruct, fieldIsStruct = field.Type.(*ast.StructType)
		switch {
		case len(field.Names) == 0 && fieldIsStruct:
			fields = append(fields, getFields(fieldStruct)...)
		case len(field.Names) != 0:
			for _, name := range field.Names {
				if name.IsExported() {
					fields = append(fields, StructField{
						Name: name.Name,
						Type: printNode(field.Type),
						AST:  field.Type,
					})
				}
			}
		}
	}
	return fields
}

func (spec Spec) FieldsN() int {
	return len(spec.Fields)
}

func (spec Spec) FieldNames() []string {
	var names = make([]string, 0, spec.FieldsN())
	for _, field := range spec.Fields {
		names = append(names, field.Name)
	}
	return names
}

type StructField struct {
	Name   string
	Type   string
	Tag    map[string][]string
	RawTag string
	AST    ast.Node
}

func printNode(node ast.Node) string {
	var buf = &bytes.Buffer{}
	var fset = token.NewFileSet()
	printer.Fprint(buf, fset, node)
	return buf.String()
}
