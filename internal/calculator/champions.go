package calculator

import (
	"math"
	"sort"

	"github.com/sMteX/necromerger-helper/internal/data"
	"github.com/sMteX/necromerger-helper/internal/models"
)

type ChampionEfficiencyInput struct {
	ChampionID   models.ChampionID
	ManaCap      int
	SlimeCap     int
	DarknessCap  int
	SpeedPercent int // e.g. 185 means 1.85× multiplier
}

type StationEfficiencyResult struct {
	StationName     string
	Level           int
	Hacked          bool
	Resource        models.ResourceType
	ExpectedSummons float64
}

type ChampionEfficiencyResult struct {
	Stations []StationEfficiencyResult // sorted by ExpectedSummons descending
}

// pointsPerL2Equiv returns the expected summoning points per L2-part-equivalent
// when merging all the way to maximum level.
// Formula: Σ_k MergePoints[k] / 2^(k+1)
func pointsPerL2Equiv(creature models.CreatureType) float64 {
	total := 0.0
	for k, pts := range creature.MergePoints {
		total += float64(pts) / math.Pow(2, float64(k+1))
	}
	return total
}

// expectedL2EquivPerTap returns the expected L2-part-equivalents contributed
// toward the given champion per tap of a station variant.
// L1 parts count as 0.5 (two needed to make one L2), L2 parts count as 1.0.
func expectedL2EquivPerTap(variant models.StationVariant, relevantCreatures map[models.CreatureID]struct{}) map[models.CreatureID]float64 {
	result := make(map[models.CreatureID]float64)
	for _, drop := range variant.Drops {
		if _, ok := relevantCreatures[drop.CreatureID]; !ok {
			continue
		}
		weight := 1.0
		if drop.PartLevel == models.PartLevelOne {
			weight = 0.5
		}
		result[drop.CreatureID] += drop.Probability * weight
	}
	return result
}

func CalculateChampionEfficiency(input ChampionEfficiencyInput) ChampionEfficiencyResult {
	champion, ok := data.ChampionsById[input.ChampionID]
	if !ok {
		return ChampionEfficiencyResult{}
	}

	speedMultiplier := float64(input.SpeedPercent) / 100.0

	relevantCreatures := make(map[models.CreatureID]struct{}, len(champion.Creatures))
	for _, cid := range champion.Creatures {
		relevantCreatures[cid] = struct{}{}
	}

	capFor := func(resource models.ResourceType) int {
		switch resource {
		case models.ResourceMana:
			return input.ManaCap
		case models.ResourceSlime:
			return input.SlimeCap
		case models.ResourceDarkness:
			return input.DarknessCap
		}
		return 0
	}

	var results []StationEfficiencyResult

	for _, station := range data.SummoningStations {
		cap := capFor(station.Resource)
		if cap <= 0 {
			continue
		}

		for _, variant := range station.Variants {
			taps := cap / variant.Cost

			l2EquivPerTap := expectedL2EquivPerTap(variant, relevantCreatures)

			expectedPointsPerTap := 0.0
			for cid, l2Equiv := range l2EquivPerTap {
				creature := data.Creatures[cid]
				expectedPointsPerTap += l2Equiv * pointsPerL2Equiv(creature)
			}

			totalPoints := float64(taps) * expectedPointsPerTap * speedMultiplier
			expectedSummons := totalPoints / float64(champion.Threshold)

			results = append(results, StationEfficiencyResult{
				StationName:     station.Name,
				Level:           variant.Level,
				Hacked:          variant.Hacked,
				Resource:        station.Resource,
				ExpectedSummons: expectedSummons,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].ExpectedSummons > results[j].ExpectedSummons
	})

	return ChampionEfficiencyResult{Stations: results}
}
