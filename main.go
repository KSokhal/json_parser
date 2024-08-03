package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func dataFromFile(filePath string) (data string) {
	content, err := os.ReadFile(filePath) // Read the entire file content
	check(err)

	// Convert the content to a string
	data = string(content)
	return
}

func main() {
	data := dataFromFile("tests/test.json")

	tokens := tokenizer(data)

	fmt.Println(tokens)
	fmt.Println(len(tokens))
	node := parser(tokens)

	fmt.Println(node)
}
