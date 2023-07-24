package main

import (
	"fmt"
	pp "parser"
)

func main() {

	p1 := pp.Load("program1.txt")
	p2 := pp.Load("program2.txt")

	result := p1.Compare(p2)

	fmt.Printf(result.Print())
	//var p parser.Instruction;
	//pp.load("hello.txt")
	//fmt.Printf("Hello World!\n")
}

// > go mod init parser
