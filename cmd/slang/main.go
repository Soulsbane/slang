package main

import (
	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
)

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	slang.LookupWord(args.Word, args.ListAll)
}
