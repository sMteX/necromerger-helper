package calculator

import (
	"math"

	"github.com/sMteX/necromerger-helper/internal/data"
	"github.com/sMteX/necromerger-helper/internal/models"
)

// Station base caps per level (index 0 = L1, ..., index 5 = L6).
var manaPoolCaps = []int{2000, 4000, 6000, 10000, 15000, 20000}
var slimeVatCaps = []int{1500, 3000, 5000, 7500, 10000, 15000}
var darkStoreCaps = []int{1000, 2000, 3000, 5000, 7000, 10000}

// L1 rune costs for each resource cap station type.
var manaPoolL1Cost = models.RuneCosts{models.RuneIce: 10, models.RunePoison: 5}
var slimeVatL1Cost = models.RuneCosts{models.RunePoison: 10, models.RuneBlood: 5}
var darkStoreL1Cost = models.RuneCosts{models.RuneBlood: 10, models.RuneMoon: 5}

// Per-resource relic bonus per level (index = level - 1), as decimals.
var perResourceRelicBonus = []float64{0.02, 0.04, 0.06, 0.08, 0.10, 0.13, 0.16, 0.20, 0.25, 0.30}

// All-resources relic bonus per level (index = level - 1), as decimals.
var allResourceRelicBonus = []float64{0.03, 0.06, 0.09, 0.12, 0.15}

// ResourceCapInput contains all values needed to calculate resource caps.
// Percentage values are decimals (e.g. 0.06 means +6%).
type ResourceCapInput struct {
	ManaPools  [6]int
	SlimeVats  [6]int
	DarkStores [6]int

	ServOResource models.ResourceType // empty string if Serv-O not applicable
	ServOUpgraded bool

	// Base boost skins (flat cap increase)
	WizardSkin bool // +2000 Mana
	OozingSkin bool // +2000 Slime
	SidSkin    bool // +2000 Darkness

	// Multiplicative skins
	SantaSkin    bool // +5% Mana
	BirthdaySkin bool // +5% Slime
	WitchSkin    bool // +5% Darkness
	GoodSkin     bool // +2% All
	RoyaltySkin  bool // +5% All

	GoldenBoosts bool // +25% All

	// Spell levels: per-resource (pages 1-3, +5%/level, max 5) and all-resources (page 6, +5%/level, max 2)
	ManaSpell, SlimeSpell, DarknessSpell, AllResourcesSpell int

	// Relic levels: per-resource (levels 1-10) and all-resources (levels 1-5)
	ManaRelic, SlimeRelic, DarknessRelic, AllResourcesRelic int

	CapacityExp1 int // Capacity Experiment I level (0-9)

	// PearlBonus is the per-resource % bonus from pearls as shown in game stats.
	// Already includes the all-resources pearl contribution. Decimal format.
	PearlBonus map[models.ResourceType]float64

	CapacityExp2 int // Capacity Experiment II level (0-9)
}

// ResourceCapResult holds calculated caps and the effective multiplier per resource.
// The multiplier is needed externally to calculate station option costs.
type ResourceCapResult struct {
	Mana, Slime, Darkness            int
	ManaMulti, SlimeMulti, DarkMulti float64
}

// StationOptionResult describes how many stations of a given level close a resource gap.
type StationOptionResult struct {
	Level    int
	Count    int
	RuneCost models.RuneCosts
}

func relicBonus(level int, table []float64) float64 {
	if level <= 0 || level > len(table) {
		return 0
	}
	return table[level-1]
}

func stationTotalCap(counts [6]int, caps []int) int {
	total := 0
	for i := 0; i < 6; i++ {
		total += counts[i] * caps[i]
	}
	return total
}

func capacityExp1Bonus(level int) float64 {
	if level <= 0 {
		return 0
	}
	for _, exp := range data.Experiments {
		if exp.ID == models.ExpCapacity && level <= len(exp.Levels) {
			return exp.Levels[level-1].Value
		}
	}
	return 0
}

func capacityExp2Multiplier(level int) float64 {
	if level <= 0 {
		return 1.0
	}
	for _, exp := range data.Experiments {
		if exp.ID == models.ExpCapacity2 && level <= len(exp.Levels) {
			return exp.Levels[level-1].Value
		}
	}
	return 1.0
}

// CalculateResourceCaps computes the effective cap for each resource type.
func CalculateResourceCaps(input ResourceCapInput) ResourceCapResult {
	exp2 := capacityExp2Multiplier(input.CapacityExp2)

	// Additive bonus components shared across all resources.
	shared := 0.0
	if input.GoldenBoosts {
		shared += 0.25
	}
	if input.GoodSkin {
		shared += 0.02
	}
	if input.RoyaltySkin {
		shared += 0.05
	}
	shared += float64(input.AllResourcesSpell) * 0.05
	shared += relicBonus(input.AllResourcesRelic, allResourceRelicBonus)
	shared += capacityExp1Bonus(input.CapacityExp1)

	// Mana
	manaBase := 20000 + stationTotalCap(input.ManaPools, manaPoolCaps)
	if input.ServOResource == models.ResourceMana {
		if input.ServOUpgraded {
			manaBase += 9000
		} else {
			manaBase += 6000
		}
	}
	if input.WizardSkin {
		manaBase += 2000
	}
	manaMulti := 1.0 + shared +
		float64(input.ManaSpell)*0.05 +
		relicBonus(input.ManaRelic, perResourceRelicBonus) +
		input.PearlBonus[models.ResourceMana]
	if input.SantaSkin {
		manaMulti += 0.05
	}

	// Slime
	slimeBase := 15000 + stationTotalCap(input.SlimeVats, slimeVatCaps)
	if input.ServOResource == models.ResourceSlime {
		if input.ServOUpgraded {
			slimeBase += 7500
		} else {
			slimeBase += 4500
		}
	}
	if input.OozingSkin {
		slimeBase += 2000
	}
	slimeMulti := 1.0 + shared +
		float64(input.SlimeSpell)*0.05 +
		relicBonus(input.SlimeRelic, perResourceRelicBonus) +
		input.PearlBonus[models.ResourceSlime]
	if input.BirthdaySkin {
		slimeMulti += 0.05
	}

	// Darkness
	darkBase := 10000 + stationTotalCap(input.DarkStores, darkStoreCaps)
	if input.ServOResource == models.ResourceDarkness {
		if input.ServOUpgraded {
			darkBase += 6000
		} else {
			darkBase += 3000
		}
	}
	if input.SidSkin {
		darkBase += 2000
	}
	darkMulti := 1.0 + shared +
		float64(input.DarknessSpell)*0.05 +
		relicBonus(input.DarknessRelic, perResourceRelicBonus) +
		input.PearlBonus[models.ResourceDarkness]
	if input.WitchSkin {
		darkMulti += 0.05
	}

	return ResourceCapResult{
		Mana:       int(math.Round(float64(manaBase) * manaMulti * exp2)),
		Slime:      int(math.Round(float64(slimeBase) * slimeMulti * exp2)),
		Darkness:   int(math.Round(float64(darkBase) * darkMulti * exp2)),
		ManaMulti:  manaMulti * exp2,
		SlimeMulti: slimeMulti * exp2,
		DarkMulti:  darkMulti * exp2,
	}
}

// ResourceTargets returns the per-resource cap targets for a combined threshold.
// Respects any fixed targets; remaining combined budget is split by the documented ratios.
func ResourceTargets(threshold int, fixed map[models.ResourceType]int) (mana, slime, darkness int) {
	if fixed == nil {
		fixed = map[models.ResourceType]int{}
	}
	fixedMana, hasMana := fixed[models.ResourceMana]
	fixedSlime, hasSlime := fixed[models.ResourceSlime]
	fixedDark, hasDark := fixed[models.ResourceDarkness]

	switch {
	case hasMana && hasSlime && hasDark:
		return fixedMana, fixedSlime, fixedDark
	case hasMana:
		rem := float64(threshold - fixedMana)
		return fixedMana, int(math.Round(rem * 2 / 3)), int(math.Round(rem / 3))
	case hasSlime:
		rem := float64(threshold - fixedSlime)
		return int(math.Round(rem * 3 / 4)), fixedSlime, int(math.Round(rem / 4))
	case hasDark:
		rem := float64(threshold - fixedDark)
		return int(math.Round(rem * 2 / 3)), int(math.Round(rem / 3)), fixedDark
	default:
		// ~69% Mana, 16.8% Slime, 14.2% Darkness
		t := float64(threshold)
		return int(math.Round(t * 0.69)), int(math.Round(t * 0.168)), int(math.Round(t * 0.142))
	}
}

// StationOptions returns how many stations of each level (1-6) are needed to close gap,
// given the effective multiplier already applied to base cap contributions.
func StationOptions(gap int, caps []int, l1Cost models.RuneCosts, multiplier float64) []StationOptionResult {
	if gap <= 0 {
		return nil
	}
	options := make([]StationOptionResult, 6)
	for i := 0; i < 6; i++ {
		effectiveCap := float64(caps[i]) * multiplier
		count := int(math.Ceil(float64(gap) / effectiveCap))
		l1Equivalent := count * (1 << i) // 2^i L1 stations needed per one level-(i+1) station
		runeCost := make(models.RuneCosts, len(l1Cost))
		for rune, unitCost := range l1Cost {
			runeCost[rune] = l1Equivalent * unitCost
		}
		options[i] = StationOptionResult{Level: i + 1, Count: count, RuneCost: runeCost}
	}
	return options
}

func ManaStationOptions(gap int, multiplier float64) []StationOptionResult {
	return StationOptions(gap, manaPoolCaps, manaPoolL1Cost, multiplier)
}

func SlimeStationOptions(gap int, multiplier float64) []StationOptionResult {
	return StationOptions(gap, slimeVatCaps, slimeVatL1Cost, multiplier)
}

func DarknessStationOptions(gap int, multiplier float64) []StationOptionResult {
	return StationOptions(gap, darkStoreCaps, darkStoreL1Cost, multiplier)
}
