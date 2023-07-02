package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

type visitor struct {
	fset *token.FileSet
}

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // Needed for `go run ... -- file`
			continue
		}

		fmt.Printf("PROVIDED ARG: %s\n", filePath)

		f, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	if fdecl, ok := node.(*ast.FuncDecl); ok {
		fmt.Printf("Found function declaration: %s\n", fdecl.Name.Name)
	}

	return v
}
