package legendaries

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
		return m, m.HandleUpKeyFn(func(cursor int) bool {
			return cursor > int(fieldLichHave) && cursor != int(fieldLichPlan)
		})
	case "down":
		return m, m.HandleDownKeyFn(func(cursor int) bool {
			return cursor < int(fieldSoulStalkerPlan) && cursor != int(fieldSoulStalkerHave)
		})
	case "tab":
		if m.Cursor >= int(fieldLichHave) && m.Cursor <= int(fieldSoulStalkerHave) {
			m.CurrentInput().Blur()
			m.Cursor += int(fieldLichPlan) - int(fieldLichHave)
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case "shift+tab":
		if m.Cursor >= int(fieldLichPlan) && m.Cursor <= int(fieldSoulStalkerPlan) {
			m.CurrentInput().Blur()
			m.Cursor -= int(fieldLichPlan) - int(fieldLichHave)
			return m, m.CurrentInput().Focus()
		}
		return m, nil
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
	legendary := legendaryIdByFieldIndex[i]
	if i >= fieldLichHave && i <= fieldSoulStalkerHave {
		if v, err := strconv.Atoi(value); err == nil {
			m.PossessedLegendaries[legendary] = v
		}
		return
	}
	if v, err := strconv.Atoi(value); err == nil {
		m.LegendaryCounts[legendary] = v
	}
}
