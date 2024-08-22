package main

import (
	"errors"
	"fmt"
	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
	"github.com/tiagomelo/go-clipboard/clipboard"
)

func handleCopyToClipboard(definition string) {
	c := clipboard.New()

	if err := c.CopyText(definition); err != nil {
		fmt.Println(err)
	}
}

func handleListResults(results []slang.Result, listAll bool, copyToClipboard bool) {
	if listAll {
		for _, result := range results {
			fmt.Println(result.Definition + "\n")
		}
	} else {
		fmt.Println(results[0].Definition)

		if copyToClipboard {
			handleCopyToClipboard(results[0].Definition)
		}
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
		handleListResults(results, args.ListAll, args.Copy)
	}
}
