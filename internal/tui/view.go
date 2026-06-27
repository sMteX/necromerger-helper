package tui

import (
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

func (m *AppModel) renderMainMenu() tea.View {
	content := lipgloss.JoinVertical(lipgloss.Center,
		shared.Styles.Header.Render("NecroMerger helper"),
		m.renderMainMenuChoices(),
	)
	content = shared.Styles.MainContainer.
		Width(m.windowWidth).
		Height(m.windowHeight).
		Align(lipgloss.Center, lipgloss.Top).
		Render(content)

	return tea.View{
		Content:   content,
		AltScreen: true,
	}
}

func (m *AppModel) renderMainMenuChoices() string {
	choices := []string{"Resource cap", "Prestige plan", "Station efficiency"}
	var styledChoices []string
	for i, choice := range choices {
		if i == m.cursor {
			styledChoices = append(styledChoices, styleActiveOption.Render(choice))
		} else {
			styledChoices = append(styledChoices, styleInactiveOption.Render(choice))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, styledChoices...)
}
