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

	var p rune
	var current string
	var quote_state int = 0
	var line_state int = 0

	for _, ch := range data {
		if ch == 13 {
		} else if (ch == ' ') && ((p == ' ') || (p == 10)) && (quote_state == 0) {
		} else if ch == 10 {
			if line_state == 0 {
				line_state += 1
			}
		} else if ch == '"' {
			if quote_state == 0 {
				quote_state += 1
			} else if quote_state == 1 {
				if p != '\\' {
					quote_state = 0
				}
			}
			current = current + string(ch)
		} else if ch == '{' {
			if (line_state == 1) && (len(current) > 0) {
				var child Instruction
				child.value = current

				current = ""

				stack[len(stack)-1].children = append(stack[len(stack)-1].children, child)
				stack = append(stack, &stack[len(stack)-1].children[len(stack[len(stack)-1].children)-1])

				line_state = 0
			}
		} else if ch == '}' {
			if len(current) > 0 {
				var child Instruction
				child.value = current

				stack[len(stack)-1].children = append(stack[len(stack)-1].children, child)

				current = ""
			}

			stack = stack[:len(stack)-1]

		} else if line_state == 1 {
			if len(current) > 0 {
				var child Instruction
				child.value = current

				stack[len(stack)-1].children = append(stack[len(stack)-1].children, child)

				current = ""
			}

			current = current + string(ch)
			line_state = 0
		} else {
			current = current + string(ch)
		}
		p = ch
	}

	return result
}
