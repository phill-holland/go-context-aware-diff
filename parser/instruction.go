package parser

type Instruction struct {
	value    string
	children []Instruction
}

func (x Instruction) Print() string {

	var result string
	var prefix string

	if len(x.value) > 0 {
		if len(x.value) > 1 {
			if x.value[0] == '-' {
				prefix = string("- ")
			} else if x.value[0] == '+' {
				prefix = string("+ ")
			}
		}

		result += x.value + string("\n")

		if len(x.children) > 0 {
			result += prefix + string("{") + string("\n")
		}
	}

	for _, it := range x.children {
		result += it.Print()
	}

	if len(x.children) > 0 && len(x.value) > 0 {
		result += prefix + string("}") + string("\n")
	}

	return result
}

func (self Instruction) Compare(source Instruction) Instruction {
	var result Instruction
	var right map[string]int
	right = make(map[string]int)

	var idx int
	idx = 0

	for _, it := range self.children {
		temp := it.value

		_, ok := right[temp]
		if !ok {
			right[temp] = idx
		}
		idx = idx + 1
	}

	idx = 0
	for _, it := range source.children {
		temp := it.value
		val, ok := right[temp]
		if ok {
			if val == idx {
				t := self.children[idx].Compare(it)
				t.value = temp
				result.children = append(result.children, t)
			} else {
				var ttt Instruction
				ttt.value = temp + string("\n")
				result.children = append(result.children, ttt)
			}

			delete(right, temp)
		} else {
			result.children = append(result.children, it.Prefix(string("+ ")))
		}

		idx = idx + 1
	}

	for _, num := range right {
		var tt Instruction
		tt = self.children[num].Prefix(string("- "))

		result.children = append(result.children[:num+1], result.children[num:]...)
		result.children[num] = tt
	}

	return result
}

func (source Instruction) Prefix(value string) Instruction {
	var result Instruction

	result.value = value + source.value

	for _, it := range source.children {
		result.children = append(result.children, it.Prefix(value))
	}

	return result
}
