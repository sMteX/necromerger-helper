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
			// The submodel needs a WindowSizeMsg to initialize its layout, but bubbletea
			// only sends one at program start — before any task is selected. Synthesize it
			// so the submodel receives it on the next update cycle.
			w, h := m.windowWidth, m.windowHeight
			return m, func() tea.Msg {
				return tea.WindowSizeMsg{Width: w, Height: h}
			}
		}
	}
	return m, nil
}
