package runes

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
		if m.cursor > int(fieldIce) {
			m.CurrentInput().Blur()
			m.cursor--
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldCosmic) {
			m.CurrentInput().Blur()
			m.cursor++
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.currentField().Input, cmd = m.CurrentInput().Update(msg)
	if m.CurrentInput().Err == nil {
		m.parseFieldValues(fieldIndex(m.cursor), m.CurrentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseFieldValues(i fieldIndex, value string) {
	if v, err := strconv.Atoi(value); err == nil {
		m.PossessedRunes[runeByFieldType[i]] = v
	}
}
