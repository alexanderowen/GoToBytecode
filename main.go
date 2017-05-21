package main

import (
	bv "BytecodeVisitor"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

var descMap = make(map[string]string)

func i64tobyteslice(x int64) []byte {
	b := make([]byte, 8)
	b[7] = byte(x)
	b[6] = byte(x >> 8)
	b[5] = byte(x >> 16)
	b[4] = byte(x >> 24)
	b[3] = byte(x >> 32)
	b[2] = byte(x >> 40)
	b[1] = byte(x >> 48)
	b[0] = byte(x >> 56)
	return b
}

func main() {
	descMap["int64"] = "I64"
	descMap["float64"] = "F64"
	descMap["string"] = "S"
	descMap["nil"] = "V"

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		os.Exit(-1)
	}

	fn := os.Args[1]
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file\n")
		os.Exit(-1)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fn, file, 0)
	if err != nil {
		log.Fatal(err) // parse error
		os.Exit(-1)
	}

	conf := types.Config{Importer: importer.Default()}

	//pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	_, err = conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
		os.Exit(-1)
	}

	// Inspect (DFS-walk) the AST
	// Anon. func is called on encounter of each node
	var bytecode []byte = []byte{0xCA, 0xFE, 0xDE, 0xAD}
	var functionCount uint64
	//bv.Inspect(f, func(n ast.Node) bool {
	ast.Inspect(f, func(n ast.Node) bool {
		var flag bool
		switch n.(type) {
		case nil:
			flag = true
		case *ast.FuncDecl:
			functionCount += 1
		default:

		}
		if !flag {
			fmt.Printf("%T --> %+v\n", n, n)
		}
		return true
	})
	bytecode = append(bytecode, byte(functionCount))

	// Make the functionInfo structures
	var funcs [][]byte = make([][]byte, functionCount)
	for i := range funcs {
		funcs[i] = make([]byte, 0) //make room for 30 instructions
	}
	var idx int = -1
	bv.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case nil:

		case *ast.FuncDecl:
			idx += 1
			t := n.(*ast.FuncDecl)
			name := t.Name.Name
			if unicode.IsUpper([]rune(name)[0]) { // access flags
				funcs[idx] = append(funcs[idx], 0x01)
			} else {
				funcs[idx] = append(funcs[idx], 0x00)
			}
			funcs[idx] = append(funcs[idx], i64tobyteslice(int64(len(name)))...) // length of name
			funcs[idx] = append(funcs[idx], []byte(name)...)

			// descriptor
			typ := t.Type
			desc := "("
			if typ.Params.List == nil {
				desc += "V"
			} else {
				for i := range typ.Params.List {
					s := fmt.Sprintf("%s", typ.Params.List[i].Type)
					desc += descMap[s]
				}
			}
			desc += ")"
			if typ.Results == nil {
				desc += "V"
			} else {
				for i := range typ.Results.List {
					s := fmt.Sprintf("%s", typ.Results.List[i].Type)
					desc += descMap[s]
				}
			}
			funcs[idx] = append(funcs[idx], []byte(desc)...)

		default:

		}
		return true
	})
	for i := range funcs {
		bytecode = append(bytecode, funcs[i]...)
	}

	outFn := strings.Split(fn, ".go")[0] + ".gobc"
	ioutil.WriteFile(outFn, bytecode, 0777)

	/*
		fmt.Printf("Package  %q\n", pkg.Path())
		fmt.Printf("Name:    %s\n", pkg.Name())
		fmt.Printf("Imports: %s\n", pkg.Imports())
		fmt.Printf("Scope:   %s\n", pkg.Scope())
	*/
}
