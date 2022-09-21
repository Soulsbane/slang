package main

import (
	"fmt"

	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
)

func main() {
	arg.MustParse(&args)

	definition, err := slang.LookupWord(args.WordToFind, args.ListAll)

	if err != nil {
		panic(err)
	}

	fmt.Println(definition)
}
