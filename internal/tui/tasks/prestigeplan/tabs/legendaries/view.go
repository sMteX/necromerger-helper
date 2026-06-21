package legendaries

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

var (
	nameColumn  = lipgloss.NewStyle().Width(12)
	countColumn = lipgloss.NewStyle().Width(8).AlignHorizontal(lipgloss.Right)
	bonusColumn = lipgloss.NewStyle().Width(8).AlignHorizontal(lipgloss.Right).MarginLeft(1).MarginRight(3)
	arrow       = lipgloss.NewStyle().Render("  →  ")
)

func (m *Model) View() string {
	tableWidth := nameColumn.GetWidth() + 2*(countColumn.GetWidth()+countColumn.GetHorizontalFrameSize()) + lipgloss.Width(arrow) + bonusColumn.GetWidth() + bonusColumn.GetHorizontalFrameSize()

	lines := []string{
		lipgloss.JoinHorizontal(lipgloss.Top,
			nameColumn.Bold(true).Render("Legendary"),
			countColumn.Bold(true).Render("Have  "),
			arrow,
			countColumn.Bold(true).Render("Plan  "),
			bonusColumn.Bold(true).Render("Bonus"),
		),
		m.renderGroupHeading(models.Group1, tableWidth),
		m.renderRow(models.Lich),
		m.renderRow(models.Gorgon),
		m.renderRow(models.Harpy),
		m.renderGroupHeading(models.Group2, tableWidth),
		m.renderRow(models.Reaper),
		m.renderRow(models.Cyclops),
		m.renderRow(models.Archdemon),
		m.renderGroupHeading(models.Group3, tableWidth),
		m.renderRow(models.TheCursed),
		m.renderRow(models.TheColossus),
		m.renderRow(models.TheInfernal),
		m.renderGroupHeading(models.Group4, tableWidth),
		m.renderRow(models.RoboChicken),
		m.renderRow(models.ShieldBot),
		m.renderRow(models.SoulStalker),
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m *Model) renderGroupHeading(group models.LegendaryGroup, tableWidth int) string {
	groupHeading := lipgloss.NewStyle().Foreground(shared.Colors.Good).MarginTop(1)
	return groupHeading.Render(groupHeadingText(fmt.Sprintf("Group %d", group), shared.FormatPercentageBonus(m.LegendaryGroupBonuses[group]), tableWidth))
}

var legendaryToFieldInputs = map[models.LegendaryID][2]fieldIndex{
	models.Lich:        {fieldLichHave, fieldLichPlan},
	models.Gorgon:      {fieldGorgonHave, fieldGorgonPlan},
	models.Harpy:       {fieldHarpyHave, fieldHarpyPlan},
	models.Reaper:      {fieldReaperHave, fieldReaperPlan},
	models.Cyclops:     {fieldCyclopsHave, fieldCyclopsPlan},
	models.Archdemon:   {fieldArchdemonHave, fieldArchdemonPlan},
	models.TheCursed:   {fieldTheCursedHave, fieldTheCursedPlan},
	models.TheColossus: {fieldTheColossusHave, fieldTheColossusPlan},
	models.TheInfernal: {fieldTheInfernalHave, fieldTheInfernalPlan},
	models.RoboChicken: {fieldRoboChickenHave, fieldRoboChickenPlan},
	models.ShieldBot:   {fieldShieldBotHave, fieldShieldBotPlan},
	models.SoulStalker: {fieldSoulStalkerHave, fieldSoulStalkerPlan},
}

func (m *Model) renderRow(legendary models.LegendaryID) string {
	fieldIndexes := legendaryToFieldInputs[legendary]

	haveColumn := func() lipgloss.Style {
		if m.PossessedLegendaries[legendary] >= m.LegendaryCounts[legendary] {
			return countColumn.Foreground(shared.Colors.Good)
		}
		return countColumn.Foreground(shared.Colors.Bad)
	}()

	haveField, planField := m.fields[fieldIndexes[0]], m.fields[fieldIndexes[1]]
	haveFieldText, planFieldText := "", ""
	haveFocused := m.cursor == int(fieldIndexes[0])
	planFocused := m.cursor == int(fieldIndexes[1])
	// have field text
	if haveFocused {
		haveFieldText = haveColumn.Render(shared.PadLeft("< "+haveField.Input.Value()+" >", haveColumn.GetWidth()))
	} else {
		haveFieldText = haveColumn.Render(shared.PadLeft(haveField.Input.Value()+"  ", haveColumn.GetWidth()))
	}
	// plan field text
	if planFocused {
		planFieldText = countColumn.Render(shared.PadLeft("< "+planField.Input.Value()+" >", countColumn.GetWidth()))
	} else {
		planFieldText = countColumn.Render(shared.PadLeft(planField.Input.Value()+"  ", countColumn.GetWidth()))
	}

	l := data.LegendariesById[legendary]
	return nameColumn.Render(l.Name) +
		haveFieldText +
		arrow +
		planFieldText +
		bonusColumn.Render(shared.FormatPercentageBonus(m.LegendaryBonuses[legendary]))
}

func groupHeadingText(name, bonus string, width int) string {
	// e.g. "── Group 3 ─────────────────── +80% ──"
	// 2 dashes on each side fixed
	// padding 1 space around each text
	// Group 4 = 7 chars, not incl. padding
	innerFillLength := width - 2*2 - 2*2 - lipgloss.Width(name) - lipgloss.Width(bonus)
	return fmt.Sprintf("── %s %s %s ──", name, strings.Repeat("─", innerFillLength), bonus)
}

func (m *Model) GetHelpItems() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  change legendary"),
		shared.Styles.Help.Render("Tab / Shift+Tab  change have/planned"),
		shared.Styles.Help.Render("← / →  change value"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
