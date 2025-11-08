package words

import (
	"encoding/json"
	"errors"
	"net/http"
)

const WORD_API = "https://random-word-api.herokuapp.com/word?length=5"

func GetWord() (string, error) {
	resp, err := http.Get(WORD_API)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	word := []string{}
	if err := decoder.Decode(&word); err != nil {
		return "", err
	}
	if len(word) != 1 {
		return "", errors.New("got an unexpected response from external api")
	}
	return word[0], nil
}
