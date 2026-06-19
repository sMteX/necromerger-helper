package prestigePlan

import (
	"image/color"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *Model) renderRuneTab() string {
	runeColumn := lipgloss.NewStyle().Width(10)
	valueColumn := lipgloss.NewStyle().Width(9).AlignHorizontal(lipgloss.Right)

	lines := []string{
		lipgloss.NewStyle().Bold(true).Render(
			runeColumn.Render("Rune") + valueColumn.Render("Have") + valueColumn.Render("Total") + valueColumn.Render("Need"),
		),
		strings.Repeat("─", runeColumn.GetWidth()+3*valueColumn.GetWidth()),
		m.renderRuneTableRow(models.RuneIce),
		m.renderRuneTableRow(models.RunePoison),
		m.renderRuneTableRow(models.RuneBlood),
		m.renderRuneTableRow(models.RuneMoon),
		m.renderRuneTableRow(models.RuneDeath),
		m.renderRuneTableRow(models.RuneCosmic),
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		lines...,
	)
}

var runeColorMap = map[models.RuneType]color.Color{
	models.RuneIce:    shared.Colors.RuneIce,
	models.RunePoison: shared.Colors.RunePoison,
	models.RuneBlood:  shared.Colors.RuneBlood,
	models.RuneMoon:   shared.Colors.RuneMoon,
	models.RuneDeath:  shared.Colors.RuneDeath,
	models.RuneCosmic: shared.Colors.RuneCosmic,
}

func (m *Model) renderRuneTableRow(rune models.RuneType) string {
	runeColumn := lipgloss.NewStyle().Width(10)
	valueColumn := lipgloss.NewStyle().Width(9).AlignHorizontal(lipgloss.Right)

	needColumn := func() lipgloss.Style {
		if m.currentRunes[rune] >= m.totalRunesNeeded[rune] {
			return valueColumn.Foreground(shared.Colors.Good)
		}
		return valueColumn.Foreground(shared.Colors.Bad)
	}()

	needRunes := int(math.Max(float64(m.totalRunesNeeded[rune]-m.currentRunes[rune]), 0))

	return runeColumn.Foreground(runeColorMap[rune]).Render(string(rune)) +
		valueColumn.Render(shared.FormatNumberLong(m.currentRunes[rune])) +
		valueColumn.Render(shared.FormatNumberLong(m.totalRunesNeeded[rune])) +
		needColumn.Render(shared.FormatNumberLong(needRunes))
}

func (m *Model) getRuneTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
