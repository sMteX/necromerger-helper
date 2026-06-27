package stationefficiency

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necromerger-helper/internal/calculator"
	"github.com/sMteX/necromerger-helper/internal/models"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

type Model struct {
	shared.TabModel
	windowHeight, windowWidth int

	selectedChampion     models.ChampionID
	manaCap              int
	slimeCap             int
	darknessCap          int
	championSpeedPercent int // to make it easier to fill in directly from the info screen - 185 (%) = 1.85x multiplier

	result *calculator.ChampionEfficiencyResult
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) recalculate() {
	*m.result = calculator.CalculateChampionEfficiency(calculator.ChampionEfficiencyInput{
		ChampionID:   m.selectedChampion,
		ManaCap:      m.manaCap,
		SlimeCap:     m.slimeCap,
		DarknessCap:  m.darknessCap,
		SpeedPercent: m.championSpeedPercent,
	})
}

func New() *Model {
	resultPtr := new(calculator.ChampionEfficiencyResult)
	m := &Model{
		selectedChampion:     models.ChampionPeasant,
		championSpeedPercent: 100,
		result:               resultPtr,
		TabModel:             shared.NewTabModel(int(fieldIndexCount)),
	}
	m.Fields[fieldChampion] = shared.InputField{
		Label: "Champion",
		Options: []string{
			string(models.ChampionPeasant),
			string(models.ChampionKnight),
			string(models.ChampionCleric),
			string(models.ChampionPaladin),
			string(models.ChampionRival),
		},
		Step:           0,
		Width:          7,
		CharacterLimit: 7,
		InitialValue:   string(m.selectedChampion),
	}
	m.Fields[fieldManaCap] = shared.InputField{
		Label:          "Mana cap",
		Step:           0,
		Width:          7,
		CharacterLimit: 7,
		InitialValue:   strconv.Itoa(m.manaCap),
		Validate:       shared.InputValidationIntInRange(1, 9_999_999),
	}
	m.Fields[fieldSlimeCap] = shared.InputField{
		Label:          "Slime cap",
		Step:           0,
		Width:          7,
		CharacterLimit: 7,
		InitialValue:   strconv.Itoa(m.slimeCap),
		Validate:       shared.InputValidationIntInRange(1, 9_999_999),
	}
	m.Fields[fieldDarknessCap] = shared.InputField{
		Label:          "Darkness cap",
		Step:           0,
		Width:          7,
		CharacterLimit: 7,
		InitialValue:   strconv.Itoa(m.darknessCap),
		Validate:       shared.InputValidationIntInRange(1, 9_999_999),
	}
	m.Fields[fieldSpeedPercent] = shared.InputField{
		Label:          "Champion Speed %",
		Step:           0,
		Width:          3,
		CharacterLimit: 3,
		InitialValue:   strconv.Itoa(m.championSpeedPercent),
		Validate:       shared.InputValidationIntInRange(100, 999),
	}
	m.InitializeInputs()
	m.recalculate()
	return m
}

type fieldIndex int8

const (
	fieldChampion fieldIndex = iota
	fieldManaCap
	fieldSlimeCap
	fieldDarknessCap
	fieldSpeedPercent
	fieldIndexCount
)
