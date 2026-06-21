package legendaries

import (
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	cursor int
	fields []shared.InputField

	// part of the `model.Plan` this tab changes
	LegendaryCounts      map[models.LegendaryID]int
	PossessedLegendaries map[models.LegendaryID]int
	// ensure these are updated too
	LegendaryBonuses      map[models.LegendaryID]float64
	LegendaryGroupBonuses map[models.LegendaryGroup]float64
}

func (m *Model) Init() tea.Cmd {
	return nil
}

var legendaryIdByFieldIndex = map[fieldIndex]models.LegendaryID{
	fieldLichHave:        models.Lich,
	fieldGorgonHave:      models.Gorgon,
	fieldHarpyHave:       models.Harpy,
	fieldReaperHave:      models.Reaper,
	fieldCyclopsHave:     models.Cyclops,
	fieldArchdemonHave:   models.Archdemon,
	fieldTheCursedHave:   models.TheCursed,
	fieldTheColossusHave: models.TheColossus,
	fieldTheInfernalHave: models.TheInfernal,
	fieldRoboChickenHave: models.RoboChicken,
	fieldShieldBotHave:   models.ShieldBot,
	fieldSoulStalkerHave: models.SoulStalker,

	fieldLichPlan:        models.Lich,
	fieldGorgonPlan:      models.Gorgon,
	fieldHarpyPlan:       models.Harpy,
	fieldReaperPlan:      models.Reaper,
	fieldCyclopsPlan:     models.Cyclops,
	fieldArchdemonPlan:   models.Archdemon,
	fieldTheCursedPlan:   models.TheCursed,
	fieldTheColossusPlan: models.TheColossus,
	fieldTheInfernalPlan: models.TheInfernal,
	fieldRoboChickenPlan: models.RoboChicken,
	fieldShieldBotPlan:   models.ShieldBot,
	fieldSoulStalkerPlan: models.SoulStalker,
}

func NewModel() *Model {
	m := &Model{
		fields: make([]shared.InputField, fieldIndexCount),
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
	}

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
	for i := fieldLichHave; i <= fieldSoulStalkerHave; i++ {
		legendary := data.LegendariesById[legendaryIdByFieldIndex[i]]
		limits := legendaryCountLimits[legendary.ID]
		m.fields[i] = shared.InputField{
			Label:          legendary.Name,
			Step:           1,
			Width:          5,
			CharacterLimit: 2,
			InitialValue:   strconv.Itoa(m.PossessedLegendaries[legendary.ID]),
			Validate:       shared.InputValidationIntInRange(limits[0], limits[1]),
		}
	}
	for i := fieldLichPlan; i <= fieldSoulStalkerPlan; i++ {
		legendary := data.LegendariesById[legendaryIdByFieldIndex[i]]
		limits := legendaryCountLimits[legendary.ID]
		m.fields[i] = shared.InputField{
			Label:          legendary.Name,
			Step:           1,
			Width:          5,
			CharacterLimit: 2,
			InitialValue:   strconv.Itoa(m.LegendaryCounts[legendary.ID]),
			Validate:       shared.InputValidationIntInRange(limits[0], limits[1]),
		}
	}

	for i, field := range m.fields {
		m.fields[i].Input = field.CreateInput()
	}
	return m
}

func (m *Model) currentField() *shared.InputField {
	return &m.fields[m.cursor]
}

func (m *Model) CurrentInput() *textinput.Model {
	return &m.currentField().Input
}

type fieldIndex int8

const (
	fieldLichHave fieldIndex = iota
	fieldGorgonHave
	fieldHarpyHave
	fieldReaperHave
	fieldCyclopsHave
	fieldArchdemonHave
	fieldTheCursedHave
	fieldTheColossusHave
	fieldTheInfernalHave
	fieldRoboChickenHave
	fieldShieldBotHave
	fieldSoulStalkerHave
	fieldLichPlan
	fieldGorgonPlan
	fieldHarpyPlan
	fieldReaperPlan
	fieldCyclopsPlan
	fieldArchdemonPlan
	fieldTheCursedPlan
	fieldTheColossusPlan
	fieldTheInfernalPlan
	fieldRoboChickenPlan
	fieldShieldBotPlan
	fieldSoulStalkerPlan
	fieldIndexCount
)
