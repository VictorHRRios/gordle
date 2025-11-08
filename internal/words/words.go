package words

import (
	"encoding/json"
	"errors"
	"net/http"
)

const WORD_API = "https://random-word-api.herokuapp.com/word?length=5"

type Session struct {
	Active      bool
	CurrentWord []byte
	GuessNum    int
	Guess       []byte
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
	for charIdx, char := range s.CurrentWord {
		if guess[charIdx] != char {
			guess[charIdx] = '#'
		}
	}
	s.Guess = guess
	return string(s.Guess)
}

func StartSession() (Session, error) {
	word, err := GetWord()
	if err != nil {
		return Session{}, err
	}
	return Session{
		Active:      true,
		CurrentWord: []byte(word),
	}, nil
}

func (s *Session) EndSession() {
	s.Active = false
	s.Guess = nil
	s.GuessNum = 0
	s.CurrentWord = nil
}

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
