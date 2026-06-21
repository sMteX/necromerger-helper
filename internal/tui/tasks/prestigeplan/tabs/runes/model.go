package runes

import (
	"strconv"

	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	shared.TabModel

	// part of the `model.Plan` this tab changes
	PossessedRunes map[models.RuneType]int

	result *calculator.PrestigePlanResult
}

var runeByFieldType = map[fieldIndex]models.RuneType{
	fieldIce:    models.RuneIce,
	fieldPoison: models.RunePoison,
	fieldBlood:  models.RuneBlood,
	fieldMoon:   models.RuneMoon,
	fieldDeath:  models.RuneDeath,
	fieldCosmic: models.RuneCosmic,
}

func NewModel(resultPtr *calculator.PrestigePlanResult) *Model {
	m := &Model{
		TabModel: shared.NewTabModel(int(fieldIndexCount)),
		PossessedRunes: map[models.RuneType]int{
			models.RuneIce:    10,
			models.RunePoison: 100,
			models.RuneBlood:  1000,
			models.RuneMoon:   10000,
			models.RuneDeath:  20000,
			models.RuneCosmic: 10000,
		},
		result: resultPtr,
	}
	for i := fieldIce; i <= fieldCosmic; i++ {
		r := runeByFieldType[i]
		m.Fields[i] = shared.InputField{
			Label:          string(r),
			Step:           0,
			Width:          6,
			CharacterLimit: 6,
			Validate:       shared.InputValidationIntInRange(0, 10000000),
			InitialValue:   strconv.Itoa(m.PossessedRunes[r]),
		}
	}
	m.InitializeInputs()
	return m
}

type fieldIndex int8

const (
	fieldIce fieldIndex = iota
	fieldPoison
	fieldBlood
	fieldMoon
	fieldDeath
	fieldCosmic
	fieldIndexCount
)
