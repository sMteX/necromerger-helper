package prestigePlan

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

var experimentsByFieldIndex = map[fieldIndex]models.Experiment{
	fieldExperimentsSeasoning1:   data.ExperimentSeasoning,
	fieldExperimentsStrength1:    data.ExperimentStrength,
	fieldExperimentsTaste1:       data.ExperimentTaste,
	fieldExperimentsCapacity1:    data.ExperimentCapacity,
	fieldExperimentsBodySnatcher: data.ExperimentBodySnatcher,
	fieldExperimentsWeakening:    data.ExperimentWeakening,
	fieldExperimentsDamageCap:    data.ExperimentDamageCap,
	fieldExperimentsIceChest:     data.ExperimentIceChest,
	fieldExperimentsPoisonChest:  data.ExperimentPoisonChest,
	fieldExperimentsBloodChest:   data.ExperimentBloodChest,
	fieldExperimentsMoonChest:    data.ExperimentMoonChest,
	fieldExperimentsDeathChest:   data.ExperimentDeathChest,
	fieldExperimentsCosmicChest:  data.ExperimentCosmicChest,
	fieldExperimentsSeasoning2:   data.ExperimentSeasoning2,
	fieldExperimentsStrength2:    data.ExperimentStrength2,
	fieldExperimentsTaste2:       data.ExperimentTaste2,
	fieldExperimentsCapacity2:    data.ExperimentCapacity2,
}

func (m *Model) addExperimentsTabFields() {
	for i := fieldExperimentsSeasoning1; i <= fieldExperimentsCapacity2; i++ {
		e := experimentsByFieldIndex[i]
		maxLevel := e.Levels[len(e.Levels)-1].Level
		m.fields[i] = inputField{
			step:           1,
			width:          1,
			characterLimit: 1,
			initialValue:   strconv.Itoa(m.plannedExperiments[e.ID]),
			validate:       inputValidationIntInRange(0, maxLevel),
		}
	}
}

var (
	experimentsNameColumn        = lipgloss.NewStyle().Width(20)
	experimentsLevelColumn       = lipgloss.NewStyle().Width(9).AlignHorizontal(lipgloss.Center)
	experimentsEffectColumn      = lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center)
	experimentsCurrentCostColumn = lipgloss.NewStyle().Width(6).AlignHorizontal(lipgloss.Right)
	experimentsNextCostColumn    = lipgloss.NewStyle().Width(14).MarginLeft(3)
)

func (m *Model) renderExperimentsTab() string {
	tableWidth := experimentsNameColumn.GetWidth() + experimentsLevelColumn.GetWidth() + experimentsEffectColumn.GetWidth() + experimentsCurrentCostColumn.GetWidth() + experimentsNextCostColumn.GetWidth() + experimentsNextCostColumn.GetHorizontalFrameSize()
	//arrow := " → "

	lines := []string{
		lipgloss.NewStyle().Bold(true).Render(
			experimentsNameColumn.Render("Experiment") + experimentsLevelColumn.Render("Level") + experimentsEffectColumn.Render("Effect") + experimentsCurrentCostColumn.Render("Cost") + experimentsNextCostColumn.Render("Next level"),
		),
		renderExperimentHeadingText("Pre-100", tableWidth),
		m.renderExperimentRow(fieldExperimentsSeasoning1),
		m.renderExperimentRow(fieldExperimentsStrength1),
		m.renderExperimentRow(fieldExperimentsTaste1),
		m.renderExperimentRow(fieldExperimentsCapacity1),
		m.renderExperimentRow(fieldExperimentsBodySnatcher),
		m.renderExperimentRow(fieldExperimentsWeakening),
		m.renderExperimentRow(fieldExperimentsDamageCap),
		m.renderExperimentRow(fieldExperimentsIceChest),
		m.renderExperimentRow(fieldExperimentsPoisonChest),
		m.renderExperimentRow(fieldExperimentsBloodChest),
		m.renderExperimentRow(fieldExperimentsMoonChest),
		m.renderExperimentRow(fieldExperimentsDeathChest),
		m.renderExperimentRow(fieldExperimentsCosmicChest),
		renderExperimentHeadingText("Post-100", tableWidth),
		m.renderExperimentRow(fieldExperimentsSeasoning2),
		m.renderExperimentRow(fieldExperimentsStrength2),
		m.renderExperimentRow(fieldExperimentsTaste2),
		m.renderExperimentRow(fieldExperimentsCapacity2),
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

func (m *Model) renderExperimentRow(i fieldIndex) string {
	e := experimentsByFieldIndex[i]
	// TODO: careful with indices, `plannedExperiments[...]` can be 0 - not planned
	var currentLevel *models.ExperimentLevel
	if m.plannedExperiments[e.ID] > 0 {
		// let's assume the `plannedExperiments[]` is either 0 (not planned) or in bounds (after subtracting 1)
		currentLevel = &e.Levels[m.plannedExperiments[e.ID]-1]
	}
	var nextLevel *models.ExperimentLevel
	if m.plannedExperiments[e.ID] < len(e.Levels) {
		nextLevel = &e.Levels[m.plannedExperiments[e.ID]]
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
		experimentsNameColumn.Render(e.Name),
		m.renderExperimentsRowLevelInput(i),
		experimentsEffectColumn.Render(m.getEffectText(e.ID, e.Tier, currentLevel)),
		experimentsCurrentCostColumn.Render(currentCost),
		experimentsNextCostColumn.Foreground(shared.Colors.Dim).Render(nextCost),
	)
}

func (m *Model) renderExperimentsRowLevelInput(i fieldIndex) string {
	e := experimentsByFieldIndex[i]
	if m.cursor == int(i) {
		current := m.currentInput().Value()
		valueText := "none"
		if current != "0" {
			valueText = fmt.Sprintf("lvl %s", current)
		}
		return experimentsLevelColumn.Foreground(shared.Colors.Good).Render(fmt.Sprintf("< %s >", valueText))
	}
	return experimentsLevelColumn.Render(fmt.Sprintf("lvl %d", m.plannedExperiments[e.ID]))
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

func (m *Model) getExperimentsTabHelp() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("← / →  change level"),
		shared.Styles.Help.Render("F1 - F4  switch tab"),
		shared.Styles.Help.Render("q / ctrl+c  exit"),
	}
}

func (m *Model) handleExperimentsTabKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > int(fieldExperimentsSeasoning1) {
			m.currentInput().Blur()
			m.cursor--
			return m, m.currentInput().Focus()
		}
		return m, nil
	case "down":
		if m.cursor < int(fieldExperimentsCapacity2) {
			m.currentInput().Blur()
			m.cursor++
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
			m.parseExperimentsTabFieldValues(fieldIndex(m.cursor), newVal)
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
		m.parseExperimentsTabFieldValues(fieldIndex(m.cursor), m.currentInput().Value())
		// TODO: recalculate m.calculatedOutputs from m.baseInputs
	}
	return m, cmd
}

func (m *Model) parseExperimentsTabFieldValues(i fieldIndex, value string) {
	e := experimentsByFieldIndex[i]
	if v, err := strconv.Atoi(value); err == nil {
		m.plannedExperiments[e.ID] = v
	}
}
