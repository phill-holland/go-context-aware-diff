package main

import (
	"fmt"
	pp "parser"
)

func main() {

	pp.Load("hello.txt")
	//var p parser.Instruction;
	//pp.load("hello.txt")
	fmt.Printf("Hello World!\n")
}

// > go mod init parser
