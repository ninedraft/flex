package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"testing"
)

func TestGenerator(test *testing.T) {
	var generator = Generator{
		Spec: Spec{
			Name: "User",
			Fields: []StructField{
				{Name: "Name", Type: "string"},
				{Name: "Age", Type: "int"},
			},
		},
	}
	var methodsToGenerate = generator.MethodGenerators()
	for methodName, methodGenerator := range methodsToGenerate {
		methodGenerator := methodGenerator
		test.Run(methodName, func(test *testing.T) {
			var buf = &bytes.Buffer{}
			if err := methodGenerator().Render(buf); err != nil {
				test.Log(lineNumbers("\n" + buf.String()))
				test.Fatal(err)
			}
			if err := checkAST(buf.String()); err != nil {
				test.Log(lineNumbers("\n" + buf.String()))
				test.Fatal(err)
			}
			test.Log(buf)
		})
	}
}

func lineNumbers(text string) string {
	var buf = &bytes.Buffer{}
	for i, line := range strings.Split(text, "\n") {
		fmt.Fprintf(buf, "%30d %s\n", i, line)
	}
	return buf.String()
}

func checkAST(text string) error {
	var _, err = format.Source([]byte(text))
	return err
}
