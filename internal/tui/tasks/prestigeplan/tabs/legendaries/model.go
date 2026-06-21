package legendaries

import (
	"strconv"

	"github.com/sMteX/necromerger-helper/internal/calculator"
	"github.com/sMteX/necromerger-helper/internal/data"
	"github.com/sMteX/necromerger-helper/internal/models"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

// LoadFrom replaces the tab's values with those from a loaded plan.
// The field indices are split into two ranges: Have fields (what the player
// currently owns) and Plan fields (target counts). Both the model maps and the
// textinput display values are updated.
func (m *Model) LoadFrom(plan models.Plan) {
	m.LegendaryCounts = plan.LegendaryCounts
	m.PossessedLegendaries = plan.PossessedLegendaries
	for i := fieldLichHave; i <= fieldSoulStalkerHave; i++ {
		id := legendaryIdByFieldIndex[i]
		m.Fields[i].Input.SetValue(strconv.Itoa(plan.PossessedLegendaries[id]))
	}
	for i := fieldLichPlan; i <= fieldSoulStalkerPlan; i++ {
		id := legendaryIdByFieldIndex[i]
		m.Fields[i].Input.SetValue(strconv.Itoa(plan.LegendaryCounts[id]))
	}
}

type Model struct {
	shared.TabModel

	// part of the `model.Plan` this tab changes
	LegendaryCounts      map[models.LegendaryID]int
	PossessedLegendaries map[models.LegendaryID]int

	result *calculator.PrestigePlanResult
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

func NewModel(resultPtr *calculator.PrestigePlanResult) *Model {
	m := &Model{
		TabModel:             shared.NewTabModel(int(fieldIndexCount)),
		LegendaryCounts:      make(map[models.LegendaryID]int),
		PossessedLegendaries: make(map[models.LegendaryID]int),
		result:               resultPtr,
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
		m.Fields[i] = shared.InputField{
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
		m.Fields[i] = shared.InputField{
			Label:          legendary.Name,
			Step:           1,
			Width:          5,
			CharacterLimit: 2,
			InitialValue:   strconv.Itoa(m.LegendaryCounts[legendary.ID]),
			Validate:       shared.InputValidationIntInRange(limits[0], limits[1]),
		}
	}

	m.InitializeInputs()
	return m
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
