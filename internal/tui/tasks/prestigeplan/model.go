package prestigeplan

import (
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
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

	cursor int
	fields []inputField

	plan   models.Plan
	result calculator.PrestigePlanResult
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
		cursor:      int(fieldExperimentsSeasoning1),
		plan: models.Plan{
			DevourerLevel:   200,
			FeatTiers:       27,
			OtherMultiplier: 1.72,
			GroupBonusCount: 1,
			LeftoverShards:  123456,
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
	m.fields = make([]inputField, fieldIndexCount)
	m.addBaseTabFields()
	m.addLegendariesTabFields()
	m.addRunesTabFields()
	m.addExperimentsTabFields()
	m.initializeInputModels()
	m.recalculate()
	return m
}

func (m *Model) initializeInputModels() {
	for i, field := range m.fields {
		f := textinput.New()
		f.Prompt = ""
		f.CharLimit = field.characterLimit
		f.Validate = field.validate
		f.SetVirtualCursor(true)
		f.SetWidth(field.width)
		f.SetValue(field.initialValue)
		m.fields[i].input = f
	}
}

func (m *Model) currentInput() *textinput.Model {
	return &m.fields[m.cursor].input
}

type inputField struct {
	label          string   // not going to be used every time, but it's convenient to have it here
	step           int      // step == 0 means the field is text-only; step > 0 enables ←/→ increment/decrement.
	options        []string // if len(options) > 0, step is ignored and the field is a select-like
	width          int
	characterLimit int
	validate       func(string) error
	initialValue   string
	input          textinput.Model
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

func inputValidationIntInRange(min, max int) func(string) error {
	return func(s string) error {
		v, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("must be a whole number")
		}
		if v < min || v > max {
			return fmt.Errorf("must be between %d and %d", min, max)
		}
		return nil
	}
}
