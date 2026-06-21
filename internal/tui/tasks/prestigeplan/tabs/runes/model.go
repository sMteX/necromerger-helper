package runes

import (
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	cursor int
	fields []shared.InputField

	// part of the `model.Plan` this tab changes
	PossessedRunes map[models.RuneType]int
	// ensure these are updated too
	RuneTotal  models.RuneCosts
	RuneNeeded models.RuneCosts
}

func (m *Model) Init() tea.Cmd {
	return nil
}

var runeByFieldType = map[fieldIndex]models.RuneType{
	fieldIce:    models.RuneIce,
	fieldPoison: models.RunePoison,
	fieldBlood:  models.RuneBlood,
	fieldMoon:   models.RuneMoon,
	fieldDeath:  models.RuneDeath,
	fieldCosmic: models.RuneCosmic,
}

func NewModel() *Model {
	m := &Model{
		fields: make([]shared.InputField, fieldIndexCount),
		PossessedRunes: map[models.RuneType]int{
			models.RuneIce:    10,
			models.RunePoison: 100,
			models.RuneBlood:  1000,
			models.RuneMoon:   10000,
			models.RuneDeath:  20000,
			models.RuneCosmic: 10000,
		},
	}
	for i := fieldIce; i <= fieldCosmic; i++ {
		r := runeByFieldType[i]
		m.fields[i] = shared.InputField{
			Label:          string(r),
			Step:           0,
			Width:          6,
			CharacterLimit: 6,
			Validate:       shared.InputValidationIntInRange(0, 10000000),
			InitialValue:   strconv.Itoa(m.PossessedRunes[r]),
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
	fieldIce fieldIndex = iota
	fieldPoison
	fieldBlood
	fieldMoon
	fieldDeath
	fieldCosmic
	fieldIndexCount
)
