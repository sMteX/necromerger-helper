package prestigeplan

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *Model) addLegendariesTabFields() {
	legendaryCountLimits := map[models.LegendaryID][2]int{
		models.Lich:        {0, 35},
		models.Gorgon:      {0, 35},
		models.Harpy:       {0, 35},
		models.Reaper:      {0, 35},
		models.Cyclops:     {0, 35},
		models.Archdemon:   {0, 4},
		models.TheCursed:   {0, 1},
		models.TheColossus: {0, 1},
		models.TheInfernal: {0, 1},
		models.RoboChicken: {0, 4},
		models.ShieldBot:   {0, 4},
		models.SoulStalker: {0, 35},
	}
	for i := fieldLegendariesLichHave; i <= fieldLegendariesSoulStalkerHave; i++ {
		legendary := data.LegendariesById[legendaryIdByFieldIndex[i]]
		limits := legendaryCountLimits[legendary.ID]
		m.fields[i] = inputField{
			label:          legendary.Name,
			step:           1,
			width:          5,
			characterLimit: 2,
			initialValue:   strconv.Itoa(m.plan.PossessedLegendaries[legendary.ID]),
			validate:       inputValidationIntInRange(limits[0], limits[1]),
		}
	}
	for i := fieldLegendariesLichPlan; i <= fieldLegendariesSoulStalkerPlan; i++ {
		legendary := data.LegendariesById[legendaryIdByFieldIndex[i]]
		limits := legendaryCountLimits[legendary.ID]
		m.fields[i] = inputField{
			label:          legendary.Name,
			step:           1,
			width:          5,
			characterLimit: 2,
			initialValue:   strconv.Itoa(m.plan.LegendaryCounts[legendary.ID]),
			validate:       inputValidationIntInRange(limits[0], limits[1]),
		}
	}
}

var (
	legendaryTabNameColumn  = lipgloss.NewStyle().Width(12)
	legendaryTabCountColumn = lipgloss.NewStyle().Width(8).AlignHorizontal(lipgloss.Right)
	legendaryTabBonusColumn = lipgloss.NewStyle().Width(8).AlignHorizontal(lipgloss.Right).MarginLeft(1).MarginRight(3)
	legendaryTabArrow       = lipgloss.NewStyle().Render("  →  ")
)

func (m *Model) renderLegendariesTab() string {
	tableWidth := legendaryTabNameColumn.GetWidth() + 2*(legendaryTabCountColumn.GetWidth()+legendaryTabCountColumn.GetHorizontalFrameSize()) + lipgloss.Width(legendaryTabArrow) + legendaryTabBonusColumn.GetWidth() + legendaryTabBonusColumn.GetHorizontalFrameSize()

	lines := []string{
		lipgloss.JoinHorizontal(lipgloss.Top,
			legendaryTabNameColumn.Bold(true).Render("Legendary"),
			legendaryTabCountColumn.Bold(true).Render("Have  "),
			legendaryTabArrow,
			legendaryTabCountColumn.Bold(true).Render("Plan  "),
			legendaryTabBonusColumn.Bold(true).Render("Bonus"),
		),
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
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m *Model) renderLegendaryGroupHeading(group models.LegendaryGroup, tableWidth int) string {
	groupHeading := lipgloss.NewStyle().Foreground(shared.Colors.Good).MarginTop(1)
	return groupHeading.Render(renderGroupHeadingText(fmt.Sprintf("Group %d", group), shared.FormatPercentageBonus(m.result.LegendaryGroupBonuses[group]), tableWidth))
}

var legendaryToFieldInputs = map[models.LegendaryID][2]fieldIndex{
	models.Lich:        {fieldLegendariesLichHave, fieldLegendariesLichPlan},
	models.Gorgon:      {fieldLegendariesGorgonHave, fieldLegendariesGorgonPlan},
	models.Harpy:       {fieldLegendariesHarpyHave, fieldLegendariesHarpyPlan},
	models.Reaper:      {fieldLegendariesReaperHave, fieldLegendariesReaperPlan},
	models.Cyclops:     {fieldLegendariesCyclopsHave, fieldLegendariesCyclopsPlan},
	models.Archdemon:   {fieldLegendariesArchdemonHave, fieldLegendariesArchdemonPlan},
	models.TheCursed:   {fieldLegendariesTheCursedHave, fieldLegendariesTheCursedPlan},
	models.TheColossus: {fieldLegendariesTheColossusHave, fieldLegendariesTheColossusPlan},
	models.TheInfernal: {fieldLegendariesTheInfernalHave, fieldLegendariesTheInfernalPlan},
	models.RoboChicken: {fieldLegendariesRoboChickenHave, fieldLegendariesRoboChickenPlan},
	models.ShieldBot:   {fieldLegendariesShieldBotHave, fieldLegendariesShieldBotPlan},
	models.SoulStalker: {fieldLegendariesSoulStalkerHave, fieldLegendariesSoulStalkerPlan},
}

func (m *Model) renderLegendaryRow(legendary models.LegendaryID) string {
	fieldIndexes := legendaryToFieldInputs[legendary]

	haveColumn := func() lipgloss.Style {
		if m.plan.PossessedLegendaries[legendary] >= m.plan.LegendaryCounts[legendary] {
			return legendaryTabCountColumn.Foreground(shared.Colors.Good)
		}
		return legendaryTabCountColumn.Foreground(shared.Colors.Bad)
	}()

	haveField, planField := m.fields[fieldIndexes[0]], m.fields[fieldIndexes[1]]
	haveFieldText, planFieldText := "", ""
	haveFocused := m.cursor == int(fieldIndexes[0])
	planFocused := m.cursor == int(fieldIndexes[1])
	// have field text
	if haveFocused {
		haveFieldText = haveColumn.Render(shared.PadLeft("< "+haveField.input.Value()+" >", haveColumn.GetWidth()))
	} else {
		haveFieldText = haveColumn.Render(shared.PadLeft(haveField.input.Value()+"  ", haveColumn.GetWidth()))
	}
	// plan field text
	if planFocused {
		planFieldText = legendaryTabCountColumn.Render(shared.PadLeft("< "+planField.input.Value()+" >", legendaryTabCountColumn.GetWidth()))
	} else {
		planFieldText = legendaryTabCountColumn.Render(shared.PadLeft(planField.input.Value()+"  ", legendaryTabCountColumn.GetWidth()))
	}

	l := data.LegendariesById[legendary]
	return legendaryTabNameColumn.Render(l.Name) +
		haveFieldText +
		legendaryTabArrow +
		planFieldText +
		legendaryTabBonusColumn.Render(shared.FormatPercentageBonus(m.result.LegendaryBonuses[legendary]))
}

func renderGroupHeadingText(name, bonus string, width int) string {
	// e.g. "── Group 3 ─────────────────── +80% ──"
	// 2 dashes on each side fixed
	// padding 1 space around each text
	// Group 4 = 7 chars, not incl. padding
	innerFillLength := width - 2*2 - 2*2 - lipgloss.Width(name) - lipgloss.Width(bonus)
	return fmt.Sprintf("── %s %s %s ──", name, strings.Repeat("─", innerFillLength), bonus)
}

func (m *Model) getLegendariesTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  change legendary"),
		shared.Styles.Help.Render("Tab / Shift+Tab  change have/planned"),
		shared.Styles.Help.Render("← / →  change value"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}
func (m *Model) handleLegendariesTabKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > int(fieldLegendariesLichHave) && m.cursor != int(fieldLegendariesLichPlan) {
			m.currentInput().Blur()
			m.cursor--
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldLegendariesSoulStalkerPlan) && m.cursor != int(fieldLegendariesSoulStalkerHave) {
			m.currentInput().Blur()
			m.cursor++
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "tab":
		if m.cursor >= int(fieldLegendariesLichHave) && m.cursor <= int(fieldLegendariesSoulStalkerHave) {
			m.currentInput().Blur()
			m.cursor += int(fieldLegendariesLichPlan) - int(fieldLegendariesLichHave)
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "shift+tab":
		if m.cursor >= int(fieldLegendariesLichPlan) && m.cursor <= int(fieldLegendariesSoulStalkerPlan) {
			m.currentInput().Blur()
			m.cursor -= int(fieldLegendariesLichPlan) - int(fieldLegendariesLichHave)
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "left", "right":
		// For arrow-adjustable fields, ←/→ increment/decrement the value directly.
		// For text-only fields (step == 0), fall through so the textinput handles
		// cursor movement within the text.
		field := m.fields[fieldIndex(m.cursor)]
		if field.step > 0 {
			cur, err := strconv.Atoi(m.currentInput().Value())
			if err != nil {
				return m, nil
			}
			if msg.String() == "left" {
				cur -= field.step
			} else {
				cur += field.step
			}
			newVal := strconv.Itoa(cur)
			if field.validate != nil {
				if err := field.validate(newVal); err != nil {
					// didn't pass validate, don't change anything
					return m, nil
				}
			}
			m.currentInput().SetValue(newVal)
			m.parseLegendariesTabFieldValues(fieldIndex(m.cursor), newVal)
			// TODO: recalculate m.calculatedOutputs from m.baseInputs
			return m, nil
		}
		return m, nil
	}

	// Everything else — character input, backspace, and ←/→ cursor movement for
	// text-only fields — goes to the focused textinput.
	var cmd tea.Cmd
	m.fields[m.cursor].input, cmd = m.currentInput().Update(msg)
	if m.currentInput().Err == nil {
		m.parseLegendariesTabFieldValues(fieldIndex(m.cursor), m.currentInput().Value())
		// TODO: recalculate m.calculatedOutputs from m.baseInputs
	}
	return m, cmd
}

var legendaryIdByFieldIndex = map[fieldIndex]models.LegendaryID{
	fieldLegendariesLichHave:        models.Lich,
	fieldLegendariesGorgonHave:      models.Gorgon,
	fieldLegendariesHarpyHave:       models.Harpy,
	fieldLegendariesReaperHave:      models.Reaper,
	fieldLegendariesCyclopsHave:     models.Cyclops,
	fieldLegendariesArchdemonHave:   models.Archdemon,
	fieldLegendariesTheCursedHave:   models.TheCursed,
	fieldLegendariesTheColossusHave: models.TheColossus,
	fieldLegendariesTheInfernalHave: models.TheInfernal,
	fieldLegendariesRoboChickenHave: models.RoboChicken,
	fieldLegendariesShieldBotHave:   models.ShieldBot,
	fieldLegendariesSoulStalkerHave: models.SoulStalker,

	fieldLegendariesLichPlan:        models.Lich,
	fieldLegendariesGorgonPlan:      models.Gorgon,
	fieldLegendariesHarpyPlan:       models.Harpy,
	fieldLegendariesReaperPlan:      models.Reaper,
	fieldLegendariesCyclopsPlan:     models.Cyclops,
	fieldLegendariesArchdemonPlan:   models.Archdemon,
	fieldLegendariesTheCursedPlan:   models.TheCursed,
	fieldLegendariesTheColossusPlan: models.TheColossus,
	fieldLegendariesTheInfernalPlan: models.TheInfernal,
	fieldLegendariesRoboChickenPlan: models.RoboChicken,
	fieldLegendariesShieldBotPlan:   models.ShieldBot,
	fieldLegendariesSoulStalkerPlan: models.SoulStalker,
}

func (m *Model) parseLegendariesTabFieldValues(i fieldIndex, value string) {
	legendary := legendaryIdByFieldIndex[i]
	if i >= fieldLegendariesLichHave && i <= fieldLegendariesSoulStalkerHave {
		if v, err := strconv.Atoi(value); err == nil {
			m.plan.PossessedLegendaries[legendary] = v
		}
		m.recalculate()
		return
	}
	if v, err := strconv.Atoi(value); err == nil {
		m.plan.LegendaryCounts[legendary] = v
	}
	m.recalculate()
}
