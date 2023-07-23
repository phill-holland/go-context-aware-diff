package parser

import (
	"fmt"
	"io/ioutil"
)

func Load(filename string) Instruction {
	var result Instruction
	fileContent, err := ioutil.ReadFile(filename)
	if err == nil {
		data := string(fileContent)
		fmt.Println(data)
		result = Parse(data)
	} else {
		fmt.Println("File reading error", err)
	}

	return result
}

func Parse(data string) Instruction {
	var result Instruction
	var stack []*Instruction

	stack = append(stack, &result)

	return result
}
