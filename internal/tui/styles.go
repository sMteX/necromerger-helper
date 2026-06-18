package tui

import (
	"charm.land/lipgloss/v2"
)

var (
	// Colors
	colorHighlight = lipgloss.Color("#00a7f8") // Pinkish/Purple for active
	colorDim       = lipgloss.Color("#626262") // Gray for inactive

	// Styles
	styleActiveOption = lipgloss.NewStyle().
				Foreground(colorHighlight)

	styleInactiveOption = lipgloss.NewStyle().
				Foreground(colorDim)
)
