package experiments

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necromerger-helper/internal/models"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

var (
	nameColumn        = lipgloss.NewStyle().Width(20)
	levelColumn       = lipgloss.NewStyle().Width(9).AlignHorizontal(lipgloss.Center)
	effectColumn      = lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center)
	currentCostColumn = lipgloss.NewStyle().Width(6).AlignHorizontal(lipgloss.Right)
	nextCostColumn    = lipgloss.NewStyle().Width(14).MarginLeft(3)
)

func (m *Model) View() string {
	tableWidth := nameColumn.GetWidth() + levelColumn.GetWidth() + effectColumn.GetWidth() + currentCostColumn.GetWidth() + nextCostColumn.GetWidth() + nextCostColumn.GetHorizontalFrameSize()
	//arrow := " → "

	lines := []string{
		lipgloss.NewStyle().Bold(true).Render(
			nameColumn.Render("Experiment") + levelColumn.Render("Level") + effectColumn.Render("Effect") + currentCostColumn.Render("Cost") + nextCostColumn.Render("Next level"),
		),
		renderGroupHeader("Pre-100", tableWidth),
		m.renderRow(fieldSeasoning1),
		m.renderRow(fieldStrength1),
		m.renderRow(fieldTaste1),
		m.renderRow(fieldCapacity1),
		m.renderRow(fieldBodySnatcher),
		m.renderRow(fieldWeakening),
		m.renderRow(fieldDamageCap),
		m.renderRow(fieldIceChest),
		m.renderRow(fieldPoisonChest),
		m.renderRow(fieldBloodChest),
		m.renderRow(fieldMoonChest),
		m.renderRow(fieldDeathChest),
		m.renderRow(fieldCosmicChest),
		renderGroupHeader("Post-100", tableWidth),
		m.renderRow(fieldSeasoning2),
		m.renderRow(fieldStrength2),
		m.renderRow(fieldTaste2),
		m.renderRow(fieldCapacity2),
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		lines...,
	)
}

func renderGroupHeader(text string, width int) string {
	// e.g. "── Group 3 ─────────────────────"
	// 2 dashes on left side fixed
	// padding 1 space around text
	fillLength := width - 2 - 2 - lipgloss.Width(text)
	return lipgloss.NewStyle().Foreground(shared.Colors.Good).MarginTop(1).Render(fmt.Sprintf("── %s %s", text, strings.Repeat("─", fillLength)))
}

func (m *Model) renderRow(i fieldIndex) string {
	e := experimentsByFieldIndex[i]
	// TODO: careful with indices, `plannedExperiments[...]` can be 0 - not planned
	var currentLevel *models.ExperimentLevel
	if m.ExperimentLevels[e.ID] > 0 {
		// let's assume the `plannedExperiments[]` is either 0 (not planned) or in bounds (after subtracting 1)
		currentLevel = &e.Levels[m.ExperimentLevels[e.ID]-1]
	}
	var nextLevel *models.ExperimentLevel
	if m.ExperimentLevels[e.ID] < len(e.Levels) {
		nextLevel = &e.Levels[m.ExperimentLevels[e.ID]]
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
		m.renderExperimentsRowLevelInput(i),
		effectColumn.Render(m.getEffectText(e.ID, e.Tier, currentLevel)),
		currentCostColumn.Render(currentCost),
		nextCostColumn.Foreground(shared.Colors.Dim).Render(nextCost),
	)
}

func (m *Model) renderExperimentsRowLevelInput(i fieldIndex) string {
	e := experimentsByFieldIndex[i]
	if m.Cursor == int(i) {
		current := m.CurrentInput().Value()
		valueText := "none"
		if current != "0" {
			valueText = fmt.Sprintf("lvl %s", current)
		}
		return levelColumn.Foreground(shared.Colors.Good).Render(fmt.Sprintf("< %s >", valueText))
	}
	return levelColumn.Render(fmt.Sprintf("lvl %d", m.ExperimentLevels[e.ID]))
}

func (m *Model) getEffectText(experiment models.ExperimentID, tier models.ExperimentTier, level *models.ExperimentLevel) string {
	arrow := " → "
	if level == nil {
		return "──"
	}
	return lipgloss.JoinHorizontal(lipgloss.Top,
		shared.FormatExperimentValue(experiment, tier, level.PrevValue),
		arrow,
		shared.FormatExperimentValue(experiment, tier, level.Value),
	)
}

func (m *Model) GetHelpItems() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  change experiment"),
		shared.Styles.Help.Render("← / →  change level"),
	}
}
