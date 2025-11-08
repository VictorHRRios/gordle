package repl

import (
	"bufio"
	"io"
	"strings"

	"github.com/VictorHRRios/gordle/internal/handlers"
	"github.com/fatih/color"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	handler := handlers.Init()
	for {
		io.WriteString(out, "> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		commands := strings.Split(line, " ")

		command := commands[0]
		params := commands[1:]
		output, err := handler.Exec(command, params...)
		if err != nil {
			io.WriteString(out, color.RedString(err.Error()))
			io.WriteString(out, "\n")
		}

		io.WriteString(out, output)
	}
}
