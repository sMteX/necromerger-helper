package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui"
)

func main() {
	model := tui.New()
	program := tea.NewProgram(model)
	f, err := tea.LogToFile("debug.log", "debug")
	defer f.Close()
	if err != nil {
		fmt.Println("fatal error while opening debug log file:", err)
		os.Exit(1)
	}
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running TUI: %v\n", err)
		os.Exit(1)
	}
}
