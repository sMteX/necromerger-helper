package calculator

import (
	"math"

	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

func CalculateTotalRunes(plan models.Plan) (total models.RuneCosts, needed models.RuneCosts) {
	total = make(models.RuneCosts)
	needed = make(models.RuneCosts)

	// Calculate total legendaries that MUST be produced from stations
	totalNeeded := make(map[models.LegendaryID]int)
	for id, count := range plan.LegendaryCounts {
		totalNeeded[id] += count
		recipe := data.LegendaryRecipes[id]
		for _, reqID := range recipe.Requires {
			totalNeeded[reqID] += count
		}
	}

	// Calculate what we already have (including nested)
	totalHave := make(map[models.LegendaryID]int)
	for id, count := range plan.PossessedLegendaries {
		totalHave[id] += count
		recipe := data.LegendaryRecipes[id]
		for _, reqID := range recipe.Requires {
			totalHave[reqID] += count
		}
	}

	// How many of each station-produced legendary we need to buy
	toBuy := make(map[models.LegendaryID]int)
	for id, recipe := range data.LegendaryRecipes {
		if recipe.StationID == "" {
			continue
		}
		diff := totalNeeded[id] - totalHave[id]
		if diff > 0 {
			toBuy[id] = diff
		}
	}

	// Calculate total runes for EVERYTHING targeted (regardless of possession)
	for id, count := range totalNeeded {
		recipe := data.LegendaryRecipes[id]
		if recipe.StationID == "" {
			continue
		}
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := data.StationCosts[recipe.StationID]
		for runeType, amount := range stationCost {
			total[runeType] += amount * multiplier * count
		}
	}

	// Calculate "Needed" runes (only for what we need to buy, minus possessed runes)
	neededRunes := make(models.RuneCosts)
	for id, count := range toBuy {
		recipe := data.LegendaryRecipes[id]
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := data.StationCosts[recipe.StationID]
		for runeType, amount := range stationCost {
			neededRunes[runeType] += amount * multiplier * count
		}
	}

	for runeType, amount := range neededRunes {
		possessed := plan.PossessedRunes[runeType]
		if amount > possessed {
			needed[runeType] = amount - possessed
		} else {
			needed[runeType] = 0
		}
	}

	return total, needed
}

func GetLegendaryRuneCost(id models.LegendaryID) models.RuneCosts {
	costs := make(models.RuneCosts)
	recipe, ok := data.LegendaryRecipes[id]
	if !ok {
		return costs
	}

	if recipe.StationID != "" {
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := data.StationCosts[recipe.StationID]
		for runeType, amount := range stationCost {
			costs[runeType] += amount * multiplier
		}
	}

	for _, reqID := range recipe.Requires {
		reqCosts := GetLegendaryRuneCost(reqID)
		for runeType, amount := range reqCosts {
			costs[runeType] += amount
		}
	}

	return costs
}
