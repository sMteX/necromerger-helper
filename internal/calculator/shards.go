package calculator

import (
	"math"

	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

func CalculateTimeShards(plan models.Plan) int {
	base := data.DevourerBaseShards[plan.DevourerLevel]
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

	for _, leg := range data.Legendaries {
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
