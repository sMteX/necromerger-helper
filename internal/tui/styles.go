package tui

import (
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

var (
	// Colors
	colorHighlight = lipgloss.Color("#00a7f8") // Pinkish/Purple for active

	// Styles
	styleActiveOption = lipgloss.NewStyle().
				Foreground(colorHighlight)

	styleInactiveOption = lipgloss.NewStyle().
				Foreground(shared.Colors.Dim)
)
