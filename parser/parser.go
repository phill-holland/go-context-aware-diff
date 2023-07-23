package parser

import (
	"fmt"
	"io/ioutil"
)

func load(filename string) int32 {
	fileContent, err := ioutil.ReadFile(filename)
	if err == nil {
		data := string(fileContent)
		fmt.Println(data)
	}
	return 1
}
