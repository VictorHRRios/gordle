package handlers

import (
	"bytes"
	"errors"
	"os"

	"github.com/VictorHRRios/gordle/internal/words"
	"github.com/fatih/color"
)

type Handler struct {
	handlerFns map[string]Command
	session    words.Session
}

type Command struct {
	Action        func(...string) (string, error)
	Documentation string
}

func Init() *Handler {
	h := &Handler{}
	h.handlerFns = make(map[string]Command)
	h.RegisterCommand("exit", "exit the application", exitCommand)
	h.RegisterCommand("start", "start the game", h.startCommand)
	h.RegisterCommand("end", "end the game", h.endCommand)
	h.RegisterCommand("help", "get help", h.helpCommand)
	return h
}

func (h *Handler) Guess(name string, params ...string) (string, error) {
	return h.session.MakeGuess([]byte(name)), nil
}

func (h *Handler) Exec(name string, params ...string) (string, error) {
	exec, ok := h.handlerFns[name]
	if ok {
		return exec.Action()
	}
	if h.session.Active && len(name) == 5 {
		return h.Guess(name)
	}
	return "", errors.New("command not found")
}

func (h *Handler) Status() bool {
	return h.session.Active
}

func NewCommand(doc string, fn func(...string) (string, error)) Command {
	return Command{
		Action:        fn,
		Documentation: doc,
	}
}

func (h *Handler) RegisterCommand(name string, doc string, fn func(...string) (string, error)) {
	h.handlerFns[name] = NewCommand(doc, fn)
}

func (h *Handler) startCommand(params ...string) (string, error) {
	var err error
	h.session, err = words.StartSession()
	if err != nil {
		return "", err
	}
	return "game started...", nil
}

func (h *Handler) endCommand(params ...string) (string, error) {
	h.session.EndSession()
	return "ok", nil
}

func exitCommand(params ...string) (string, error) {
	println("exiting program...")
	os.Exit(0)
	return "ok", nil
}

func (h *Handler) helpCommand(params ...string) (string, error) {
	var out bytes.Buffer
	out.WriteString("Supported Commands\n")
	for name, param := range h.handlerFns {
		out.WriteString(name)
		out.WriteString(" : ")
		out.WriteString(param.Documentation + "\n")
	}

	return color.BlueString(out.String()), nil
}
