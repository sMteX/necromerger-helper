package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

type ExperimentSummary struct {
	ID                models.ExperimentID `json:"id"`
	CurrentLevel      int                 `json:"currentLevel"`
	CurrentLevelValue string              `json:"currentLevelValue"`
	NextLevelCost     string              `json:"nextLevelCost"`
	NextLevelValue    string              `json:"nextLevelValue"`
	MaxLevel          bool                `json:"maxLevel"`
}

type RecalculateResponse struct {
	TotalShards      int                                         `json:"totalShards"`
	BaseShards       int                                         `json:"baseShards"`
	FeatMultiplier   float64                                     `json:"featMultiplier"`
	LegendMultiplier float64                                     `json:"legendMultiplier"`
	OtherMultiplier  float64                                     `json:"otherMultiplier"`
	ExperimentCost   int                                         `json:"experimentCost"`
	Remaining        int                                         `json:"remaining"`
	Experiments      map[models.ExperimentID]ExperimentSummary   `json:"experiments"`
	RuneTotal        map[models.RuneType]int                     `json:"runeTotal"`
	RuneNeeded       map[models.RuneType]int                     `json:"runeNeeded"`
	LegendaryRunes   map[models.LegendaryID]calculator.RuneCosts `json:"legendaryRunes"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func RecalculateHandler(w http.ResponseWriter, r *http.Request) {
	var plan models.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shards := calculator.CalculateTimeShards(plan)
	baseShards := calculator.DevourerBaseShards[plan.DevourerLevel]
	featMultiplier := 1.0 + (float64(plan.FeatTiers) * 0.10)

	expCost := calculator.CalculateExperimentCost(plan)
	totalRunes, neededRunes := calculator.CalculateTotalRunes(plan)

	// Experiment summaries
	expSummaries := make(map[models.ExperimentID]ExperimentSummary)
	for _, exp := range calculator.Experiments {
		currentLevel := plan.ExperimentLevels[exp.ID]

		summary := ExperimentSummary{
			ID:           exp.ID,
			CurrentLevel: currentLevel,
		}

		if currentLevel > 0 {
			prevVal := 0.0
			if currentLevel-1 < len(exp.Levels) {
				prevVal = exp.Levels[currentLevel-1].PrevValue
				summary.CurrentLevelValue = fmt.Sprintf("%s -> %s",
					calculator.FormatExperimentValue(exp.ID, exp.Tier, prevVal),
					calculator.FormatExperimentValue(exp.ID, exp.Tier, exp.Levels[currentLevel-1].Value))
			}
		} else {
			summary.CurrentLevelValue = calculator.FormatExperimentValue(exp.ID, exp.Tier, 0)
		}

		if currentLevel < len(exp.Levels) {
			nextLevel := exp.Levels[currentLevel]
			summary.NextLevelCost = calculator.FormatLargeNumber(nextLevel.Cost)
			summary.NextLevelValue = calculator.FormatExperimentValue(exp.ID, exp.Tier, nextLevel.Value)
		} else {
			summary.MaxLevel = true
		}
		expSummaries[exp.ID] = summary
	}

	// Legendary rune costs
	legRunes := make(map[models.LegendaryID]calculator.RuneCosts)
	for id := range calculator.LegendaryRecipes {
		legRunes[id] = calculator.GetLegendaryRuneCost(id)
	}

	resp := RecalculateResponse{
		TotalShards:     shards,
		BaseShards:      baseShards,
		FeatMultiplier:  featMultiplier,
		OtherMultiplier: 1.0 + plan.OtherMultiplier,
		ExperimentCost:  expCost,
		Remaining:       shards - expCost,
		Experiments:     expSummaries,
		RuneTotal:       totalRunes,
		RuneNeeded:      neededRunes,
		LegendaryRunes:  legRunes,
	}

	// Calculate legend multiplier for the breakdown
	if baseShards > 0 && featMultiplier > 0 && resp.OtherMultiplier > 0 {
		resp.LegendMultiplier = float64(shards-plan.LeftoverShards) / (float64(baseShards) * featMultiplier * resp.OtherMultiplier)
	} else {
		resp.LegendMultiplier = 1.0
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func ParsePlanFromJSON(r *http.Request) (models.Plan, error) {
	var p models.Plan
	err := json.NewDecoder(r.Body).Decode(&p)
	return p, err
}
