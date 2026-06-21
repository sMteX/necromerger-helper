package prestigeplan

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/base"
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

	baseTab *base.Model
	plan    models.Plan
	result  calculator.PrestigePlanResult
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) recalculate() {
	m.result = calculator.Calculate(m.plan)
}

func New() *Model {
	m := &Model{
		selectedTab: planTabExperiments,
		baseTab:     base.NewModel(),
		plan: models.Plan{
			LegendaryCounts: map[models.LegendaryID]int{
				models.Lich:        11,
				models.Gorgon:      10,
				models.Harpy:       1,
				models.Reaper:      1,
				models.Cyclops:     1,
				models.Archdemon:   4,
				models.TheCursed:   1,
				models.TheColossus: 1,
				models.TheInfernal: 1,
				models.RoboChicken: 4,
				models.ShieldBot:   4,
				models.SoulStalker: 15,
			},
			ExperimentLevels: map[models.ExperimentID]int{
				models.ExpSeasoning:    6,
				models.ExpStrength:     2,
				models.ExpTaste:        6,
				models.ExpCapacity:     6,
				models.ExpBodySnatcher: 1,
				models.ExpWeakening:    4,
				models.ExpDamageCap:    2,
				models.ExpIceChest:     6,
				models.ExpPoisonChest:  7,
				models.ExpBloodChest:   6,
				models.ExpMoonChest:    6,
				models.ExpDeathChest:   6,
				models.ExpCosmicChest:  5,
				models.ExpSeasoning2:   2,
				models.ExpStrength2:    2,
				models.ExpTaste2:       2,
				models.ExpCapacity2:    1,
			},
			PossessedRunes: map[models.RuneType]int{
				models.RuneIce:    10,
				models.RunePoison: 100,
				models.RuneBlood:  1000,
				models.RuneMoon:   10000,
				models.RuneDeath:  20000,
				models.RuneCosmic: 10000,
			},
			PossessedLegendaries: map[models.LegendaryID]int{
				models.Lich:        11,
				models.Gorgon:      0,
				models.Harpy:       0,
				models.Reaper:      0,
				models.Cyclops:     0,
				models.Archdemon:   4,
				models.TheCursed:   1,
				models.TheColossus: 1,
				models.TheInfernal: 1,
				models.RoboChicken: 4,
				models.ShieldBot:   4,
				models.SoulStalker: 4,
			},
		},
	}
	m.addLegendariesTabFields()
	m.addRunesTabFields()
	m.addExperimentsTabFields()
	m.recalculate()
	return m
}

// currentInput returns the current input model for the currently selected tab
func (m *Model) currentInput() *textinput.Model {
	switch m.selectedTab {
	case planTabBase:
		return m.baseTab.CurrentInput()
	}
	return nil
}

type fieldIndex int8

const (
	// Base fields
	fieldBaseDevourerLevel fieldIndex = iota
	fieldBaseFeatTiers
	fieldBaseOtherMultiplier
	fieldBaseGroupBonusCount
	fieldBaseLeftoverShards
	// Legendaries fields
	fieldLegendariesLichHave
	fieldLegendariesGorgonHave
	fieldLegendariesHarpyHave
	fieldLegendariesReaperHave
	fieldLegendariesCyclopsHave
	fieldLegendariesArchdemonHave
	fieldLegendariesTheCursedHave
	fieldLegendariesTheColossusHave
	fieldLegendariesTheInfernalHave
	fieldLegendariesRoboChickenHave
	fieldLegendariesShieldBotHave
	fieldLegendariesSoulStalkerHave
	fieldLegendariesLichPlan
	fieldLegendariesGorgonPlan
	fieldLegendariesHarpyPlan
	fieldLegendariesReaperPlan
	fieldLegendariesCyclopsPlan
	fieldLegendariesArchdemonPlan
	fieldLegendariesTheCursedPlan
	fieldLegendariesTheColossusPlan
	fieldLegendariesTheInfernalPlan
	fieldLegendariesRoboChickenPlan
	fieldLegendariesShieldBotPlan
	fieldLegendariesSoulStalkerPlan
	// Runes fields
	fieldRunesIce
	fieldRunesPoison
	fieldRunesBlood
	fieldRunesMoon
	fieldRunesDeath
	fieldRunesCosmic
	// Experiments fields
	fieldExperimentsSeasoning1
	fieldExperimentsStrength1
	fieldExperimentsTaste1
	fieldExperimentsCapacity1
	fieldExperimentsBodySnatcher
	fieldExperimentsWeakening
	fieldExperimentsDamageCap
	fieldExperimentsIceChest
	fieldExperimentsPoisonChest
	fieldExperimentsBloodChest
	fieldExperimentsMoonChest
	fieldExperimentsDeathChest
	fieldExperimentsCosmicChest
	fieldExperimentsSeasoning2
	fieldExperimentsStrength2
	fieldExperimentsTaste2
	fieldExperimentsCapacity2
	// special constant that automatically updates and refers to the amount of these constants (like `len(fieldIndex)`)
	fieldIndexCount
)
