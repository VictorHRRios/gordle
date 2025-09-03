package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	word     [5]rune
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		word:     [5]rune{' ', ' ', ' ', ' ', ' '},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		char := msg.String()
		// Cool, what was the actual key pressed?
		switch {

		case len(char) == 1:
			r := rune(char[0])
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				m.word[m.cursor] = r

				if m.cursor < len(m.word)-1 {
					m.cursor++
				}
			}
		// These keys should exit the program.
		case char == "ctrl+c":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case char == "left":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case char == "right":
			if m.cursor < len(m.word)-1 {
				m.cursor++
			}

		case char == "backspace":
			m.word[m.cursor] = ' '
			if m.cursor > 0 {
				m.cursor--
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case char == "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Welcome to Gordle!\n\n"

	// Iterate over our choices
	for _, choice := range m.word {
		// Render the row
		s += fmt.Sprintf("%c", choice)
	}
	s += "\n"

	for i := range m.word {
		if m.cursor == i {
			s += "^"
		} else {
			s += "-"
		}
	}

	// The footer
	s += "\nPress ctrl+c to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
