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
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		os.Exit(-1)
	}

	fn := os.Args[1]
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file\n")
		os.Exit(-1)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fn, file, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}

	conf := types.Config{Importer: importer.Default()}

	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	// Inspect (DFS-walk) the AST
	// Anon. func is called on encounter of each node
	bv.Inspect(f, func(n ast.Node) bool {
		/*
			if debug {
				switch x := n.(type) {
				default:
					fmt.Printf("%T --> %+v\n", x, x)
				}
			}
		*/
		fmt.Printf("%T --> %+v\n", n, n)
		return true
	})

	fmt.Printf("Package  %q\n", pkg.Path())
	fmt.Printf("Name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
}
