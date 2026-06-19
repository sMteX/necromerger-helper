package prestigePlan

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "f1":
			m.selectedTab = planTabBase
		case "f2":
			m.selectedTab = planTabLegendaries
		case "f3":
			m.selectedTab = planTabRunes
		case "f4":
			m.selectedTab = planTabExperiments
		}
	}
	return m, nil
}
