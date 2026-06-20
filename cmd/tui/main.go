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
	f, err := tea.LogToFile("debug.log", "debug")
	defer f.Close()
	if err != nil {
		fmt.Println("fatal error while opening debug log file:", err)
		os.Exit(1)
	}
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running TUI: %v\n", err)
		os.Exit(1)
	}
}
