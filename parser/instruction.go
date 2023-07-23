package parser

import "fmt"

type Instruction struct {
	value    string
	children []Instruction
}

func (x Instruction) Print() {
	fmt.Println(x.value)
}
