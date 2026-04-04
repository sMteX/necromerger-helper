package models

type LegendaryGroup int

const (
	Group1 LegendaryGroup = iota + 1
	Group2
	Group3
	Group4
)

type LegendaryID string

const (
	Lich        LegendaryID = "lich"
	Gorgon      LegendaryID = "gorgon"
	Harpy       LegendaryID = "harpy"
	Reaper      LegendaryID = "reaper"
	Cyclops     LegendaryID = "cyclops"
	Archdemon   LegendaryID = "archdemon"
	TheCursed   LegendaryID = "the_cursed"
	TheColossus LegendaryID = "the_colossus"
	TheInfernal LegendaryID = "the_infernal"
	RoboChicken LegendaryID = "robo_chicken"
	ShieldBot   LegendaryID = "shield_bot"
	SoulStalker LegendaryID = "soul_stalker"
)

type Legendary struct {
	ID           LegendaryID
	Name         string
	Group        LegendaryGroup
	FirstBonus   float64 // e.g., 0.10 for 10%
	Subsequent   float64 // e.g., 0.05 for 5%
	MaxInstances int     // 0 for uncapped
}

type ExperimentID string

const (
	ExpSeasoning    ExperimentID = "seasoning"
	ExpStrength     ExperimentID = "strength"
	ExpTaste        ExperimentID = "taste"
	ExpCapacity     ExperimentID = "capacity"
	ExpBodySnatcher ExperimentID = "body_snatcher"
	ExpWeakening    ExperimentID = "weakening"
	ExpDamageCap    ExperimentID = "damage_cap"
	ExpIceChest     ExperimentID = "ice_chest"
	ExpPoisonChest  ExperimentID = "poison_chest"
	ExpBloodChest   ExperimentID = "blood_chest"
	ExpMoonChest    ExperimentID = "moon_chest"
	ExpDeathChest   ExperimentID = "death_chest"
	ExpCosmicChest  ExperimentID = "cosmic_chest"
	ExpSeasoning2   ExperimentID = "seasoning_2"
	ExpStrength2    ExperimentID = "strength_2"
	ExpTaste2       ExperimentID = "taste_2"
	ExpCapacity2    ExperimentID = "capacity_2"
)

type ExperimentTier int

const (
	TierPre100 ExperimentTier = iota + 1
	TierPost100
)

type ExperimentLevel struct {
	Level     int
	Cost      int
	Value     float64
	PrevValue float64
}

type Experiment struct {
	ID        ExperimentID
	Name      string
	Tier      ExperimentTier
	Levels    []ExperimentLevel
	IsSpecial bool // For things like Body Snatcher
}

type Plan struct {
	ID               int
	Name             string
	DevourerLevel    int
	FeatTiers        int
	OtherMultiplier  float64 // The "Other" category (skins, etc.)
	GroupBonusCount  int     // Number of times group bonuses can be claimed (default 1)
	LegendaryCounts  map[LegendaryID]int
	ExperimentLevels map[ExperimentID]int
	Notes            string
}
