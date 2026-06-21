package runes

import (
	"image/color"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necromerger-helper/internal/models"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
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
		lipgloss.JoinHorizontal(lipgloss.Top,
			runeColumn.Bold(true).Render("Rune"),
			valueColumn.Bold(true).Render("Have"),
			valueColumn.Bold(true).Render("Total"),
			lipgloss.NewStyle().Bold(true).Foreground(shared.Colors.Bad).Render("Need"),
			"/",
			lipgloss.NewStyle().Bold(true).Foreground(shared.Colors.Good).Render("Excess"),
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

	needRunes := m.result.RuneTotal[r] - m.PossessedRunes[r]

	needColumn := func() lipgloss.Style {
		if needRunes <= 0 {
			return valueColumn.Foreground(shared.Colors.Good)
		}
		return valueColumn.Foreground(shared.Colors.Bad)
	}()

	valueText := ""
	if m.Cursor == int(i) {
		valueText = valueColumn.Render(m.CurrentInput().View())
	} else {
		valueText = valueColumn.Render(shared.FormatNumberLong(m.PossessedRunes[r]))
	}
	return runeColumn.Foreground(runeColorMap[r]).Render(string(r)) +
		valueText +
		valueColumn.Render(shared.FormatNumberLong(m.result.RuneTotal[r])) +
		needColumn.Render(shared.FormatNumberLong(int(math.Abs(float64(needRunes)))))
}

func (m *Model) GetHelpItems() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
	}
}
