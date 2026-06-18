package prestigePlan

import (
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m Model) renderBaseTab() string {
	labelStyle := lipgloss.NewStyle().Width(30)
	valueStyle := lipgloss.NewStyle().Width(10).AlignHorizontal(lipgloss.Right)

	selectedInput := 0
	// TODO: actual inputs
	inputs := [][]string{
		{"Devourer level:", "300"},
		// TODO: consider arrow input
		{"Max Feat Tier:", "29"},
		{"'Others' multiplier [%]:", "210"},
		// TODO: consider arrow input
		{"Legendary group bonus count:", "1"},
		{"Leftover shards:", "123456"},
	}
	var lines []string
	for i, input := range inputs {
		if i == selectedInput {
			lines = append(lines, labelStyle.Foreground(shared.Colors.Good).Render(input[0])+valueStyle.Foreground(shared.Colors.Good).Render(input[1]))
		} else {
			lines = append(lines, labelStyle.Render(input[0])+valueStyle.Render(input[1]))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		lines...,
	)
}

func (m Model) getBaseTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
