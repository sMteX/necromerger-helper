package prestigePlan

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m Model) renderExperimentsTab() string {
	// TODO: clean up this duplicated style mess
	nameColumn := lipgloss.NewStyle().Width(20)
	levelColumn := lipgloss.NewStyle().Width(5).AlignHorizontal(lipgloss.Center)
	effectColumn := lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center)
	currentCostColumn := lipgloss.NewStyle().Width(6).AlignHorizontal(lipgloss.Right)
	nextCostColumn := lipgloss.NewStyle().Width(14).MarginLeft(3)

	tableWidth := nameColumn.GetWidth() + levelColumn.GetWidth() + effectColumn.GetWidth() + currentCostColumn.GetWidth() + nextCostColumn.GetWidth() + nextCostColumn.GetHorizontalFrameSize()
	//arrow := " → "

	lines := []string{
		lipgloss.NewStyle().Bold(true).Render(
			nameColumn.Render("Experiment") + levelColumn.Render("Level") + effectColumn.Render("Effect") + currentCostColumn.Render("Cost") + nextCostColumn.Render("Next level"),
		),
		renderExperimentHeadingText("Pre-100", tableWidth),
		m.renderExperimentRow(models.ExpSeasoning),
		m.renderExperimentRow(models.ExpStrength),
		m.renderExperimentRow(models.ExpTaste),
		m.renderExperimentRow(models.ExpCapacity),
		m.renderExperimentRow(models.ExpBodySnatcher),
		m.renderExperimentRow(models.ExpWeakening),
		m.renderExperimentRow(models.ExpDamageCap),
		m.renderExperimentRow(models.ExpIceChest),
		m.renderExperimentRow(models.ExpPoisonChest),
		m.renderExperimentRow(models.ExpBloodChest),
		m.renderExperimentRow(models.ExpMoonChest),
		m.renderExperimentRow(models.ExpDeathChest),
		m.renderExperimentRow(models.ExpCosmicChest),
		renderExperimentHeadingText("Post-100", tableWidth),
		m.renderExperimentRow(models.ExpSeasoning2),
		m.renderExperimentRow(models.ExpStrength2),
		m.renderExperimentRow(models.ExpTaste2),
		m.renderExperimentRow(models.ExpCapacity2),
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		lines...,
	)
}

func renderExperimentHeadingText(text string, width int) string {
	// e.g. "── Group 3 ─────────────────────"
	// 2 dashes on left side fixed
	// padding 1 space around text
	fillLength := width - 2 - 2 - lipgloss.Width(text)
	return lipgloss.NewStyle().Foreground(shared.Colors.Good).MarginTop(1).Render(fmt.Sprintf("── %s %s", text, strings.Repeat("─", fillLength)))
}

func (m Model) renderExperimentRow(experiment models.ExperimentID) string {
	nameColumn := lipgloss.NewStyle().Width(20)
	levelColumn := lipgloss.NewStyle().Width(5).AlignHorizontal(lipgloss.Center)
	effectColumn := lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center)
	currentCostColumn := lipgloss.NewStyle().Width(6).AlignHorizontal(lipgloss.Right)
	nextCostColumn := lipgloss.NewStyle().Foreground(shared.Colors.Dim).Width(14).MarginLeft(3)

	e := data.ExperimentsById[experiment]
	// TODO: careful with indices, `plannedExperiments[...]` can be 0 - not planned
	var currentLevel *models.ExperimentLevel
	if m.plannedExperiments[experiment] > 0 {
		// let's assume the `plannedExperiments[]` is either 0 (not planned) or in bounds (after subtracting 1)
		currentLevel = &e.Levels[m.plannedExperiments[experiment]-1]
	}
	var nextLevel *models.ExperimentLevel
	if m.plannedExperiments[experiment] < len(e.Levels) {
		nextLevel = &e.Levels[m.plannedExperiments[experiment]]
	}

	currentCost := "──"
	if currentLevel != nil {
		currentCost = shared.FormatLargeNumber(currentLevel.Cost)
	}

	nextCost := "(max)"
	if nextLevel != nil {
		nextCost = fmt.Sprintf("(next: %s)", shared.FormatLargeNumber(nextLevel.Cost))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		nameColumn.Render(e.Name),
		levelColumn.Render(fmt.Sprintf("lvl %d", m.plannedExperiments[experiment])),
		effectColumn.Render(m.getEffectText(experiment, e.Tier, currentLevel)),
		currentCostColumn.Render(currentCost),
		nextCostColumn.Render(nextCost),
	)
}

func (m Model) getEffectText(experiment models.ExperimentID, tier models.ExperimentTier, level *models.ExperimentLevel) string {
	arrow := " → "
	if level == nil {
		return "──"
	}
	return fmt.Sprintf("%s%s%s",
		shared.FormatExperimentValue(experiment, tier, level.PrevValue),
		arrow,
		shared.FormatExperimentValue(experiment, tier, level.Value),
	)
}

func (m Model) getExperimentsTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("← / →  change level"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
