package experiments

import (
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
		if m.cursor > int(fieldSeasoning1) {
			m.CurrentInput().Blur()
			m.cursor--
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldCapacity2) {
			m.CurrentInput().Blur()
			m.cursor++
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "left", "right":
		field := m.currentField()
		if field.Step > 0 {
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
			m.parseFieldValues(fieldIndex(m.cursor), newVal)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.currentField().Input, cmd = m.CurrentInput().Update(msg)
	if m.CurrentInput().Err == nil {
		m.parseFieldValues(fieldIndex(m.cursor), m.CurrentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseFieldValues(i fieldIndex, value string) {
	e := experimentsByFieldIndex[i]
	if v, err := strconv.Atoi(value); err == nil {
		m.ExperimentLevels[e.ID] = v
	}
}
