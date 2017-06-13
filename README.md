## Go (.go) to Bytecode (.gobc)  
A tool to convert Go source code to a bytecode specification that can be executed on a stack-based VM.  

### Current Progress  
Documents:
- GoDataTypes: enumerates Go's basic types and their opeartions, as well as helpful information 
- BytecodeFileFormat: defines a basic file format for the .gobc file   
- GoBytecodeInstructions: defines basic instructions for the bytecode 
 
Current goal is to emit bytecode for basic operations on int64, float64, and string.  
Execution of 'main.go' will emit the magicNumber, functionCount, and the first few fields of the Code attribute.  
### TODO list  
- Basic operations in Go need to be converted into bytecode (for example, '5 + 6')  
    - Let 'x' and 'y' be two non-const variables of type int64. Convert 'x + y' to bytecode.
    - Note that Go handles constant expressions in a peculiar way. Refer to https://blog.golang.org/constants to understand how Go handles constant expressions. Simply put, Go evaluates constant expressions at compile-time, so the bytecode instruction should NOT be 'i64add 5 6'.  

- (Long term) A few of Go's more unique features have not been discussed in transition to bytecode. To name a few
    - First class functions
    - Goroutines
    - Interfaces
    - The '...' variadic parameter

### Commands
- main.go needs to be built with 'go build'  
`go build main.go`  
`./main filename.go`  

- Examine the bytes of the .gobc file with xxd  
`xxd -g 1 filename.gobc`

### Relevant Links
- Golang's Source Repository, branch for go1.8: https://github.com/golang/go/tree/release-branch.go1.8  
- Golang's Language Specification: https://golang.org/ref/spec  
- Go compiler, Go in Go: https://talks.golang.org/2015/gogo.slide#1  
- Java's bytecode file format: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html   
- Documentation concerning go/types, Go's type checker: https://github.com/golang/example/tree/master/gotypes   

### Version  
- go1.8.3 

