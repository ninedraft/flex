package generator

import (
	"go/ast"
	"reflect"
	"unicode"
	"unicode/utf8"
)

type StructReflex struct {
	Comments   []string
	Fields     []Field
	typeSpec   *ast.TypeSpec
	structSpec *ast.StructType
}

func StructReflexFromAST(typeSpec *ast.TypeSpec) StructReflex {
	var fields []Field
	var structSpec = typeSpec.Type.(*ast.StructType)
	for _, field := range structSpec.Fields.List {
		for _, name := range field.Names {
			if FieldIsExported(name.Name) {
				// TODO: (ninedraft) Add type extraction and put type info to TypeNameField
				fields = append(fields, Field{
					Name:     name.Name,
					TypeName: "TODO!",
					Comments: commentsFromAST(field.Doc.List),
					Tag:      reflect.StructTag(field.Tag.Value),
				})
			}
		}
	}
	return StructReflex{
		Comments: commentsFromAST(typeSpec.Doc.List),
		Fields:   fields,
		typeSpec: typeSpec,
	}
}

type Field struct {
	Name     string
	TypeName string
	Comments []string
	Tag      reflect.StructTag
}

func (structReflex StructReflex) Name() string {
	return structReflex.typeSpec.Name.Name
}

func (structReflex StructReflex) ForEachField(walker func(field *ast.Field)) {
	for _, field := range structReflex.structSpec.Fields.List {
		walker(field)
	}
}

func (structReflex StructReflex) ExportedFieldNum() int {
	var n int
	structReflex.ForEachField(func(field *ast.Field) {
		n++
	})
	return n
}

func (structReflex StructReflex) ExportedFieldNames() []string {
	var names []string
	structReflex.ForEachField(func(field *ast.Field) {
		names = append(names, field.Names[0].Name)
	})
	return names
}

func FieldIsExported(name string) bool {
	var firstRune, _ = utf8.DecodeLastRuneInString(name)
	return unicode.IsUpper(firstRune)
}
