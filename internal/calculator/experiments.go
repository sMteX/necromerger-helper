package calculator

import (
	"fmt"

	"github.com/sMteX/necro-prestige-planner/internal/models"
)

var Experiments = []models.Experiment{
	{
		ID: models.ExpSeasoning, Name: "Seasoning Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	{
		ID: models.ExpStrength, Name: "Strength Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	{
		ID: models.ExpTaste, Name: "Taste Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 0.20, PrevValue: 0.0},
			{Level: 2, Cost: 500, Value: 0.40, PrevValue: 0.20},
			{Level: 3, Cost: 1000, Value: 0.60, PrevValue: 0.40},
			{Level: 4, Cost: 2000, Value: 0.80, PrevValue: 0.60},
			{Level: 5, Cost: 4000, Value: 1.00, PrevValue: 0.80},
			{Level: 6, Cost: 20000, Value: 1.25, PrevValue: 1.00},
			{Level: 7, Cost: 50000, Value: 1.50, PrevValue: 1.25},
			{Level: 8, Cost: 100000, Value: 2.00, PrevValue: 1.50},
			{Level: 9, Cost: 500000, Value: 3.00, PrevValue: 2.00},
		},
	},
	{
		ID: models.ExpCapacity, Name: "Capacity Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100, Value: 0.05, PrevValue: 0.0},
			{Level: 2, Cost: 750, Value: 0.10, PrevValue: 0.05},
			{Level: 3, Cost: 1500, Value: 0.15, PrevValue: 0.10},
			{Level: 4, Cost: 2500, Value: 0.20, PrevValue: 0.15},
			{Level: 5, Cost: 5000, Value: 0.25, PrevValue: 0.20},
			{Level: 6, Cost: 25000, Value: 0.30, PrevValue: 0.25},
			{Level: 7, Cost: 75000, Value: 0.35, PrevValue: 0.30},
			{Level: 8, Cost: 150000, Value: 0.40, PrevValue: 0.35},
			{Level: 9, Cost: 750000, Value: 0.50, PrevValue: 0.40},
		},
	},
	{
		ID: models.ExpBodySnatcher, Name: "Body Snatcher", Tier: models.TierPre100, IsSpecial: true,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1.0, PrevValue: 1.0},
		},
	},
	{
		ID: models.ExpWeakening, Name: "Weakening Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 750, Value: 2.50, PrevValue: 5.00},
			{Level: 2, Cost: 15000, Value: 1.75, PrevValue: 2.50},
			{Level: 3, Cost: 25000, Value: 1.50, PrevValue: 1.75},
			{Level: 4, Cost: 50000, Value: 1.40, PrevValue: 1.50},
			{Level: 5, Cost: 500000, Value: 1.35, PrevValue: 1.40},
			{Level: 6, Cost: 5000000, Value: 1.30, PrevValue: 1.35},
			{Level: 7, Cost: 25000000, Value: 1.27, PrevValue: 1.30},
			{Level: 8, Cost: 100000000, Value: 1.25, PrevValue: 1.27},
		},
	},
	{
		ID: models.ExpDamageCap, Name: "Damage Cap Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25000, Value: 0.65, PrevValue: 0.50},
			{Level: 2, Cost: 1000000, Value: 0.75, PrevValue: 0.65},
			{Level: 3, Cost: 50000000, Value: 0.80, PrevValue: 0.75},
		},
	},
	{
		ID: models.ExpIceChest, Name: "Ice Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 10, Value: 1, PrevValue: 0}, {Level: 2, Cost: 500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 1500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 5000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 10000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 50000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 100000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpPoisonChest, Name: "Poison Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 1, PrevValue: 0}, {Level: 2, Cost: 1000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 3000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 7500, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 12500, Value: 5, PrevValue: 4}, {Level: 6, Cost: 75000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 150000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1500000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpBloodChest, Name: "Blood Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1, PrevValue: 0}, {Level: 2, Cost: 2500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 5000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 10000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 50000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 100000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 250000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 750000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 2500000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpMoonChest, Name: "Moon Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 250, Value: 1, PrevValue: 0}, {Level: 2, Cost: 3000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 7500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 15000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 75000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 150000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 500000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 1000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 5000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpDeathChest, Name: "Death Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 500, Value: 1, PrevValue: 0}, {Level: 2, Cost: 5000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 10000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 50000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 150000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 500000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 1000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 5000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 10000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpCosmicChest, Name: "Cosmic Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 5000, Value: 1, PrevValue: 0}, {Level: 2, Cost: 10000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 15000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 75000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 250000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 1000000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 5000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 10000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 25000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpSeasoning2, Name: "Seasoning Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 5, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 10, PrevValue: 5},
			{Level: 3, Cost: 1000000, Value: 15, PrevValue: 10},
			{Level: 4, Cost: 2000000, Value: 20, PrevValue: 15},
			{Level: 5, Cost: 3000000, Value: 25, PrevValue: 20},
			{Level: 6, Cost: 5000000, Value: 30, PrevValue: 25},
			{Level: 7, Cost: 10000000, Value: 35, PrevValue: 30},
			{Level: 8, Cost: 20000000, Value: 40, PrevValue: 35},
			{Level: 9, Cost: 50000000, Value: 50, PrevValue: 40},
		},
	},
	{
		ID: models.ExpStrength2, Name: "Strength Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 1.2, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 1.4, PrevValue: 1.2},
			{Level: 3, Cost: 1000000, Value: 1.6, PrevValue: 1.4},
			{Level: 4, Cost: 2000000, Value: 1.8, PrevValue: 1.6},
			{Level: 5, Cost: 3000000, Value: 2.0, PrevValue: 1.8},
			{Level: 6, Cost: 5000000, Value: 2.2, PrevValue: 2.0},
			{Level: 7, Cost: 10000000, Value: 2.4, PrevValue: 2.2},
			{Level: 8, Cost: 20000000, Value: 2.6, PrevValue: 2.4},
			{Level: 9, Cost: 50000000, Value: 3.0, PrevValue: 2.6},
		},
	},
	{
		ID: models.ExpTaste2, Name: "Taste Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 2, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 3, PrevValue: 2},
			{Level: 3, Cost: 3000000, Value: 4, PrevValue: 3},
			{Level: 4, Cost: 5000000, Value: 5, PrevValue: 4},
			{Level: 5, Cost: 10000000, Value: 6, PrevValue: 5},
			{Level: 6, Cost: 20000000, Value: 7, PrevValue: 6},
			{Level: 7, Cost: 30000000, Value: 8, PrevValue: 7},
			{Level: 8, Cost: 50000000, Value: 9, PrevValue: 8},
			{Level: 9, Cost: 100000000, Value: 10, PrevValue: 9},
		},
	},
	{
		ID: models.ExpCapacity2, Name: "Capacity Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 1.1, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 1.2, PrevValue: 1.1},
			{Level: 3, Cost: 3000000, Value: 1.3, PrevValue: 1.2},
			{Level: 4, Cost: 5000000, Value: 1.4, PrevValue: 1.3},
			{Level: 5, Cost: 10000000, Value: 1.5, PrevValue: 1.4},
			{Level: 6, Cost: 20000000, Value: 1.6, PrevValue: 1.5},
			{Level: 7, Cost: 30000000, Value: 1.8, PrevValue: 1.6},
			{Level: 8, Cost: 50000000, Value: 1.9, PrevValue: 1.8},
			{Level: 9, Cost: 100000000, Value: 2, PrevValue: 1.9},
		},
	},
}

func CalculateExperimentCost(plan models.Plan) int {
	total := 0
	for _, exp := range Experiments {
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
	for _, e := range Experiments {
		if e.Tier == models.TierPre100 {
			res = append(res, e)
		}
	}
	return res
}

func GetPost100Experiments() []models.Experiment {
	var res []models.Experiment
	for _, e := range Experiments {
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
