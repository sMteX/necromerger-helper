package prestigePlan

import (
	"image/color"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

var runeByFieldType = map[fieldIndex]models.RuneType{
	fieldRunesIce:    models.RuneIce,
	fieldRunesPoison: models.RunePoison,
	fieldRunesBlood:  models.RuneBlood,
	fieldRunesMoon:   models.RuneMoon,
	fieldRunesDeath:  models.RuneDeath,
	fieldRunesCosmic: models.RuneCosmic,
}

func (m *Model) addRunesTabFields() {
	for i := fieldRunesIce; i <= fieldRunesCosmic; i++ {
		r := runeByFieldType[i]
		m.fields[i] = inputField{
			label:          string(r),
			step:           0,
			width:          6,
			characterLimit: 6,
			validate:       inputValidationIntInRange(0, 10000000),
			initialValue:   strconv.Itoa(m.plan.PossessedRunes[r]),
		}
	}
}

func (m *Model) renderRuneTab() string {
	runeColumn := lipgloss.NewStyle().Width(10)
	valueColumn := lipgloss.NewStyle().Width(9)

	lines := []string{
		lipgloss.NewStyle().Bold(true).Render(
			runeColumn.Render("Rune") + valueColumn.Render("Have") + valueColumn.Render("Total") + valueColumn.Render("Need"),
		),
		strings.Repeat("─", runeColumn.GetWidth()+3*valueColumn.GetWidth()),
		m.renderRuneTableRow(fieldRunesIce),
		m.renderRuneTableRow(fieldRunesPoison),
		m.renderRuneTableRow(fieldRunesBlood),
		m.renderRuneTableRow(fieldRunesMoon),
		m.renderRuneTableRow(fieldRunesDeath),
		m.renderRuneTableRow(fieldRunesCosmic),
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

func (m *Model) renderRuneTableRow(i fieldIndex) string {
	runeColumn := lipgloss.NewStyle().Width(10)
	valueColumn := lipgloss.NewStyle().Width(9)
	r := runeByFieldType[i]

	needColumn := func() lipgloss.Style {
		if m.plan.PossessedRunes[r] >= m.result.RuneTotal[r] {
			return valueColumn.Foreground(shared.Colors.Good)
		}
		return valueColumn.Foreground(shared.Colors.Bad)
	}()

	needRunes := m.result.RuneNeeded[r]

	valueText := ""
	if m.cursor == int(i) {
		valueText = valueColumn.Render(m.currentInput().View())
	} else {
		valueText = valueColumn.Render(shared.FormatNumberLong(m.plan.PossessedRunes[r]))
	}
	return runeColumn.Foreground(runeColorMap[r]).Render(string(r)) +
		valueText +
		valueColumn.Render(shared.FormatNumberLong(m.result.RuneTotal[r])) +
		needColumn.Render(shared.FormatNumberLong(needRunes))
}

func (m *Model) getRuneTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}

func (m *Model) handleRunesTabKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > int(fieldRunesIce) {
			m.currentInput().Blur()
			m.cursor--
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldRunesCosmic) {
			m.currentInput().Blur()
			m.cursor++
			return m, m.currentInput().Focus()
		}
		return m, nil
	}

	// Everything else — character input, backspace, and ←/→ cursor movement for
	// text-only fields — goes to the focused textinput.
	var cmd tea.Cmd
	m.fields[m.cursor].input, cmd = m.currentInput().Update(msg)
	if m.currentInput().Err == nil {
		m.parseRunesTabFieldValues(fieldIndex(m.cursor), m.currentInput().Value())
		// TODO: recalculate m.calculatedOutputs from m.baseInputs
	}
	return m, cmd
}

func (m *Model) parseRunesTabFieldValues(i fieldIndex, value string) {
	if v, err := strconv.Atoi(value); err == nil {
		m.plan.PossessedRunes[runeByFieldType[i]] = v
	}
	m.recalculate()
}
