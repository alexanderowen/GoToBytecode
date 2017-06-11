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

var DEBUG bool = false
var descMap = make(map[string]string) // map to convert from Golang description of types to gobc description. Example "int64" -> "I64"

func main() {
	descMap["int64"] = "I64"
	descMap["float64"] = "F64"
	descMap["string"] = "S"
	descMap["nil"] = "V"

	// check command usage
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		os.Exit(-1)
	}

	// open argv[1]
	fn := os.Args[1]
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file\n")
		os.Exit(-1)
	}

	// parse the file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fn, file, 0)
	if err != nil {
		log.Fatal(err) // parse error
		os.Exit(-1)
	}

	// type check the file
	conf := types.Config{Importer: importer.Default()}
	_, err = conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Type Error:\n\t")
		log.Fatal(err) // type error
		os.Exit(-1)
	}

	// begin generating bytecode
	var bytecode []byte = []byte{0xCA, 0xFE, 0xDE, 0xAD} // slice containing the bytecode data

	if DEBUG {
		// print out AST node information
		ast.Inspect(f, func(n ast.Node) bool {
			switch n.(type) {
			case nil:
				// nop
			default:
				fmt.Printf("%T --> %+v\n", n, n)
			}
			return true
		})
	}

	// determine function count
	var functionCount uint64
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			functionCount += 1
		default:
			// nop
		}
		return true
	})
	bytecode = append(bytecode, byte(functionCount))

	// make the functionInfo structures
	var funcs [][]byte = make([][]byte, functionCount)
	for i := range funcs {
		funcs[i] = make([]byte, 0)
	}
	var idx int = -1
	bv.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case nil:

		case *ast.FuncDecl:
			idx += 1
			t := n.(*ast.FuncDecl)
			name := t.Name.Name
			// access flags
			if unicode.IsUpper([]rune(name)[0]) {
				funcs[idx] = append(funcs[idx], 0x01)
			} else {
				funcs[idx] = append(funcs[idx], 0x00)
			}
			// function name
			funcs[idx] = append(funcs[idx], bv.Ui64tobyteslice(uint64(len(name)))...) // length of name
			funcs[idx] = append(funcs[idx], []byte(name)...)                          // name

			// function descriptor
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
			funcs[idx] = append(funcs[idx], bv.Ui64tobyteslice(uint64(len(desc)))...) // length of descriptor
			funcs[idx] = append(funcs[idx], []byte(desc)...)                          // descriptor

			var attributeCount uint32 = 1 //as of now, only 1 attribute: Code
			funcs[idx] = append(funcs[idx], bv.Ui32tobyteslice(attributeCount)...)

			//Code attribute
			funcs[idx] = append(funcs[idx], 0x01) //attributeType

		default:

		}
		return true
	})
	for i := range funcs {
		bytecode = append(bytecode, funcs[i]...)
	}

	// output to .gobc file
	outFn := strings.Split(fn, ".go")[0] + ".gobc"
	ioutil.WriteFile(outFn, bytecode, 0777)
}
