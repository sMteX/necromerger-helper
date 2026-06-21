package experiments

import (
	"strconv"

	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

type Model struct {
	shared.TabModel

	// part of the `model.Plan` this tab changes
	ExperimentLevels map[models.ExperimentID]int
}

var experimentsByFieldIndex = map[fieldIndex]models.Experiment{
	fieldSeasoning1:   data.ExperimentSeasoning,
	fieldStrength1:    data.ExperimentStrength,
	fieldTaste1:       data.ExperimentTaste,
	fieldCapacity1:    data.ExperimentCapacity,
	fieldBodySnatcher: data.ExperimentBodySnatcher,
	fieldWeakening:    data.ExperimentWeakening,
	fieldDamageCap:    data.ExperimentDamageCap,
	fieldIceChest:     data.ExperimentIceChest,
	fieldPoisonChest:  data.ExperimentPoisonChest,
	fieldBloodChest:   data.ExperimentBloodChest,
	fieldMoonChest:    data.ExperimentMoonChest,
	fieldDeathChest:   data.ExperimentDeathChest,
	fieldCosmicChest:  data.ExperimentCosmicChest,
	fieldSeasoning2:   data.ExperimentSeasoning2,
	fieldStrength2:    data.ExperimentStrength2,
	fieldTaste2:       data.ExperimentTaste2,
	fieldCapacity2:    data.ExperimentCapacity2,
}

func NewModel() *Model {
	m := &Model{
		TabModel: shared.NewTabModel(int(fieldIndexCount)),
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
	}

	for i := fieldSeasoning1; i <= fieldCapacity2; i++ {
		e := experimentsByFieldIndex[i]
		maxLevel := e.Levels[len(e.Levels)-1].Level
		m.Fields[i] = shared.InputField{
			Step:           1,
			Width:          1,
			CharacterLimit: 1,
			InitialValue:   strconv.Itoa(m.ExperimentLevels[e.ID]),
			Validate:       shared.InputValidationIntInRange(0, maxLevel),
		}
	}

	m.InitializeInputs()
	return m
}

// LoadFrom replaces the tab's experiment levels with those from a loaded plan,
// updating both the model map and the textinput display values.
func (m *Model) LoadFrom(plan models.Plan) {
	m.ExperimentLevels = plan.ExperimentLevels
	for i := fieldSeasoning1; i <= fieldCapacity2; i++ {
		e := experimentsByFieldIndex[i]
		m.Fields[i].Input.SetValue(strconv.Itoa(plan.ExperimentLevels[e.ID]))
	}
}

type fieldIndex int8

const (
	fieldSeasoning1 fieldIndex = iota
	fieldStrength1
	fieldTaste1
	fieldCapacity1
	fieldBodySnatcher
	fieldWeakening
	fieldDamageCap
	fieldIceChest
	fieldPoisonChest
	fieldBloodChest
	fieldMoonChest
	fieldDeathChest
	fieldCosmicChest
	fieldSeasoning2
	fieldStrength2
	fieldTaste2
	fieldCapacity2
	fieldIndexCount
)
