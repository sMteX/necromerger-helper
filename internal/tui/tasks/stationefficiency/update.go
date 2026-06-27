package stationefficiency

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necromerger-helper/internal/models"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// bubbletea only sends WindowSizeMsg once, before the user picks a task,
		// so the prestige planner never gets Init() called by the framework.
		// Use the first WindowSizeMsg as the init signal to activate the cursor field.
		if !m.CurrentInput().Focused() {
			return m, m.CurrentInput().Focus()
		}
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			return m.handleKeyPress(msg)
		}
	}
	return m, m.HandleNonKeyMsg(msg)
}

func (m *Model) handleKeyPress(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		return m, m.HandleUpKey(int(fieldChampion))
	case "down":
		return m, m.HandleDownKey(int(fieldSpeedPercent))
	case "left", "right":
		if newVal, changed := m.HandleStepKeys(msg.String()); changed {
			m.parseFieldValues(fieldIndex(m.Cursor), newVal)
			m.recalculate()
		}
		return m, nil
	}

	cmd, changed := m.HandleInputKeyMsg(msg)
	if changed {
		m.parseFieldValues(fieldIndex(m.Cursor), m.CurrentInput().Value())
		m.recalculate()
	}
	return m, cmd
}

func (m *Model) parseFieldValues(i fieldIndex, value string) {
	switch i {
	case fieldChampion:
		m.selectedChampion = models.ChampionID(value)
	case fieldManaCap:
		if v, err := strconv.Atoi(value); err == nil {
			m.manaCap = v
		}
	case fieldSlimeCap:
		if v, err := strconv.Atoi(value); err == nil {
			m.slimeCap = v
		}
	case fieldDarknessCap:
		if v, err := strconv.Atoi(value); err == nil {
			m.darknessCap = v
		}
	case fieldSpeedPercent:
		if v, err := strconv.Atoi(value); err == nil {
			m.championSpeedPercent = v
		}
	}
}
