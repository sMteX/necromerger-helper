package tui

import (
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m *AppModel) renderMainMenu() tea.View {
	content := lipgloss.JoinVertical(lipgloss.Center,
		styleMenuHeader.Render("NecroMerger helper"),
		m.renderMainMenuChoices(),
	)
	fw, fh := styleMainContainer.GetFrameSize()
	content = styleMainContainer.
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
