package calculator

import (
	"testing"

	"github.com/sMteX/necro-prestige-planner/internal/models"
)

func TestCalculateTimeShards(t *testing.T) {
	plan := models.Plan{
		DevourerLevel:   50,
		FeatTiers:       0,
		OtherMultiplier: 0,
		GroupBonusCount: 1,
		LegendaryCounts: map[models.LegendaryID]int{
			models.Lich:   1,
			models.Gorgon: 1,
			models.Harpy:  1,
		},
	}

	// Base for level 50 is 750
	// Feats: 1.0
	// Legendaries: 1 + (0.10 + 0.10 + 0.10) + 0.20 (Group 1) = 1.5
	// Total: 750 * 1.5 = 1125
	expected := 1125
	result := CalculateTimeShards(plan)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Test extra group bonus
	plan.GroupBonusCount = 2
	plan.LegendaryCounts = map[models.LegendaryID]int{
		models.Lich:   2,
		models.Gorgon: 2,
		models.Harpy:  2,
	}
	// Legendaries: 1 + (0.10+0.05 + 0.10+0.05 + 0.10+0.05) + 0.20*2 = 1 + 0.45 + 0.40 = 1.85
	// Total: 750 * 1.0 * 1.85 = 1387.5 -> 1387 (using math.Floor)
	expected = 1387
	result = CalculateTimeShards(plan)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCalculateExperimentCost(t *testing.T) {
	plan := models.Plan{
		ExperimentLevels: map[models.ExperimentID]int{
			models.ExpSeasoning: 2, // Level 1: 25, Level 2: 250 -> Total: 275
		},
	}

	expected := 275
	result := CalculateExperimentCost(plan)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	plan.ExperimentLevels[models.ExpTaste] = 1 // Level 1: 50
	expected = 275 + 50
	result = CalculateExperimentCost(plan)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
