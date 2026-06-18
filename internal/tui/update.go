package tui

import "charm.land/bubbletea/v2"

func (m *AppModel) handleUpdateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			// TODO: dynamic, menu items
			if m.cursor < 1 {
				m.cursor++
			}
		case "enter":
			var t taskType
			switch m.cursor {
			case 0:
				t = taskResourceCap
			case 1:
				t = taskPrestigePlan
			}
			m.currentTask = &t
		}
	}
	return m, nil
}
