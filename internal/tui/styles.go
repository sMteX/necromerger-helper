package tui

import "charm.land/lipgloss/v2"

var (
	// Colors
	colorBorder    = lipgloss.Color("#7D56F4") // Purplish/Blue
	colorHighlight = lipgloss.Color("#00a7f8") // Pinkish/Purple for active
	colorDim       = lipgloss.Color("#626262") // Gray for inactive

	// Styles
	styleMenuHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorBorder).
			MarginBottom(1)

	styleMainContainer = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorBorder).
				Padding(1, 4)

	styleActiveOption = lipgloss.NewStyle().
				Foreground(colorHighlight)

	styleInactiveOption = lipgloss.NewStyle().
				Foreground(colorDim)
)
