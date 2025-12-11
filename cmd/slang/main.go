package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/slang/pkg/slang"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tiagomelo/go-clipboard/clipboard"
)

func getOutputTable(maxWidth int) table.Writer {

	outputTable := table.NewWriter()
	outputTable.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft, WidthMax: maxWidth},
	})

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

func handleListResults(results []slang.Result, listAll bool, copyToClipboard bool, maxWidth int) {
	outputTable := getOutputTable(maxWidth)

	if listAll {
		for _, result := range results {
			outputTable.AppendRow(table.Row{result})
		}
	} else {
		outputTable.AppendRow(table.Row{results[0].Definition})

		if copyToClipboard {
			handleCopyToClipboard(results[0].Definition)
		}
	}

	outputTable.Render()
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	results, err := slang.LookupWord(args.Word)

	if errors.Is(err, slang.ErrNoResultsFound) {
		fmt.Println("No results found")
	} else if errors.Is(err, slang.ErrFetchingResult) {
		fmt.Println("Failed to fetch definition")
	} else {
		handleListResults(results, args.ListAll, args.Copy, args.MaxWidth)
	}
}
