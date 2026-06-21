package calculator

import (
	"math"

	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

type PrestigePlanResult struct {
	BaseShards     int
	LeftoverShards int
	TotalShards    int
	ExperimentCost int
	NetShards      int

	FeatMultiplier   float64
	LegendMultiplier float64
	OtherMultiplier  float64

	LegendaryBonuses      map[models.LegendaryID]float64
	LegendaryGroupBonuses map[models.LegendaryGroup]float64

	RuneTotal  models.RuneCosts
	RuneNeeded models.RuneCosts
}

func Calculate(plan models.Plan) PrestigePlanResult {
	baseShards := data.DevourerBaseShards[plan.DevourerLevel]
	featMultiplier := 1.0 + float64(plan.FeatTiers)*0.10
	legMultiplier, legBonuses, groupBonuses := calculateLegendBreakdown(plan)

	totalShards := int(math.Floor(float64(baseShards)*featMultiplier*legMultiplier*plan.OtherMultiplier)) + plan.LeftoverShards
	expCost := CalculateExperimentCost(plan)
	runeTotal, runeNeeded := CalculateTotalRunes(plan)

	return PrestigePlanResult{
		BaseShards:            baseShards,
		LeftoverShards:        plan.LeftoverShards,
		FeatMultiplier:        featMultiplier,
		LegendMultiplier:      legMultiplier,
		OtherMultiplier:       plan.OtherMultiplier,
		TotalShards:           totalShards,
		ExperimentCost:        expCost,
		NetShards:             totalShards - expCost,
		LegendaryBonuses:      legBonuses,
		LegendaryGroupBonuses: groupBonuses,
		RuneTotal:             runeTotal,
		RuneNeeded:            runeNeeded,
	}
}

func CalculateTimeShards(plan models.Plan) int {
	base := data.DevourerBaseShards[plan.DevourerLevel]
	if base == 0 {
		return 0
	}
	featMultiplier := 1.0 + (float64(plan.FeatTiers) * 0.10)
	legMultiplier := CalculateLegendMultiplier(plan)
	total := float64(base) * featMultiplier * legMultiplier * plan.OtherMultiplier
	return int(math.Floor(total)) + plan.LeftoverShards
}

func CalculateLegendMultiplier(plan models.Plan) float64 {
	multiplier, _, _ := calculateLegendBreakdown(plan)
	return multiplier
}

func calculateLegendBreakdown(plan models.Plan) (multiplier float64, legBonuses map[models.LegendaryID]float64, groupBonuses map[models.LegendaryGroup]float64) {
	legBonuses = make(map[models.LegendaryID]float64)
	groupBonuses = make(map[models.LegendaryGroup]float64)

	totalBonus := 0.0
	groupMinCounts := make(map[models.LegendaryGroup]int)
	for _, g := range []models.LegendaryGroup{models.Group1, models.Group2, models.Group3, models.Group4} {
		groupMinCounts[g] = math.MaxInt32
	}

	for _, leg := range data.Legendaries {
		count := plan.LegendaryCounts[leg.ID]
		if count > 0 {
			bonusCount := count
			if leg.MaxInstances > 0 && bonusCount > leg.MaxInstances {
				bonusCount = leg.MaxInstances
			}
			indBonus := leg.FirstBonus
			if bonusCount > 1 {
				indBonus += float64(bonusCount-1) * leg.Subsequent
			}
			legBonuses[leg.ID] = indBonus
			totalBonus += indBonus
		}
		if count < groupMinCounts[leg.Group] {
			groupMinCounts[leg.Group] = count
		}
	}

	groupBonusRates := map[models.LegendaryGroup]float64{
		models.Group1: 0.20, models.Group2: 0.40, models.Group3: 0.80, models.Group4: 0.60,
	}

	for _, group := range []models.LegendaryGroup{models.Group1, models.Group2, models.Group3, models.Group4} {
		minCount := groupMinCounts[group]
		if minCount == math.MaxInt32 || minCount == 0 {
			continue
		}
		allowedClaims := plan.GroupBonusCount
		var claims int
		if group == models.Group3 {
			// Group 3 legendaries are each capped at 1 copy, so minCount is always 1 when
			// the set is complete. The game still grants the full allowedClaims for this group.
			claims = allowedClaims
		} else {
			claims = int(math.Min(float64(minCount), float64(allowedClaims)))
		}
		gb := float64(claims) * groupBonusRates[group]
		groupBonuses[group] = gb
		totalBonus += gb
	}

	multiplier = 1.0 + totalBonus
	return
}
