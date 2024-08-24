package main

import (
	"errors"
	"fmt"
	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tiagomelo/go-clipboard/clipboard"
	"os"
)

func getOutputTable() table.Writer {

	outputTable := table.NewWriter()

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{"Definition"})
	outputTable.SetStyle(table.StyleRounded)
	outputTable.Style().Options.SeparateRows = true

	return outputTable
}

func handleCopyToClipboard(definition string) {
	c := clipboard.New()

	if err := c.CopyText(definition); err != nil {
		fmt.Println(err)
	}
}

func handleListResults(results []slang.Result, listAll bool, copyToClipboard bool) {
	//outputTable := getOutputTable()

	if listAll {
		for _, result := range results {
			//outputTable.AppendRow(table.Row{result.Definition})
			fmt.Println(result.Definition + "\n")
		}
	} else {
		//outputTable.AppendRow(table.Row{results[0].Definition})
		fmt.Println(results[0].Definition)

		if copyToClipboard {
			handleCopyToClipboard(results[0].Definition)
		}
	}

	//outputTable.Render()
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
