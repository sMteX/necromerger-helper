package calculator

import (
	"math"

	"github.com/sMteX/necro-prestige-planner/internal/models"
)

type RuneCosts map[models.RuneType]int

var StationCosts = map[models.StationID]RuneCosts{
	models.StationGrave:          {models.RuneIce: 20},
	models.StationSupplyCupboard: {models.RunePoison: 20},
	models.StationFoulChicken:    {models.RuneIce: 30, models.RunePoison: 15},
	models.StationAltar:          {models.RuneBlood: 20},
	models.StationLectern:        {models.RuneIce: 50, models.RuneMoon: 20},
	models.StationFridge:         {models.RunePoison: 50, models.RuneMoon: 20},
	models.StationPortal:         {models.RuneBlood: 30, models.RuneDeath: 30},
	models.StationCrashedSaucer:  {models.RuneCosmic: 20},
	models.StationSoulGrinder:    {models.RuneDeath: 200, models.RuneCosmic: 200},
}

type LegendaryRecipe struct {
	StationID models.StationID
	Levels    int
	ReturnsL1 bool
	Requires  []models.LegendaryID // For Group 3
}

var LegendaryRecipes = map[models.LegendaryID]LegendaryRecipe{
	models.Lich:        {StationID: models.StationGrave, Levels: 5, ReturnsL1: true},
	models.Gorgon:      {StationID: models.StationSupplyCupboard, Levels: 5, ReturnsL1: true},
	models.Harpy:       {StationID: models.StationAltar, Levels: 5, ReturnsL1: true},
	models.Reaper:      {StationID: models.StationLectern, Levels: 5, ReturnsL1: true},
	models.Cyclops:     {StationID: models.StationFridge, Levels: 5, ReturnsL1: true},
	models.Archdemon:   {StationID: models.StationPortal, Levels: 5, ReturnsL1: true},
	models.RoboChicken: {StationID: models.StationFoulChicken, Levels: 5, ReturnsL1: false},
	models.ShieldBot:   {StationID: models.StationCrashedSaucer, Levels: 5, ReturnsL1: true},
	models.SoulStalker: {StationID: models.StationSoulGrinder, Levels: 2, ReturnsL1: true},

	models.TheCursed:   {Requires: []models.LegendaryID{models.Lich, models.Reaper}},
	models.TheColossus: {Requires: []models.LegendaryID{models.Gorgon, models.Cyclops}},
	models.TheInfernal: {Requires: []models.LegendaryID{models.Harpy, models.Archdemon}},
}

func CalculateTotalRunes(plan models.Plan) (total RuneCosts, needed RuneCosts) {
	total = make(RuneCosts)
	needed = make(RuneCosts)

	// Calculate total legendaries that MUST be produced from stations
	totalNeeded := make(map[models.LegendaryID]int)
	for id, count := range plan.LegendaryCounts {
		totalNeeded[id] += count
		recipe := LegendaryRecipes[id]
		for _, reqID := range recipe.Requires {
			totalNeeded[reqID] += count
		}
	}

	// Calculate what we already have (including nested)
	totalHave := make(map[models.LegendaryID]int)
	for id, count := range plan.PossessedLegendaries {
		totalHave[id] += count
		recipe := LegendaryRecipes[id]
		for _, reqID := range recipe.Requires {
			totalHave[reqID] += count
		}
	}

	// How many of each station-produced legendary we need to buy
	toBuy := make(map[models.LegendaryID]int)
	for id, recipe := range LegendaryRecipes {
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
		recipe := LegendaryRecipes[id]
		if recipe.StationID == "" {
			continue
		}
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := StationCosts[recipe.StationID]
		for runeType, amount := range stationCost {
			total[runeType] += amount * multiplier * count
		}
	}

	// Calculate "Needed" runes (only for what we need to buy, minus possessed runes)
	neededRunes := make(RuneCosts)
	for id, count := range toBuy {
		recipe := LegendaryRecipes[id]
		multiplier := int(math.Pow(2, float64(recipe.Levels)))
		if recipe.ReturnsL1 {
			multiplier -= 1
		}
		stationCost := StationCosts[recipe.StationID]
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
