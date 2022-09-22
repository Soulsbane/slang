package main

import (
	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
)

func main() {
	arg.MustParse(&args)
	slang.LookupWord(args.WordToFind, args.ListAll)
}
