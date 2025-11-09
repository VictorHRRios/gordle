package words

import "testing"

func Test_getWord(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{"successful response"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := GetWord()
			if gotErr != nil {
				t.Errorf("getWord() failed: %v", gotErr)
			}
		})
	}
}

func TestStartSession(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word    string
		want    Session
		wantErr bool
	}{
		{"basic", "clonk", Session{Active: true, CurrentWord: map[rune][]int{
			'c': {0},
			'l': {1},
			'o': {2},
			'n': {3},
			'k': {4},
		},
		}, false},
		{"repeat", "manas", Session{Active: true, CurrentWord: map[rune][]int{
			'm': {0},
			'a': {1, 3},
			'n': {2},
			's': {4},
		},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := StartSession(tt.word)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("StartSession() failed: %v", gotErr)
				}
				return
			}
			wantWord := ConvertWord(tt.want.CurrentWord)
			gotWord := ConvertWord(got.CurrentWord)

			if gotWord != wantWord {
				t.Errorf("StartSession() = %v, want %v", gotWord, wantWord)
			}
		})
	}
}
