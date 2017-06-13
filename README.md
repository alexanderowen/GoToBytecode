## Go (.go) to Bytecode (.gobc)  
A tool to convert Go source code to a bytecode specification that can be executed on a stack-based VM.  

### Current Progress  
Documents:
- GoDataTypes: enumerates Go's basic types and their opeartions  
- BytecodeFileFormat: defines a basic file format for the .gobc file   
- GoBytecodeInstructions: defines basic instructions for the bytecode 
 
Current goal is to emit bytecode for basic operations on int64, float64, and string.  
Execution of 'main.go' will emit the magicNumber, functionCount, and the first few fields of the Code attribute.  
### TODO list  
- Basic operations in Go need to be converted into bytecode (for example, '5 + 6')  
    - However, Go handles constant expressions in a peculiar way. Refer to https://blog.golang.org/constants to understand how Go handles constant expressions. Simply put, Go evaluates constant expressions at compile-time, so the bytecode instruction should NOT be 'i64add 5 6'.  

### Commands
- main.go needs to be built with 'go build'  
`go build main.go`  
`./main filename.go`  

- Examine the bytes of the .gobc file with xxd  
`xxd -g 1 filename.gobc`

### Version  
- go1.8.3 : https://github.com/golang/go/tree/release-branch.go1.8

