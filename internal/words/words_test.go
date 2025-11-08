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
