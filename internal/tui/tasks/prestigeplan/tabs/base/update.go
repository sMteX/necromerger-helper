package base

import (
	"slices"
	"strconv"

	"charm.land/bubbletea/v2"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	if keyPressMsg, ok := msg.(tea.KeyPressMsg); ok {
		return m.handleKeyPress(keyPressMsg)
	}
	// Non-key messages (cursor blink ticks) must reach the active textinput.
	var cmd tea.Cmd
	m.currentField().Input, cmd = m.CurrentInput().Update(msg)
	return m, cmd
}

func (m *Model) handleKeyPress(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > int(fieldDevourerLevel) {
			m.CurrentInput().Blur()
			m.cursor--
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldLeftoverShards) {
			m.CurrentInput().Blur()
			m.cursor++
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "left", "right":
		field := m.currentField()
		if len(field.Options) > 0 {
			cur := m.CurrentInput().Value()
			idx := slices.Index(field.Options, cur)
			if idx < 0 {
				idx = 0
			}
			if msg.String() == "left" && idx > 0 {
				idx--
			} else if msg.String() == "right" && idx < len(field.Options)-1 {
				idx++
			} else {
				return m, nil
			}
			newVal := field.Options[idx]
			m.CurrentInput().SetValue(newVal)
			m.parseBaseTabFieldValues(fieldIndex(m.cursor), newVal)
			return m, nil
		} else if field.Step > 0 {
			cur, err := strconv.Atoi(m.CurrentInput().Value())
			if err != nil {
				return m, nil
			}
			if msg.String() == "left" {
				cur -= field.Step
			} else {
				cur += field.Step
			}
			newVal := strconv.Itoa(cur)
			if field.Validate != nil {
				if err := field.Validate(newVal); err != nil {
					return m, nil
				}
			}
			m.CurrentInput().SetValue(newVal)
			m.parseBaseTabFieldValues(fieldIndex(m.cursor), newVal)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.currentField().Input, cmd = m.CurrentInput().Update(msg)
	if m.CurrentInput().Err == nil {
		m.parseBaseTabFieldValues(fieldIndex(m.cursor), m.CurrentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseBaseTabFieldValues(i fieldIndex, value string) {
	switch i {
	case fieldDevourerLevel:
		if v, err := strconv.Atoi(value); err == nil {
			m.DevourerLevel = v
		}
	case fieldFeatTiers:
		if v, err := strconv.Atoi(value); err == nil {
			m.FeatTiers = v
		}
	case fieldOtherMultiplier:
		if v, err := strconv.Atoi(value); err == nil {
			m.OtherMultiplier = float64(v) / 100.0
		}
	case fieldGroupBonusCount:
		if v, err := strconv.Atoi(value); err == nil {
			m.GroupBonusCount = v
		}
	case fieldLeftoverShards:
		if v, err := strconv.Atoi(value); err == nil {
			m.LeftoverShards = v
		}
	}
}
