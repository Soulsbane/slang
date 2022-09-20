package slang

import (
	"github.com/imroc/req/v3"
)

const API_URL = "http://api.urbandictionary.com/v0/define"

type SearchResult struct {
	Type    string `json:"result_type"`
	Tags    []string
	Results []Result `json:"list"`
	Sounds  []string
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

func fetchWord(wordToFind string) (string, error) {
	client := req.C()
	var result SearchResult

	_, err := client.R().SetQueryParam("term", wordToFind).SetResult(&result).Get(API_URL)

	if err != nil {
	}

	return result.Results[0].Definition, err
}

func LookupWord(wordToFind string) (string, error) {
	return fetchWord(wordToFind)
}
