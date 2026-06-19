package prestigePlan

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m Model) renderLegendariesTab() string {
	// TODO: clean up this duplicated style mess
	nameColumn := lipgloss.NewStyle().Width(15)
	countColumn := lipgloss.NewStyle().Width(4).AlignHorizontal(lipgloss.Right)
	bonusColumn := lipgloss.NewStyle().Width(8 + 3).AlignHorizontal(lipgloss.Right)
	arrow := lipgloss.NewStyle().Render("  →  ")

	tableWidth := nameColumn.GetWidth() + 2*countColumn.GetWidth() + lipgloss.Width(arrow) + bonusColumn.GetWidth()
	rightPadding := "   "

	//var lines []string
	lines := []string{
		nameColumn.Bold(true).Render("Legendary") + countColumn.Bold(true).Render("Have") + arrow + countColumn.Bold(true).Render("Plan") + bonusColumn.Bold(true).Render("Bonus"+rightPadding),
		m.renderLegendaryGroupHeading(models.Group1, tableWidth),
		m.renderLegendaryRow(models.Lich),
		m.renderLegendaryRow(models.Gorgon),
		m.renderLegendaryRow(models.Harpy),
		m.renderLegendaryGroupHeading(models.Group2, tableWidth),
		m.renderLegendaryRow(models.Reaper),
		m.renderLegendaryRow(models.Cyclops),
		m.renderLegendaryRow(models.Archdemon),
		m.renderLegendaryGroupHeading(models.Group3, tableWidth),
		m.renderLegendaryRow(models.TheCursed),
		m.renderLegendaryRow(models.TheColossus),
		m.renderLegendaryRow(models.TheInfernal),
		m.renderLegendaryGroupHeading(models.Group4, tableWidth),
		m.renderLegendaryRow(models.RoboChicken),
		m.renderLegendaryRow(models.ShieldBot),
		m.renderLegendaryRow(models.SoulStalker),
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		lines...,
	)
}

func (m Model) renderLegendaryGroupHeading(group models.LegendaryGroup, tableWidth int) string {
	groupHeading := lipgloss.NewStyle().Foreground(shared.Colors.Good).MarginTop(1)
	return groupHeading.Render(renderGroupHeadingText(fmt.Sprintf("Group %d", group), shared.FormatPercentageBonus(m.calculatedOutputs.legendaryGroupBonuses[group]), tableWidth))
}

func (m Model) renderLegendaryRow(legendary models.LegendaryID) string {
	nameColumn := lipgloss.NewStyle().Width(15)
	countColumn := lipgloss.NewStyle().Width(4).AlignHorizontal(lipgloss.Right)
	bonusColumn := lipgloss.NewStyle().Width(8 + 3).AlignHorizontal(lipgloss.Right)
	arrow := lipgloss.NewStyle().Render("  →  ")
	rightPadding := "   "

	haveColumn := func() lipgloss.Style {
		if m.currentLegendaries[legendary] >= m.plannedLegendaries[legendary] {
			return countColumn.Foreground(shared.Colors.Good)
		}
		return countColumn.Foreground(shared.Colors.Bad)
	}()

	l := data.LegendariesById[legendary]
	return nameColumn.Render(l.Name) +
		haveColumn.Render(fmt.Sprintf("%d", m.currentLegendaries[legendary])) +
		arrow +
		countColumn.Render(fmt.Sprintf("%d", m.plannedLegendaries[legendary])) +
		bonusColumn.Render(shared.FormatPercentageBonus(m.calculatedOutputs.legendaryBonuses[legendary])+rightPadding)
}

func renderGroupHeadingText(name, bonus string, width int) string {
	// e.g. "── Group 3 ─────────────────── +80% ──"
	// 2 dashes on each side fixed
	// padding 1 space around each text
	// Group 4 = 7 chars, not incl. padding
	innerFillLength := width - 2*2 - 2*2 - lipgloss.Width(name) - lipgloss.Width(bonus)
	return fmt.Sprintf("── %s %s %s ──", name, strings.Repeat("─", innerFillLength), bonus)
}

func (m Model) getLegendariesTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("← / →  navigate"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
