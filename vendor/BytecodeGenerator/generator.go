/* Generates the bytecode from ast.File; see 'Generate' function
 *
 */

package BytecodeGenerator

import (
	"fmt"
	"go/ast"
	"unicode"
)

// Traverses ast.File and generates the bytecode representation
func Generate(f *ast.File) (bytecode []byte, err error) {
	bytecode = []byte{0xCA, 0xFE, 0xDE, 0xAD} // slice containing the bytecode data
	var descMap = make(map[string]string)     // map to convert from Golang description of types to gobc description. Example "int64" -> "I64"
	descMap["int64"] = "I64"
	descMap["float64"] = "F64"
	descMap["string"] = "S"
	descMap["nil"] = "V"

	// determine functionCount
	var functionCount uint64
	ast.Inspect(f, func(n ast.Node) bool { //uses ast.Inspect because this requires no specific visitor
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
	// walk AST to populate functionInfo structures
	Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case nil:

		case *ast.FuncDecl:
			idx += 1
			t := n.(*ast.FuncDecl)
			name := t.Name.Name
			// accessFlags
			if unicode.IsUpper([]rune(name)[0]) {
				funcs[idx] = append(funcs[idx], 0x01)
			} else {
				funcs[idx] = append(funcs[idx], 0x00)
			}
			// function name
			funcs[idx] = append(funcs[idx], Ui64tobyteslice(uint64(len(name)))...) // nameLength
			funcs[idx] = append(funcs[idx], []byte(name)...)                       // name

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
			funcs[idx] = append(funcs[idx], Ui64tobyteslice(uint64(len(desc)))...) // descriptorLength
			funcs[idx] = append(funcs[idx], []byte(desc)...)                       // descriptor

			// attributes
			var attributeCount uint32 = 1 //as of now, only 1 attribute: Code
			funcs[idx] = append(funcs[idx], Ui32tobyteslice(attributeCount)...)

			//Code attribute
			funcs[idx] = append(funcs[idx], 0x01) //attributeType

			//TODO: Populate the code attribute; convert instructions to bytecode

		default:

		}
		return true
	})
	for i := range funcs {
		bytecode = append(bytecode, funcs[i]...)
	}
	return bytecode, err
}
