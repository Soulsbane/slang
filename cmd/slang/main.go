package main

import (
	"fmt"
	"os"

	"github.com/Soulsbane/slang/pkg/slang"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	wordToFind := args[0]
	definition, err := slang.LookupWord(wordToFind)

	if err != nil {
		panic(err)
	}

	fmt.Println(definition)
}
