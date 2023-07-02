package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BreD1810/go-test-only-linter/binary"
)

type visitor struct {
	fset *token.FileSet
}

// var curPackage string
var curFilePath string
var projectName string
var declaredFuncs []string

func main() {
	projectName = getProjectName()

	buildLocation := os.Args[1]
	binaryFuncs := binary.GetBinaryFunctions(buildLocation)
	// fmt.Printf("%+q\n\n", binaryFuncs)

	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[2:] {
		if filePath == "--" { // Needed for `go run ... -- file`
			continue
		}

		// fmt.Printf("PROVIDED ARG: %s\n", filePath)
		curFilePath = filepath.Dir(filePath)

		f, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}

	for _, declaredFunc := range declaredFuncs {
		// fmt.Printf("DECLARED: %s\n", declaredFunc)
		if _, ok := binaryFuncs[declaredFunc]; !ok {
			fmt.Fprintf(os.Stderr, "%s is not used\n", declaredFunc)
		}
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	// if file, ok := node.(*ast.File); ok {
	// 	fmt.Printf("Found package: %s\n", file.Name.Name)
	// 	curPackage = file.Name.Name
	// }
	if fdecl, ok := node.(*ast.FuncDecl); ok {
		if fdecl.Recv != nil {
			// fmt.Printf("RECEIVER: %s\n", fdecl.Recv)
			if recv, ok := fdecl.Recv.List[0].Type.(*ast.StarExpr); ok {
				if recvType, ok := recv.X.(*ast.Ident); ok {
					declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s/%s.(*%s).%s", projectName, curFilePath, recvType.Name, fdecl.Name.Name))
				}
			}
			// fdecl.Recv.List[0].Type.(*ast.StarExpr).Name
			// switch f := fdecl.Recv.(type) {
			// case *ast.Ident:
			// 	fmt.Printf("NAME: %s\n", f)
			// }
		} else {
			declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s/%s.%s", projectName, curFilePath, fdecl.Name.Name))
			// fmt.Printf("Found function declaration: %s/%s.%s\n", projectName, curFilePath, fdecl.Name.Name)
		}
	}

	return v
}
func getProjectName() string {
	f, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("failed to read go.mod: %s", err)
	}
	lines := strings.Split(string(f), "\n")
	if len(lines) < 1 {
		log.Fatalf("go.mod is empty")
	}

	return strings.TrimPrefix(lines[0], "module ")
}
