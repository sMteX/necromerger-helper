package base

import (
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	cursor int
	fields []shared.InputField

	// part of the `model.Plan` this tab changes
	DevourerLevel   int
	FeatTiers       int
	OtherMultiplier float64
	GroupBonusCount int
	LeftoverShards  int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func NewModel() *Model {
	m := &Model{
		fields: make([]shared.InputField, fieldIndexCount),
	}
	m.fields[fieldDevourerLevel] = shared.InputField{
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
	m.fields[fieldFeatTiers] = shared.InputField{
		Label:          "Max Feat Tier",
		Step:           1,
		Width:          7,
		CharacterLimit: 2,
		InitialValue:   strconv.Itoa(m.FeatTiers),
		Validate:       shared.InputValidationIntInRange(1, 35),
	}
	m.fields[fieldOtherMultiplier] = shared.InputField{
		Label:          "'Others' Multiplier [%]",
		Step:           0,
		Width:          8,
		CharacterLimit: 3,
		InitialValue:   strconv.Itoa(int(m.OtherMultiplier * 100)),
		Validate:       shared.InputValidationIntInRange(100, 1000),
	}
	m.fields[fieldGroupBonusCount] = shared.InputField{
		Label:          "Leg. Group Bonus Count",
		Step:           1,
		Width:          5,
		CharacterLimit: 1,
		InitialValue:   strconv.Itoa(m.GroupBonusCount),
		Validate:       shared.InputValidationIntInRange(1, 3),
	}
	m.fields[fieldLeftoverShards] = shared.InputField{
		Label:          "Leftover Shards",
		Step:           0,
		Width:          11,
		CharacterLimit: 7,
		InitialValue:   strconv.Itoa(m.LeftoverShards),
		Validate:       shared.InputValidationIntInRange(0, 10000000),
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
	fieldDevourerLevel fieldIndex = iota
	fieldFeatTiers
	fieldOtherMultiplier
	fieldGroupBonusCount
	fieldLeftoverShards
	fieldIndexCount
)
