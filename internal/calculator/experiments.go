package calculator

import (
	"fmt"

	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
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

func FormatExperimentValue(id models.ExperimentID, tier models.ExperimentTier, value float64) string {
	switch id {
	case models.ExpSeasoning, models.ExpStrength, models.ExpTaste, models.ExpCapacity, models.ExpDamageCap:
		return fmt.Sprintf("%d%%", int(value*100))
	case models.ExpWeakening:
		return fmt.Sprintf("%.2f", value)
	case models.ExpIceChest, models.ExpPoisonChest, models.ExpBloodChest, models.ExpMoonChest, models.ExpDeathChest, models.ExpCosmicChest:
		return fmt.Sprintf("+%d", int(value))
	case models.ExpBodySnatcher:
		if value > 0 {
			return "Yes"
		}
		return "No"
	}

	if tier == models.TierPost100 {
		return fmt.Sprintf("x%.1f", value)
	}

	return fmt.Sprintf("%.1f", value)
}
