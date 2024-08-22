package slang

import (
	"fmt"
	"github.com/imroc/req/v3"
)

const urbanDictionaryLink = "http://api.urbandictionary.com/v0/define"

type Results struct {
	Type   string `json:"result_type"`
	Tags   []string
	List   []Result
	Sounds []string
}

type Result struct {
	Author     string
	Word       string
	Definition string
	Example    string
	Permalink  string
	Upvote     int `json:"thumbs_up"`
	Downvote   int `json:"thumbs_down"`
}

func fetchWord(wordToFind string) (Results, error) {
	client := req.C()
	var results Results

	_, err := client.R().SetQueryParam("term", wordToFind).SetSuccessResult(&results).Get(urbanDictionaryLink)
	return results, err
}

func LookupWord(wordToFind string, listAll bool) {
	results, err := fetchWord(wordToFind)

	if err != nil {
		fmt.Println("Error fetching word! Try again later.")
	} else {
		if len(results.List) > 0 {
			if listAll {
				for _, result := range results.List {
					fmt.Println(result.Definition + "\n")
				}
			} else {
				fmt.Println(results.List[0].Definition)
			}
		} else {
			fmt.Println("No results found for: " + wordToFind)
		}
	}
}
