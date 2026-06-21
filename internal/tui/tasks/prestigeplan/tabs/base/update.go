package base

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
		return m, m.HandleUpKey(int(fieldDevourerLevel))
	case "down":
		return m, m.HandleDownKey(int(fieldLeftoverShards))
	case "left", "right":
		if newVal, changed := m.HandleStepKeys(msg.String()); changed {
			m.parseFieldValues(fieldIndex(m.Cursor), newVal)
		}
		return m, nil
	}

	cmd, changed := m.HandleInputKeyMsg(msg)
	if changed {
		m.parseFieldValues(fieldIndex(m.Cursor), m.CurrentInput().Value())
	}
	return m, cmd
}

func (m *Model) parseFieldValues(i fieldIndex, value string) {
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
