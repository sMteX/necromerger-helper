package calculator

import (
	"math"

	"github.com/sMteX/necro-prestige-planner/internal/models"
)

var DevourerBaseShards = map[int]int{
	35: 150, 40: 275, 45: 500, 50: 750, 55: 1000, 60: 1500, 65: 2000, 70: 3250,
	75: 4500, 80: 5750, 85: 7500, 90: 10000, 95: 12500, 100: 15000, 150: 40000,
	200: 65000, 300: 150000, 400: 275000, 500: 450000, 600: 700000, 700: 1050000,
	800: 1550000, 900: 2250000, 1000: 3250000,
}

var Legendaries = []models.Legendary{
	{ID: models.Lich, Name: "Lich", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Gorgon, Name: "Gorgon", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Harpy, Name: "Harpy", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Reaper, Name: "Reaper", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	{ID: models.Cyclops, Name: "Cyclops", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	{ID: models.Archdemon, Name: "Archdemon", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	{ID: models.TheCursed, Name: "The Cursed", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.TheColossus, Name: "The Colossus", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.TheInfernal, Name: "The Infernal", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.RoboChicken, Name: "Robo Chicken", Group: models.Group4, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	{ID: models.ShieldBot, Name: "Shield Bot", Group: models.Group4, FirstBonus: 0.30, Subsequent: 0.15, MaxInstances: 4},
	{ID: models.SoulStalker, Name: "Soul Stalker", Group: models.Group4, FirstBonus: 0.40, Subsequent: 0.20},
}

func CalculateTimeShards(plan models.Plan) int {
	base := DevourerBaseShards[plan.DevourerLevel]
	if base == 0 {
		return 0
	}

	featMultiplier := 1.0 + (float64(plan.FeatTiers) * 0.10)
	legMultiplier := CalculateLegendMultiplier(plan)
	otherMultiplier := 1.0 + plan.OtherMultiplier

	total := float64(base) * featMultiplier * legMultiplier * otherMultiplier
	return int(math.Floor(total)) + plan.LeftoverShards
}

func CalculateLegendMultiplier(plan models.Plan) float64 {
	legBonus := 0.0
	groupMinCounts := make(map[models.LegendaryGroup]int)
	groups := []models.LegendaryGroup{models.Group1, models.Group2, models.Group3, models.Group4}
	for _, g := range groups {
		groupMinCounts[g] = math.MaxInt32
	}

	for _, leg := range Legendaries {
		count := plan.LegendaryCounts[leg.ID]
		if count > 0 {
			bonusCount := count
			if leg.MaxInstances > 0 && bonusCount > leg.MaxInstances {
				bonusCount = leg.MaxInstances
			}

			legBonus += leg.FirstBonus
			if bonusCount > 1 {
				legBonus += float64(bonusCount-1) * leg.Subsequent
			}
		}

		if count < groupMinCounts[leg.Group] {
			groupMinCounts[leg.Group] = count
		}
	}

	groupBonuses := map[models.LegendaryGroup]float64{
		models.Group1: 0.20, models.Group2: 0.40, models.Group3: 0.80, models.Group4: 0.60,
	}

	for group, minCount := range groupMinCounts {
		if minCount == math.MaxInt32 || minCount == 0 {
			continue
		}
		// The number of times group bonus is claimed is limited by the feature level
		// and how many of each we actually have.
		allowedClaims := 1 + plan.GroupBonusCount
		claims := int(math.Min(float64(minCount), float64(allowedClaims)))
		legBonus += float64(claims) * groupBonuses[group]
	}

	return 1.0 + legBonus
}
