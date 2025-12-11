package slang

import (
	"encoding/json"
	"errors"
	"net/http"
)

const urbanDictionaryLink = "http://api.urbandictionary.com/v0/define?term="

var ErrFetchingResult = errors.New("failed to fetch definition")
var ErrNoResultsFound = errors.New("no results found")

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
	var results Results

	resp, err := http.Get(urbanDictionaryLink + wordToFind)

	if err != nil {
		return nil, ErrFetchingResult
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&results)

	if err != nil {
		return nil, ErrNoResultsFound
	}

	if len(results.List) > 0 {
		return results.List, nil
	} else {
		return []Result{}, ErrNoResultsFound
	}
}
