package prestigeplan

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/base"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/experiments"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/legendaries"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/runes"
)

type planTab int8

const (
	planTabBase planTab = iota
	planTabLegendaries
	planTabRunes
	planTabExperiments
)

type Model struct {
	selectedTab               planTab
	windowHeight, windowWidth int

	baseTab        *base.Model
	legendariesTab *legendaries.Model
	runesTab       *runes.Model
	experimentsTab *experiments.Model

	result *calculator.PrestigePlanResult
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) recalculate() {
	*m.result = calculator.Calculate(m.assemblePlan())
}

func (m *Model) assemblePlan() models.Plan {
	return models.Plan{
		DevourerLevel:        m.baseTab.DevourerLevel,
		FeatTiers:            m.baseTab.FeatTiers,
		OtherMultiplier:      m.baseTab.OtherMultiplier,
		GroupBonusCount:      m.baseTab.GroupBonusCount,
		LeftoverShards:       m.baseTab.LeftoverShards,
		LegendaryCounts:      m.legendariesTab.LegendaryCounts,
		PossessedLegendaries: m.legendariesTab.PossessedLegendaries,
		PossessedRunes:       m.runesTab.PossessedRunes,
		ExperimentLevels:     m.experimentsTab.ExperimentLevels,
	}
}

func New() *Model {
	resultPtr := new(calculator.PrestigePlanResult)
	m := &Model{
		selectedTab:    planTabExperiments,
		result:         resultPtr,
		baseTab:        base.NewModel(),
		legendariesTab: legendaries.NewModel(resultPtr),
		runesTab:       runes.NewModel(resultPtr),
		experimentsTab: experiments.NewModel(),
	}
	m.recalculate()
	return m
}

// currentInput returns the current input model for the currently selected tab
func (m *Model) currentInput() *textinput.Model {
	switch m.selectedTab {
	case planTabBase:
		return m.baseTab.CurrentInput()
	case planTabLegendaries:
		return m.legendariesTab.CurrentInput()
	case planTabRunes:
		return m.runesTab.CurrentInput()
	case planTabExperiments:
		return m.experimentsTab.CurrentInput()
	}
	return nil
}
