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

	for _, filePath := range os.Args[2:] {
		if filePath == "--" { // Needed for `go run ... -- file`
			continue
		}

		if strings.HasSuffix(filePath, "/...") {
			dirName := filePath[:len(filePath)-4]
			fInfo, err := os.Stat(dirName)
			if err != nil {
				log.Fatalf("error with filepath %s: %s", filePath, err)
			}
			if !fInfo.IsDir() {
				log.Fatalf("error with filepath %s: %s is not a directory", filePath, dirName)
			}
			walkDir(dirName)
		} else {
			walkAst(filePath)
		}
	}

	for _, declaredFunc := range declaredFuncs {
		// fmt.Printf("DECLARED: %s\n", declaredFunc)
		if _, ok := binaryFuncs[declaredFunc]; !ok {
			fmt.Fprintf(os.Stderr, "%s is not used\n", declaredFunc)
		}
	}
}

func walkAst(filePath string) {
	v := visitor{fset: token.NewFileSet()}
	// fmt.Printf("PROVIDED ARG: %s\n", filePath)
	curFilePath = filepath.Dir(filePath)

	f, err := parser.ParseFile(v.fset, filePath, nil, 0)
	if err != nil {
		log.Fatalf("failed to parse file %s: %s", filePath, err)
	}
	if curFilePath == "." || f.Name.Name == "main" { // If in root, use package name
		curFilePath = f.Name.Name
	}

	ast.Walk(&v, f)
}

func walkDir(dirName string) {
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") {
			walkAst(path)
		}
		return err
	})
	if err != nil {
		log.Fatalf("error walking directory %s: %s", dirName, err)
	}
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	if fdecl, ok := node.(*ast.FuncDecl); ok {
		if fdecl.Recv != nil {
			// fmt.Printf("RECEIVER: %s\n", fdecl.Recv)
			if recv, ok := fdecl.Recv.List[0].Type.(*ast.StarExpr); ok {
				if recvType, ok := recv.X.(*ast.Ident); ok { //
					if curFilePath == "main" {
						declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s.(*%s).%s", curFilePath, recvType.Name, fdecl.Name.Name)) // Pointer receivers
					} else {
						declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s/%s.(*%s).%s", projectName, curFilePath, recvType.Name, fdecl.Name.Name)) // Pointer receivers
					}
				}
			}
			if recvIdent, ok := fdecl.Recv.List[0].Type.(*ast.Ident); ok {
				if curFilePath == "main" {
					declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s.%s.%s", curFilePath, recvIdent.Name, fdecl.Name.Name)) // Receivers
				} else {
					declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s/%s.%s.%s", projectName, curFilePath, recvIdent.Name, fdecl.Name.Name)) // Receivers
				}
			}
		} else {
			if !strings.HasPrefix(fdecl.Name.Name, "Test") && fdecl.Name.Name != "main" && fdecl.Name.Name != "init" { // Ignore Test, main and init functions
				if curFilePath == "main" {
					declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s.%s", curFilePath, fdecl.Name.Name)) // Functions
				} else {
					declaredFuncs = append(declaredFuncs, fmt.Sprintf("%s/%s.%s", projectName, curFilePath, fdecl.Name.Name)) // Functions
				}
				// fmt.Printf("Found function declaration: %s/%s.%s\n", projectName, curFilePath, fdecl.Name.Name)
			}
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
