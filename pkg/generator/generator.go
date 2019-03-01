package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type Generator struct {
	ast     ast.Node
	structs map[string]StructReflex
}

func GeneratorFromAST(AST ast.Node) *Generator {
	return &Generator{
		ast:     AST,
		structs: extractStructsFromAST(AST),
	}
}

func ParseFile(fileContent string) (*Generator, error) {
	var fileSet = token.NewFileSet()
	var AST, errParseFile = parser.ParseFile(fileSet, "file.go",
		fileContent, parser.ParseComments)
	if errParseFile != nil {
		return nil, errParseFile
	}
	return GeneratorFromAST(AST), nil
}

func ParsePackageDir(dirPath, packageName string) (*Generator, error) {
	var fileSet = token.NewFileSet()
	var fileFilter = func(fileInfo os.FileInfo) bool {
		var ext = filepath.Ext(fileInfo.Name())
		return ext == ".go"
	}
	var packages, errParseDir = parser.ParseDir(fileSet, dirPath,
		fileFilter, parser.ParseComments)
	if errParseDir != nil {
		return nil, errParseDir
	}
	var pkg, pkgExists = packages[packageName]
	if !pkgExists {
		return nil, fmt.Errorf("generator: package %q doesn't exist in dir %q", packageName, dirPath)
	}
	return GeneratorFromAST(pkg), nil
}

func extractStructsFromAST(AST ast.Node) map[string]StructReflex {
	var structs = make(map[string]StructReflex)
	ast.Inspect(AST, func(node ast.Node) bool {
		var typeSpec, isTypeSpec = node.(*ast.TypeSpec)
		if !isTypeSpec {
			return true
		}
		var structSpec, isStructSpec = typeSpec.Type.(*ast.StructType)
		if !isStructSpec {
			return true
		}
		structs[typeSpec.Name.Name] = StructReflex{
			Comments:   commentsFromAST(typeSpec.Doc.List),
			typeSpec:   typeSpec,
			structSpec: structSpec,
		}
		return true
	})
	return structs
}

func commentsFromAST(commentList []*ast.Comment) []string {
	var comments = make([]string, 0, len(commentList))
	for _, comment := range commentList {
		comments = append(comments, comment.Text)
	}
	return comments
}
