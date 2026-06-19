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

	// elements with more or less fixed size - will be subtracted from the available space to figure out the main content
	header := shared.Styles.Header.Render("Prestige planner")
	tabSelector := m.renderTabSelector()
	help := m.renderHelp()

	totalHeight := lipgloss.Height(header) + lipgloss.Height(tabSelector) + lipgloss.Height(help)
	mainContentHeight := m.windowHeight - totalHeight - fh
	summaryWidth := m.summaryWidth()

	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		tabSelector,
		lipgloss.JoinHorizontal(lipgloss.Top,
			// TODO: min main content size
			m.renderMainContent(m.windowWidth-fw-summaryWidth, mainContentHeight),
			m.renderSummary(summaryWidth, mainContentHeight), // 25 = min width, 16 = min height
		),
		// TODO: help, dynamic sizing, min terminal size
		help,
	)
	content = shared.Styles.MainContainer.
		Width(m.windowWidth).
		Height(m.windowHeight).
		//Align(lipgloss.Center, lipgloss.Top).
		Render(content)

	return tea.View{
		Content:   content,
		AltScreen: true,
	}
}
func (m Model) summaryWidth() int {
	availableWidth := m.windowWidth - shared.Styles.MainContainer.GetHorizontalFrameSize()
	// prefer around 20% of the window width, but >= 25
	preferred := math.Floor(float64(availableWidth) / 5.0)
	return int(math.Max(25, preferred))
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
		if int(m.selectedTab) == i {
			choices[i] = choiceSelectedStyle.Render(choices[i])
		} else {
			choices[i] = choiceStyle.Render(choices[i])
		}
	}
	return lipgloss.NewStyle().MarginBottom(1).Render(lipgloss.JoinHorizontal(lipgloss.Top, choices...))
}
func (m Model) renderSummary(summaryWidth, summaryHeight int) string {
	fw := shared.Styles.SubContainer.GetHorizontalFrameSize()
	var lines []string
	valueStyle := lipgloss.NewStyle().Width(10).AlignHorizontal(lipgloss.Right)
	labelStyle := lipgloss.NewStyle().Width(summaryWidth - fw - 10)

	lines = append(lines, shared.Styles.Header.Render("Time Shard Summary"))

	values := [][]string{
		{"Base", shared.FormatNumberLong(m.calculatedOutputs.summary.baseShards)},
		{"Leftovers", shared.FormatNumberLong(m.calculatedOutputs.summary.leftoverShards)},
		{"× Feats", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.summary.featMultiplier*100))},
		{"× Leggos", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.summary.legendariesMultiplier*100))},
		{"× Others", fmt.Sprintf("+%.0f%%", math.Floor(m.calculatedOutputs.summary.othersMultiplier*100))},
	}
	for _, v := range values {
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, labelStyle.Render(v[0]), valueStyle.Render(v[1])))
	}
	netStyle := func() lipgloss.Style {
		if m.calculatedOutputs.summary.netShards > 0 {
			return valueStyle.Foreground(shared.Colors.Good)
		}
		return valueStyle.Foreground(shared.Colors.Bad)
	}()
	lines = append(lines,
		strings.Repeat("─", summaryWidth-fw),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Render("Total"),
			valueStyle.Render(shared.FormatNumberLong(m.calculatedOutputs.summary.totalShards)),
		),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Foreground(shared.Colors.Bad).Render("Spent"),
			valueStyle.Foreground(shared.Colors.Bad).Render(shared.FormatNumberLong(-m.calculatedOutputs.summary.spentShards)),
		),
		strings.Repeat("─", summaryWidth-fw),
		lipgloss.JoinHorizontal(lipgloss.Top,
			labelStyle.Render("Net"),
			netStyle.Render(shared.FormatNumberLong(m.calculatedOutputs.summary.netShards)),
		),
	)
	return shared.Styles.SubContainer.
		Width(summaryWidth).
		Height(summaryHeight).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}

func (m Model) renderMainContent(maxWidth, maxHeight int) string {
	var content string
	switch m.selectedTab {
	case planTabBase:
		content = m.renderBaseTab()
	case planTabLegendaries:
		content = m.renderLegendariesTab()
	case planTabRunes:
		content = m.renderRuneTab()
	case planTabExperiments:
		content = m.renderExperimentsTab()
	default:
		content = "Main content placeholder"
	}
	return shared.Styles.SubContainer.Width(maxWidth).Height(maxHeight).Render(content)
}

func (m Model) renderHelp() string {
	maxWidth := m.windowWidth - shared.Styles.MainContainer.GetHorizontalFrameSize()

	var units []string
	switch m.selectedTab {
	case planTabBase:
		units = m.getBaseTabHelp()
	case planTabLegendaries:
		units = m.getLegendariesTabHelp()
	case planTabRunes:
		units = m.getRuneTabHelp()
	case planTabExperiments:
		units = m.getExperimentsTabHelp()
	default:
		units = []string{
			shared.Styles.Help.Render("↑ / ↓  navigate"),
			shared.Styles.Help.Render("← / →  navigate"),
			shared.Styles.Help.Render("F1 - F4  switch tab"),
			shared.Styles.Help.Render("q / ctrl+c  exit"),
		}
	}
	var lines []string

	separator := shared.Styles.Help.Render("  ·  ")
	separatorWidth := lipgloss.Width(separator)
	line := ""
	lineWidth := 0
	for _, u := range units {
		unitWidth := lipgloss.Width(u)
		if line == "" {
			line, lineWidth = u, unitWidth
		} else if lineWidth+separatorWidth+unitWidth <= maxWidth {
			line += separator + u
			lineWidth += separatorWidth + unitWidth
		} else {
			lines = append(lines, line)
			line, lineWidth = u, unitWidth
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lipgloss.NewStyle().MarginTop(1).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
