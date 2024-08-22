package main

import (
	"errors"
	"fmt"
	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
)

func handleListResults(results []slang.Result, listAll bool) {
	if listAll {
		for _, result := range results {
			fmt.Println(result.Definition + "\n")
		}
	} else {
		fmt.Println(results[0].Definition)
	}
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	results, err := slang.LookupWord(args.Word)

	if errors.As(err, &slang.NoResultsFound) {
		fmt.Println("No results found")
	} else if errors.As(err, &slang.ErrorFetchingResult) {
		fmt.Println("Failed to fetch definition")
	} else {
		handleListResults(results, args.ListAll)
	}
}
