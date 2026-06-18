package tui

import (
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *AppModel) renderMainMenu() tea.View {
	content := lipgloss.JoinVertical(lipgloss.Center,
		shared.Styles.Header.Render("NecroMerger helper"),
		m.renderMainMenuChoices(),
	)
	fw, fh := shared.Styles.MainContainer.GetFrameSize()
	content = shared.Styles.MainContainer.
		Width(m.windowWidth-fw).
		Height(m.windowHeight-fh).
		Align(lipgloss.Center, lipgloss.Top).
		Render(content)

	return tea.View{
		Content: content,
	}
}

func (m *AppModel) renderMainMenuChoices() string {
	choices := []string{"Resource cap", "Prestige plan"}
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
