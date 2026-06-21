package base

import (
	"math"
	"strconv"

	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	shared.TabModel

	// part of the `model.Plan` this tab changes
	DevourerLevel   int
	FeatTiers       int
	OtherMultiplier float64
	GroupBonusCount int
	LeftoverShards  int
}

func NewModel() *Model {
	m := &Model{
		TabModel:        shared.NewTabModel(int(fieldIndexCount)),
		DevourerLevel:   35,
		FeatTiers:       1,
		OtherMultiplier: 1,
		GroupBonusCount: 1,
	}
	m.Fields[fieldDevourerLevel] = shared.InputField{
		Label: "Devourer Level",
		Options: []string{"35", "40", "45", "50", "55", "60", "65", "70",
			"75", "80", "85", "90", "95", "100", "150",
			"200", "300", "400", "500", "600", "700",
			"800", "900", "1000"},
		Width:          8,
		CharacterLimit: 4,
		InitialValue:   strconv.Itoa(m.DevourerLevel),
		Validate:       shared.InputValidationIntInRange(1, 1000),
	}
	m.Fields[fieldFeatTiers] = shared.InputField{
		Label:          "Max Feat Tier",
		Step:           1,
		Width:          7,
		CharacterLimit: 2,
		InitialValue:   strconv.Itoa(m.FeatTiers),
		Validate:       shared.InputValidationIntInRange(1, 35),
	}
	m.Fields[fieldOtherMultiplier] = shared.InputField{
		Label:          "'Others' Multiplier [%]",
		Step:           0,
		Width:          8,
		CharacterLimit: 3,
		InitialValue:   strconv.Itoa(int(m.OtherMultiplier * 100)),
		Validate:       shared.InputValidationIntInRange(100, 1000),
	}
	m.Fields[fieldGroupBonusCount] = shared.InputField{
		Label:          "Leg. Group Bonus Count",
		Step:           1,
		Width:          5,
		CharacterLimit: 1,
		InitialValue:   strconv.Itoa(m.GroupBonusCount),
		Validate:       shared.InputValidationIntInRange(1, 3),
	}
	m.Fields[fieldLeftoverShards] = shared.InputField{
		Label:          "Leftover Shards",
		Step:           0,
		Width:          11,
		CharacterLimit: 7,
		InitialValue:   strconv.Itoa(m.LeftoverShards),
		Validate:       shared.InputValidationIntInRange(0, 10000000),
	}
	m.InitializeInputs()
	return m
}

// LoadFrom replaces the tab's values with those from a loaded plan.
// Each field is updated in two places: the model field (used by the calculator)
// and the textinput display value (what the user sees on screen).
func (m *Model) LoadFrom(plan models.Plan) {
	m.DevourerLevel = plan.DevourerLevel
	m.FeatTiers = plan.FeatTiers
	m.OtherMultiplier = plan.OtherMultiplier
	m.GroupBonusCount = plan.GroupBonusCount
	m.LeftoverShards = plan.LeftoverShards
	m.Fields[fieldDevourerLevel].Input.SetValue(strconv.Itoa(plan.DevourerLevel))
	m.Fields[fieldFeatTiers].Input.SetValue(strconv.Itoa(plan.FeatTiers))
	// OtherMultiplier is stored as a decimal (e.g. 1.72) but the input shows
	// it as an integer percentage (e.g. 172). math.Round avoids float precision
	// issues like 1.72 * 100 = 171.99999...
	m.Fields[fieldOtherMultiplier].Input.SetValue(strconv.Itoa(int(math.Round(plan.OtherMultiplier * 100))))
	m.Fields[fieldGroupBonusCount].Input.SetValue(strconv.Itoa(plan.GroupBonusCount))
	m.Fields[fieldLeftoverShards].Input.SetValue(strconv.Itoa(plan.LeftoverShards))
}

type fieldIndex int8

const (
	fieldDevourerLevel fieldIndex = iota
	fieldFeatTiers
	fieldOtherMultiplier
	fieldGroupBonusCount
	fieldLeftoverShards
	fieldIndexCount
)
