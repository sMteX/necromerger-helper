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
	return m, m.HandleNonKeyMsg(msg)
}

func (m *Model) handleKeyPress(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		return m, m.HandleUpKey(int(fieldIce))
	case "down":
		return m, m.HandleDownKey(int(fieldCosmic))
	}

	cmd, changed := m.HandleInputKeyMsg(msg)
	if changed {
		m.parseFieldValues(fieldIndex(m.Cursor), m.CurrentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseFieldValues(i fieldIndex, value string) {
	if v, err := strconv.Atoi(value); err == nil {
		m.PossessedRunes[runeByFieldType[i]] = v
	}
}
