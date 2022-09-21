package slang

import (
	"github.com/imroc/req/v3"
)

const API_URL = "http://api.urbandictionary.com/v0/define"

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

	_, err := client.R().SetQueryParam("term", wordToFind).SetResult(&results).Get(API_URL)
	return results, err
}

func LookupWord(wordToFind string, listAll bool) (string, error) {
	results, err := fetchWord(wordToFind)
	return results.List[0].Definition, err
}
