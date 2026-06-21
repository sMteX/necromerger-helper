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

	// How many of each station-produced legendary we still need to craft.
	// Non-station legendaries (e.g. those obtained from drops) are skipped —
	// they have no rune cost.
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

	// "Total" = gross rune cost to craft only the legendaries we're missing.
	// We intentionally do NOT include legendaries already possessed here, so
	// that Total - Have = Need is a meaningful equation in the UI.
	for id, count := range toBuy {
		recipe := data.LegendaryRecipes[id]
		// Each level doubles the cost; if the recipe returns the L1 legendary
		// after upgrading, subtract one copy's worth of cost.
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := data.StationCosts[recipe.StationID]
		for runeType, amount := range stationCost {
			total[runeType] += amount * multiplier * count
		}
	}

	// "Needed" = Total minus what the player already has in their rune inventory.
	// Clamped to 0 — if possessed exceeds total, no runes are needed.
	for runeType, amount := range total {
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
