package prestigePlan

import (
	"fmt"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *Model) addBaseTabFields() {
	m.fields[fieldBaseDevourerLevel] = inputField{
		label:          "Devourer Level",
		step:           1,
		width:          8,
		characterLimit: 3,
		initialValue:   strconv.Itoa(m.baseInputs.devourerLevel),
		validate:       inputValidationIntInRange(1, 1000),
	}
	m.fields[fieldBaseFeatTiers] = inputField{
		label:          "Max Feat Tier",
		step:           1,
		width:          7,
		characterLimit: 2,
		initialValue:   strconv.Itoa(m.baseInputs.featTiers),
		validate:       inputValidationIntInRange(1, 35),
	}
	// TODO: an int field instead?
	m.fields[fieldBaseOtherMultiplier] = inputField{
		label:          "'Others' Multiplier",
		step:           0,
		width:          8,
		characterLimit: 4,
		initialValue:   fmt.Sprintf("%.2f", m.baseInputs.otherMultiplier),
		validate: func(s string) error {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return fmt.Errorf("must be a number")
			}
			if v < 1.0 {
				return fmt.Errorf("must be ≥ 1.0")
			}
			return nil
		},
	}
	m.fields[fieldBaseGroupBonusCount] = inputField{
		label:          "Leg. Group Bonus Count",
		step:           1,
		width:          5,
		characterLimit: 1,
		initialValue:   strconv.Itoa(m.baseInputs.groupBonusCount),
		validate:       inputValidationIntInRange(1, 3),
	}
	m.fields[fieldBaseLeftoverShards] = inputField{
		label:          "Leftover Shards",
		step:           0,
		width:          11,
		characterLimit: 7,
		initialValue:   strconv.Itoa(m.baseInputs.leftoverShards),
		validate:       inputValidationIntInRange(0, 10000000),
	}
}

func (m *Model) renderBaseTab() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		m.renderBaseTabInput(fieldBaseDevourerLevel),
		m.renderBaseTabInput(fieldBaseFeatTiers),
		m.renderBaseTabInput(fieldBaseOtherMultiplier),
		m.renderBaseTabInput(fieldBaseGroupBonusCount),
		m.renderBaseTabInput(fieldBaseLeftoverShards),
	)
}

func (m *Model) renderBaseTabInput(i fieldIndex) string {
	labelStyle := lipgloss.NewStyle().Width(30)
	valueStyle := lipgloss.NewStyle().Width(10)

	field := m.fields[i]
	if m.cursor == int(i) {
		if field.step > 0 {
			return labelStyle.Foreground(shared.Colors.Good).Render(field.label) + valueStyle.Foreground(shared.Colors.Good).AlignHorizontal(lipgloss.Center).Render("← "+field.input.Value()+" →")
		}
		return labelStyle.Foreground(shared.Colors.Good).Render(field.label) + field.input.View()
	}
	return labelStyle.Render(field.label) + valueStyle.Render(field.input.Value())
}

func (m *Model) getBaseTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("← / →  change value"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}

func (m *Model) handleBaseTabKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > int(fieldBaseDevourerLevel) {
			m.currentInput().Blur()
			m.cursor--
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldBaseLeftoverShards) {
			m.currentInput().Blur()
			m.cursor++
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "left", "right":
		// For arrow-adjustable fields, ←/→ increment/decrement the value directly.
		// For text-only fields (step == 0), fall through so the textinput handles
		// cursor movement within the text.
		field := m.fields[fieldIndex(m.cursor)]
		if field.step > 0 {
			cur, err := strconv.Atoi(m.currentInput().Value())
			if err != nil {
				return m, nil
			}
			if msg.String() == "left" {
				cur -= field.step
			} else {
				cur += field.step
			}
			newVal := strconv.Itoa(cur)
			if field.validate != nil {
				if err := field.validate(newVal); err != nil {
					// didn't pass validate, don't change anything
					return m, nil
				}
			}
			m.fields[m.cursor].input.SetValue(newVal)
			m.parseBaseTabFieldValues(fieldIndex(m.cursor), newVal)
			// TODO: recalculate m.calculatedOutputs from m.baseInputs
			return m, nil
		}
	}

	// Everything else — character input, backspace, and ←/→ cursor movement for
	// text-only fields — goes to the focused textinput.
	var cmd tea.Cmd
	m.fields[m.cursor].input, cmd = m.currentInput().Update(msg)
	if m.currentInput().Err == nil {
		m.parseBaseTabFieldValues(fieldIndex(m.cursor), m.currentInput().Value())
		// TODO: recalculate m.calculatedOutputs from m.baseInputs
	}
	return m, cmd
}

func (m *Model) parseBaseTabFieldValues(i fieldIndex, value string) {
	switch i {
	case fieldBaseDevourerLevel:
		if v, err := strconv.Atoi(value); err == nil {
			m.baseInputs.devourerLevel = v
		}
	case fieldBaseFeatTiers:
		if v, err := strconv.Atoi(value); err == nil {
			m.baseInputs.featTiers = v
		}
	case fieldBaseOtherMultiplier:
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			m.baseInputs.otherMultiplier = v
		}
	case fieldBaseGroupBonusCount:
		if v, err := strconv.Atoi(value); err == nil {
			m.baseInputs.groupBonusCount = v
		}
	case fieldBaseLeftoverShards:
		if v, err := strconv.Atoi(value); err == nil {
			m.baseInputs.leftoverShards = v
		}
	}
}
