package prestigePlan

import (
	"slices"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *Model) addBaseTabFields() {
	m.fields[fieldBaseDevourerLevel] = inputField{
		label: "Devourer Level",
		options: []string{"35", "40", "45", "50", "55", "60", "65", "70",
			"75", "80", "85", "90", "95", "100", "150",
			"200", "300", "400", "500", "600", "700",
			"800", "900", "1000"},
		width:          8,
		characterLimit: 4,
		initialValue:   strconv.Itoa(m.plan.DevourerLevel),
		validate:       inputValidationIntInRange(1, 1000),
	}
	m.fields[fieldBaseFeatTiers] = inputField{
		label:          "Max Feat Tier",
		step:           1,
		width:          7,
		characterLimit: 2,
		initialValue:   strconv.Itoa(m.plan.FeatTiers),
		validate:       inputValidationIntInRange(1, 35),
	}
	m.fields[fieldBaseOtherMultiplier] = inputField{
		label:          "'Others' Multiplier [%]",
		step:           0,
		width:          8,
		characterLimit: 3,
		initialValue:   strconv.Itoa(int(m.plan.OtherMultiplier * 100)),
		validate:       inputValidationIntInRange(100, 1000),
	}
	m.fields[fieldBaseGroupBonusCount] = inputField{
		label:          "Leg. Group Bonus Count",
		step:           1,
		width:          5,
		characterLimit: 1,
		initialValue:   strconv.Itoa(m.plan.GroupBonusCount),
		validate:       inputValidationIntInRange(1, 3),
	}
	m.fields[fieldBaseLeftoverShards] = inputField{
		label:          "Leftover Shards",
		step:           0,
		width:          11,
		characterLimit: 7,
		initialValue:   strconv.Itoa(m.plan.LeftoverShards),
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
		if field.step > 0 || len(field.options) > 0 {
			return labelStyle.Foreground(shared.Colors.Good).Render(field.label) + valueStyle.Foreground(shared.Colors.Good).AlignHorizontal(lipgloss.Left).Render(shared.PadRight("< "+field.input.Value()+" >", valueStyle.GetWidth()))
		}
		return labelStyle.Foreground(shared.Colors.Good).Render(field.label) + "  " + field.input.View()
	}
	return labelStyle.Render(field.label) + valueStyle.Render(shared.PadRight("  "+field.input.Value(), valueStyle.GetWidth()))
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
		field := m.fields[fieldIndex(m.cursor)]
		if len(field.options) > 0 {
			cur := m.currentInput().Value()
			idx := slices.Index(field.options, cur)
			if idx < 0 {
				idx = 0
			}
			if msg.String() == "left" && idx > 0 {
				idx--
			} else if msg.String() == "right" && idx < len(field.options)-1 {
				idx++
			} else {
				return m, nil
			}
			newVal := field.options[idx]
			m.currentInput().SetValue(newVal)
			m.parseBaseTabFieldValues(fieldIndex(m.cursor), newVal)
			return m, nil
		} else if field.step > 0 {
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
					return m, nil
				}
			}
			m.currentInput().SetValue(newVal)
			m.parseBaseTabFieldValues(fieldIndex(m.cursor), newVal)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.fields[m.cursor].input, cmd = m.currentInput().Update(msg)
	if m.currentInput().Err == nil {
		m.parseBaseTabFieldValues(fieldIndex(m.cursor), m.currentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseBaseTabFieldValues(i fieldIndex, value string) {
	switch i {
	case fieldBaseDevourerLevel:
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.DevourerLevel = v
		}
	case fieldBaseFeatTiers:
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.FeatTiers = v
		}
	case fieldBaseOtherMultiplier:
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.OtherMultiplier = float64(v) / 100.0
		}
	case fieldBaseGroupBonusCount:
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.GroupBonusCount = v
		}
	case fieldBaseLeftoverShards:
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.LeftoverShards = v
		}
	}
	m.recalculate()
}
