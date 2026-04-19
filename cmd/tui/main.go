package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui"
)

func main() {
	m := tui.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running TUI: %v\n", err)
		os.Exit(1)
	}
}
