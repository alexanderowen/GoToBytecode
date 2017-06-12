package main

import (
	bg "BytecodeGenerator"
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
)

var DEBUG bool = false

func main() {
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

	// print AST node information
	if DEBUG {
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

	// generate bytecode
	bytecode, err := bg.Generate(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating bytecode:\n\t")
		log.Fatal(err)
		os.Exit(-1)
	}

	// output to .gobc file
	outFn := strings.Split(fn, ".go")[0] + ".gobc"
	ioutil.WriteFile(outFn, bytecode, 0777)
}
