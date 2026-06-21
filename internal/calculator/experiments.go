package calculator

import (
	"github.com/sMteX/necromerger-helper/internal/data"
	"github.com/sMteX/necromerger-helper/internal/models"
)

func CalculateExperimentCost(plan models.Plan) int {
	total := 0
	for _, exp := range data.Experiments {
		level, ok := plan.ExperimentLevels[exp.ID]
		if !ok || level <= 0 {
			continue
		}

		for _, expLevel := range exp.Levels {
			if expLevel.Level <= level {
				total += expLevel.Cost
			}
		}
	}
	return total
}

func GetPre100Experiments() []models.Experiment {
	var res []models.Experiment
	for _, e := range data.Experiments {
		if e.Tier == models.TierPre100 {
			res = append(res, e)
		}
	}
	return res
}

func GetPost100Experiments() []models.Experiment {
	var res []models.Experiment
	for _, e := range data.Experiments {
		if e.Tier == models.TierPost100 {
			res = append(res, e)
		}
	}
	return res
}
