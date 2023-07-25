package main

import (
	"fmt"
	"os"
	"parser"
)

func main() {

	if len(os.Args[1:]) >= 2 {

		p1 := parser.Load(os.Args[1])
		p2 := parser.Load(os.Args[2])

		result := p1.Compare(p2)

		fmt.Printf(result.Print())
	}
}
