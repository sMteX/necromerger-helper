package planmenu

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

// innerWidth is the content width passed to boxStyle.Width() on each render.
// The box border + padding adds to this, so the total modal width is larger.
const innerWidth = 30

var (
	boxStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(shared.Colors.Border).Padding(0, 2)
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(shared.Colors.Border).MarginBottom(1)
	selectedStyle = lipgloss.NewStyle().Foreground(shared.Colors.Good)
	hintStyle     = shared.Styles.Help.MarginTop(1)
)

func (m *Model) View() string {
	switch m.state {
	case stateMenu:
		return m.viewMenu()
	case stateList:
		return m.viewList()
	case stateNameEditor:
		return m.viewNameEditor()
	case stateConfirm:
		return m.viewConfirm()
	}
	return ""
}

func (m *Model) viewMenu() string {
	type entry struct {
		item  menuItem
		label string
	}
	entries := []entry{
		{menuNew, "New plan"},
		{menuLoad, "Load plan..."},
		{menuSaveAs, "Save as..."},
		{menuClose, "Close"},
	}

	lines := []string{titleStyle.Render("Plan"), ""}
	for _, e := range entries {
		if e.item == m.menuCursor {
			lines = append(lines, selectedStyle.Render("▶ "+e.label))
		} else {
			lines = append(lines, "  "+e.label)
		}
	}
	lines = append(lines, hintStyle.Render("x / Esc to close"))
	return boxStyle.Width(innerWidth).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m *Model) viewList() string {
	lines := []string{titleStyle.Render("Load plan"), ""}
	if len(m.cfg.Plans) == 0 {
		lines = append(lines, shared.Styles.Help.Render("(no saved plans)"))
	} else {
		for i, plan := range m.cfg.Plans {
			name := plan.Name
			if name == "" {
				name = "(unnamed)"
			}
			if i == m.listCursor {
				lines = append(lines, selectedStyle.Render("▶ "+name))
			} else {
				lines = append(lines, "  "+name)
			}
			if !plan.UpdatedAt.IsZero() {
				lines = append(lines, shared.Styles.Help.Render(fmt.Sprintf("    %s", plan.UpdatedAt.Format("2006-01-02 15:04"))))
			}
		}
	}
	lines = append(lines, hintStyle.Render("Enter to load  Esc to back"))
	return boxStyle.Width(innerWidth).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m *Model) viewNameEditor() string {
	lines := []string{
		titleStyle.Render(m.saveTitle),
		"",
		"Name",
		m.nameInput.View(),
		"",
		"Notes",
		m.notesInput.View(),
		hintStyle.Render("Tab to switch  Enter to save  Esc back"),
	}
	return boxStyle.Width(innerWidth).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m *Model) viewConfirm() string {
	lines := []string{
		titleStyle.Render("Unsaved changes"),
		"",
		"Discard current plan?",
		"",
	}
	opts := []struct {
		opt   confirmOption
		label string
	}{
		{confirmYes, "Yes, discard"},
		{confirmNo, "No, go back"},
	}
	for _, o := range opts {
		if o.opt == m.confirmCursor {
			lines = append(lines, selectedStyle.Render("▶ "+o.label))
		} else {
			lines = append(lines, "  "+o.label)
		}
	}
	return boxStyle.Width(innerWidth).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
