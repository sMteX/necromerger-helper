package shared

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type colors struct {
	Good, Bad color.Color

	RuneIce, RunePoison, RuneBlood, RuneMoon color.Color

	Border color.Color
}
type styles struct {
	MainContainer, Header lipgloss.Style
}

var (
	Colors = colors{
		Good:       lipgloss.Color("#4ADE80"),
		Bad:        lipgloss.Color("#F87171"),
		RuneIce:    lipgloss.Color("#7DD3FC"),
		RunePoison: lipgloss.Color("#86EFAC"),
		RuneBlood:  lipgloss.Color("#FCA5A5"),
		RuneMoon:   lipgloss.Color("#FDE68A"),
		Border:     lipgloss.Color("#7D56F4"),
	}
	Styles = styles{
		MainContainer: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Colors.Border).
			Padding(1, 4),
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(Colors.Border).
			MarginBottom(1),
	}
)
