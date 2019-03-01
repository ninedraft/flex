package generator

import "go/ast"

type StructReflex struct {
	Comments   []string
	typeSpec   *ast.TypeSpec
	structSpec *ast.StructType
}
