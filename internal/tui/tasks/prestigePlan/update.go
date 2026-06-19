package prestigePlan

import (
	tea "charm.land/bubbletea/v2"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// bubbletea only sends WindowSizeMsg once, before the user picks a task,
		// so the prestige planner never gets Init() called by the framework.
		// Use the first WindowSizeMsg as the init signal to activate the cursor field.
		if m.selectedTab == planTabBase && !m.currentInput().Focused() {
			return m, m.currentInput().Focus()
		}
		return m, nil
	case tea.KeyPressMsg:
		return m.handleKey(msg)
	}
	// Non-key messages (cursor blink ticks) must reach the active textinput.
	if m.selectedTab == planTabBase {
		var cmd tea.Cmd
		m.fields[fieldIndex(m.cursor)].input, cmd = m.currentInput().Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "f1", "f2", "f3", "f4":
		prev := m.selectedTab
		switch msg.String() {
		case "f1":
			m.selectedTab = planTabBase
		case "f2":
			m.selectedTab = planTabLegendaries
		case "f3":
			m.selectedTab = planTabRunes
		case "f4":
			m.selectedTab = planTabExperiments
		}
		if prev == planTabBase && m.selectedTab != planTabBase {
			m.currentInput().Blur()
		} else if prev != planTabBase && m.selectedTab == planTabBase {
			return m, m.currentInput().Focus()
		}
		return m, nil
	}

	switch m.selectedTab {
	case planTabBase:
		return m.handleBaseTabKey(msg)
	}
	return m, nil
}
