package slang

func fetchWord(wordToFind string) (string, error) {
	return "Hello World", nil
}

func LookupWord(wordToFind string) (string, error) {
	return fetchWord(wordToFind)
}
