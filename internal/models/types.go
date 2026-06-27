package models

type StationID string

const (
	StationGrave          StationID = "grave"
	StationSupplyCupboard StationID = "supply_cupboard"
	StationFoulChicken    StationID = "foul_chicken"
	StationAltar          StationID = "altar"
	StationLectern        StationID = "lectern"
	StationFridge         StationID = "fridge"
	StationPortal         StationID = "portal"
	StationCrashedSaucer  StationID = "crashed_saucer"
	StationSoulGrinder    StationID = "soul_grinder"
)

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

type RuneType string

const (
	RuneIce    RuneType = "Ice"
	RunePoison RuneType = "Poison"
	RuneBlood  RuneType = "Blood"
	RuneMoon   RuneType = "Moon"
	RuneDeath  RuneType = "Death"
	RuneCosmic RuneType = "Cosmic"
)

type ResourceType string

const (
	ResourceMana     ResourceType = "mana"
	ResourceSlime    ResourceType = "slime"
	ResourceDarkness ResourceType = "darkness"
)

type Plan struct {
	ID                   int                  `json:"id"`
	Name                 string               `json:"name"`
	DevourerLevel        int                  `json:"devourerLevel"`
	FeatTiers            int                  `json:"featTiers"`
	OtherMultiplier      float64              `json:"otherMultiplier"`  // The "Other" category (skins, etc.) - will already include the base 100% => e.g. 1.72
	GroupBonusCount      int                  `json:"groupBonusCount"`  // Number of times group bonuses can be claimed (default 1)
	LeftoverShards       int                  `json:"leftoverShards"`   // Shards remaining from previous prestige
	LegendaryCounts      map[LegendaryID]int  `json:"legendaryCounts"`  // the planned amount of legendaries
	ExperimentLevels     map[ExperimentID]int `json:"experimentLevels"` // the planned setup of experiments
	PossessedRunes       map[RuneType]int     `json:"possessedRunes"`
	PossessedLegendaries map[LegendaryID]int  `json:"possessedLegendaries"`
	Notes                string               `json:"notes"`
}

type LegendaryRecipe struct {
	StationID StationID
	Levels    int
	ReturnsL1 bool
	Requires  []LegendaryID // For Group 3
}

type RuneCosts map[RuneType]int

// ── Champion summoning ────────────────────────────────────────────────────────

type ChampionID string

const (
	ChampionPeasant ChampionID = "Peasant"
	ChampionKnight  ChampionID = "Knight"
	ChampionCleric  ChampionID = "Cleric"
	ChampionPaladin ChampionID = "Paladin"
	ChampionRival   ChampionID = "Rival"
)

type CreatureID string

const (
	CreatureSkeleton    CreatureID = "skeleton"
	CreatureZombie      CreatureID = "zombie"
	CreatureMummy       CreatureID = "mummy"
	CreatureEyeMonster  CreatureID = "eye_monster"
	CreatureWerewolf    CreatureID = "werewolf"
	CreatureBat         CreatureID = "bat"
	CreatureShade       CreatureID = "shade"
	CreatureBanshee     CreatureID = "banshee"
	CreatureGhoul       CreatureID = "ghoul"
	CreatureImp         CreatureID = "imp"
	CreatureAbomination CreatureID = "abomination"
	CreatureDemon       CreatureID = "demon"
)

type ChampionType struct {
	ID        ChampionID
	Name      string
	Threshold int
	Creatures []CreatureID
}

// CreatureType describes a farmable creature and its summoning point values.
// MergePoints[k] = points awarded when merging 2x Lk into L(k+1),
// where L1 is created by merging 2x L2-parts (index 0).
type CreatureType struct {
	ID          CreatureID
	Name        string
	MergePoints []int
}

type PartLevel int

const (
	PartLevelOne PartLevel = 1
	PartLevelTwo PartLevel = 2
)

type StationDrop struct {
	CreatureID  CreatureID
	PartLevel   PartLevel
	Probability float64
}

type StationVariant struct {
	Level  int
	Hacked bool
	Cost   int
	Drops  []StationDrop
}

type SummoningStation struct {
	ID       StationID
	Name     string
	Resource ResourceType
	Variants []StationVariant
}
