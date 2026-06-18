package prestigePlan

import (
	"fmt"
	"math"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m Model) View() tea.View {
	fw, fh := shared.Styles.MainContainer.GetFrameSize()
	content := lipgloss.JoinVertical(lipgloss.Left,
		shared.Styles.Header.Render("Prestige planner"),
		m.renderTabSelector(),
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.renderMainContent(80, 40),
			m.renderSummary(25, 16), // 25 = min width, 16 = min height
		),
		// TODO: help, dynamic sizing, min terminal size
	)
	content = shared.Styles.MainContainer.
		Width(m.windowWidth - fw).
		Height(m.windowHeight - fh).
		//Align(lipgloss.Center, lipgloss.Top).
		Render(content)

	return tea.View{
		Content: content,
	}
}

func (m Model) renderTabSelector() string {
	choiceStyle := lipgloss.NewStyle().PaddingRight(3)
	choiceSelectedStyle := choiceStyle.Bold(true).Foreground(shared.Colors.Good)
	choices := []string{
		"[F1] Base",
		"[F2] Legendaries",
		"[F3] Runes",
		"[F4] Experiments",
	}
	for i := range len(choices) {
		if m.selectedTab == i {
			choices[i] = choiceSelectedStyle.Render(choices[i])
		} else {
			choices[i] = choiceStyle.Render(choices[i])
		}
	}
	return lipgloss.NewStyle().MarginBottom(1).Render(lipgloss.JoinHorizontal(lipgloss.Top, choices...))
}
func (m Model) renderSummary(summaryWidth, summaryHeight int) string {
	summaryContainerStyle := shared.Styles.MainContainer.Padding(0, 2)
	fw, fh := summaryContainerStyle.GetFrameSize()
	var lines []string
	valueStyle := lipgloss.NewStyle().Width(10).Align(lipgloss.Right)
	labelStyle := lipgloss.NewStyle().Width(summaryWidth - fw - 10)

	lines = append(lines, shared.Styles.Header.Render("Time Shard Summary"))

	values := [][]string{
		{"Base", shared.FormatNumberLong(m.calculatedOutputs.baseShards)},
		{"Leftovers", shared.FormatNumberLong(m.calculatedOutputs.leftoverShards)},
		{"× Feats", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.featMultiplier*100))},
		{"× Leggos", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.legendariesMultiplier*100))},
		{"× Others", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.othersMultiplier*100))},
	}
	for _, v := range values {
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, labelStyle.Render(v[0]), valueStyle.Render(v[1])))
	}
	netStyle := func() lipgloss.Style {
		if m.calculatedOutputs.netShards > 0 {
			return valueStyle.Foreground(shared.Colors.Good)
		}
		return valueStyle.Foreground(shared.Colors.Bad)
	}()
	lines = append(lines,
		strings.Repeat("─", summaryWidth-fw),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Render("Total"),
			valueStyle.Render(shared.FormatNumberLong(m.calculatedOutputs.totalShards)),
		),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Foreground(shared.Colors.Bad).Render("Spent"),
			valueStyle.Foreground(shared.Colors.Bad).Render(shared.FormatNumberLong(-m.calculatedOutputs.spentShards)),
		),
		strings.Repeat("─", summaryWidth-fw),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Render("Net"),
			netStyle.Render(shared.FormatNumberLong(m.calculatedOutputs.netShards)),
		),
	)
	return summaryContainerStyle.
		// width is baked into the styles
		Height(summaryHeight - fh).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m Model) renderMainContent(maxWidth, maxHeight int) string {
	return ""
}
