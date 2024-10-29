package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := new(InMemStore)
	m := NewModel(store)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
