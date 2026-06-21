package runes

import (
	"image/color"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

var (
	runeColorMap = map[models.RuneType]color.Color{
		models.RuneIce:    shared.Colors.RuneIce,
		models.RunePoison: shared.Colors.RunePoison,
		models.RuneBlood:  shared.Colors.RuneBlood,
		models.RuneMoon:   shared.Colors.RuneMoon,
		models.RuneDeath:  shared.Colors.RuneDeath,
		models.RuneCosmic: shared.Colors.RuneCosmic,
	}
	runeColumn  = lipgloss.NewStyle().Width(10)
	valueColumn = lipgloss.NewStyle().Width(9)
)

func (m *Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Render(
			runeColumn.Render("Rune")+valueColumn.Render("Have")+valueColumn.Render("Total")+valueColumn.Render("Need"),
		),
		strings.Repeat("─", runeColumn.GetWidth()+3*valueColumn.GetWidth()),
		m.renderRuneRow(fieldIce),
		m.renderRuneRow(fieldPoison),
		m.renderRuneRow(fieldBlood),
		m.renderRuneRow(fieldMoon),
		m.renderRuneRow(fieldDeath),
		m.renderRuneRow(fieldCosmic),
	)
}

func (m *Model) renderRuneRow(i fieldIndex) string {
	r := runeByFieldType[i]

	needColumn := func() lipgloss.Style {
		if m.PossessedRunes[r] >= m.RuneTotal[r] {
			return valueColumn.Foreground(shared.Colors.Good)
		}
		return valueColumn.Foreground(shared.Colors.Bad)
	}()

	needRunes := m.RuneNeeded[r]

	valueText := ""
	if m.cursor == int(i) {
		valueText = valueColumn.Render(m.CurrentInput().View())
	} else {
		valueText = valueColumn.Render(shared.FormatNumberLong(m.PossessedRunes[r]))
	}
	return runeColumn.Foreground(runeColorMap[r]).Render(string(r)) +
		valueText +
		valueColumn.Render(shared.FormatNumberLong(m.RuneTotal[r])) +
		needColumn.Render(shared.FormatNumberLong(needRunes))
}

func (m *Model) GetHelpItems() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
