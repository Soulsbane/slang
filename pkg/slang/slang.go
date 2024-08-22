package slang

import (
	"errors"
	"github.com/imroc/req/v3"
)

const urbanDictionaryLink = "http://api.urbandictionary.com/v0/define"

var ErrorFetchingResult = errors.New("failed to fetch definition")
var NoResultsFound = errors.New("no results found")

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

func LookupWord(wordToFind string) ([]Result, error) {
	client := req.C()
	var results Results

	_, err := client.R().SetQueryParam("term", wordToFind).SetSuccessResult(&results).Get(urbanDictionaryLink)

	if err != nil {
		return []Result{}, ErrorFetchingResult
	} else {
		if len(results.List) > 0 {
			return results.List, nil
		} else {
			return []Result{}, NoResultsFound
		}
	}
}
