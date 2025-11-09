package words

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fatih/color"
	//"github.com/fatih/color"
)

const WORD_API = "https://random-word-api.herokuapp.com/word?length=5"

const (
	UNKNOWN int = iota
	GUESSED
	CHARACTER
	CHAR_AND_PLACEMENT
)

type Session struct {
	Active      bool
	CurrentWord map[byte][]int
	GuessNum    int
	Guess       map[byte]string
}

func (s *Session) MakeGuess(guess []byte) string {
	if !s.Active {
		return "session not active"
	}
	if len(guess) != 5 {
		return "incorret number of characters"
	}
	if s.GuessNum == 5 {
		s.EndSession()
		return "lost the game :("
	}
	var returnedGuess bytes.Buffer
	for key, value := range guess {
		if list, ok := s.CurrentWord[value]; ok {
			found := false
			for _, val := range list {
				if key == val {
					found = true
				}
			}
			if found {
				returnedGuess.WriteString(color.GreenString(string(value)))
			} else {
				returnedGuess.WriteString(color.YellowString(string(value)))
			}
		} else {
			returnedGuess.WriteString(color.BlackString(string(value)))
		}
	}
	return returnedGuess.String()
}

func StartSession(word []byte) (Session, error) {
	currentWord := map[byte][]int{}
	for cIdx, c := range word {
		word, ok := currentWord[c]
		if !ok {
			currentWord[c] = []int{cIdx}
		}
		currentWord[c] = append(word, cIdx)
	}
	return Session{
		Active:      true,
		CurrentWord: currentWord,
	}, nil
}

func (s *Session) EndSession() {
	s.Active = false
	s.Guess = map[byte]string{}
	s.GuessNum = 0
	s.CurrentWord = map[byte][]int{}
}

func ConvertWord(mapToConvert map[byte][]int) string {
	out := make([]byte, 5)
	for key, l := range mapToConvert {
		for _, value := range l {
			out[value] = key
		}
	}
	return string(out)
}

func GetWord() ([]byte, error) {
	resp, err := http.Get(WORD_API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	word := []string{}
	if err := decoder.Decode(&word); err != nil {
		return nil, err
	}
	if len(word) != 1 {
		return nil, errors.New("got an unexpected response from external api")
	}
	return []byte(word[0]), nil
}
