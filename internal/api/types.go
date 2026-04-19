package api

import (
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

// --- Prestige planner types ---

type PlanSummary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExperimentSummary struct {
	ID                models.ExperimentID `json:"id"`
	CurrentLevel      int                 `json:"currentLevel"`
	CurrentLevelValue string              `json:"currentLevelValue"`
	CurrentLevelCost  string              `json:"currentLevelCost"`
	NextLevelCost     string              `json:"nextLevelCost"`
	NextLevelValue    string              `json:"nextLevelValue"`
	MaxLevel          bool                `json:"maxLevel"`
}

type RecalculateResponse struct {
	TotalShards      int                                         `json:"totalShards"`
	BaseShards       int                                         `json:"baseShards"`
	FeatMultiplier   float64                                     `json:"featMultiplier"`
	LegendMultiplier float64                                     `json:"legendMultiplier"`
	OtherMultiplier  float64                                     `json:"otherMultiplier"`
	ExperimentCost   int                                         `json:"experimentCost"`
	Remaining        int                                         `json:"remaining"`
	Experiments      map[models.ExperimentID]ExperimentSummary   `json:"experiments"`
	RuneTotal        map[models.RuneType]int                     `json:"runeTotal"`
	RuneNeeded       map[models.RuneType]int                     `json:"runeNeeded"`
	LegendaryRunes   map[models.LegendaryID]calculator.RuneCosts `json:"legendaryRunes"`
}

// --- Resource cap calculator types ---

// StationCounts holds how many of each station level the player owns.
type StationCounts struct {
	L1 int `json:"l1,omitempty"`
	L2 int `json:"l2,omitempty"`
	L3 int `json:"l3,omitempty"`
	L4 int `json:"l4,omitempty"`
	L5 int `json:"l5,omitempty"`
	L6 int `json:"l6,omitempty"`
}

func (s StationCounts) toSlice() [6]int {
	return [6]int{s.L1, s.L2, s.L3, s.L4, s.L5, s.L6}
}

// ServOConfig describes the player's Serv-O configuration.
type ServOConfig struct {
	Resource models.ResourceType `json:"resource"`
	Upgraded bool                `json:"upgraded"`
}

// ResourceSkins holds which skins the player has unlocked that affect resource caps.
type ResourceSkins struct {
	// Base boost skins (flat cap increase)
	Wizard bool `json:"wizard,omitempty"` // +2000 Mana
	Oozing bool `json:"oozing,omitempty"` // +2000 Slime
	Sid    bool `json:"sid,omitempty"`    // +2000 Darkness
	// Multiplicative skins
	Santa    bool `json:"santa,omitempty"`    // +5% Mana
	Birthday bool `json:"birthday,omitempty"` // +5% Slime
	Witch    bool `json:"witch,omitempty"`    // +5% Darkness
	Good     bool `json:"good,omitempty"`     // +2% All
	Royalty  bool `json:"royalty,omitempty"`  // +5% All
}

// SpellLevels holds the player's resource-related spell levels.
// Mana/Slime/Darkness are per-resource spells (+5% per level, max 5).
// AllResources is the page 6 spell (+5% per level, max 2).
type SpellLevels struct {
	Mana         int `json:"mana,omitempty"`
	Slime        int `json:"slime,omitempty"`
	Darkness     int `json:"darkness,omitempty"`
	AllResources int `json:"allResources,omitempty"`
}

// RelicLevels holds the player's equipped relic levels.
// Mana/Slime/Darkness are per-resource relics (levels 1-10).
// AllResources is the all-resources relic (levels 1-5).
type RelicLevels struct {
	Mana         int `json:"mana,omitempty"`
	Slime        int `json:"slime,omitempty"`
	Darkness     int `json:"darkness,omitempty"`
	AllResources int `json:"allResources,omitempty"`
}

// ResourceCapRequest is the input for the resource cap calculator.
type ResourceCapRequest struct {
	ManaPools  StationCounts `json:"manaPools"`
	SlimeVats  StationCounts `json:"slimeVats"`
	DarkStores StationCounts `json:"darkStores"`

	ServO *ServOConfig `json:"servO,omitempty"`

	GoldenBoosts bool          `json:"goldenBoosts,omitempty"`
	Skins        ResourceSkins `json:"skins,omitempty"`
	Spells       SpellLevels   `json:"spells,omitempty"`
	Relics       RelicLevels   `json:"relics,omitempty"`

	// CapacityExp1 is the Capacity Experiment I level (0-9).
	CapacityExp1 int `json:"capacityExp1,omitempty"`
	// PearlBonus is the per-resource pearl % bonus as shown in game statistics
	// (already includes the all-resources pearl contribution).
	PearlBonus map[models.ResourceType]float64 `json:"pearlBonus,omitempty"`
	// CapacityExp2 is the Capacity Experiment II level (0-9).
	CapacityExp2 int `json:"capacityExp2,omitempty"`

	// FixedTargets lets the player pin a specific cap goal for one resource.
	// Remaining combined threshold is then split across the other two.
	FixedTargets map[models.ResourceType]int `json:"fixedTargets,omitempty"`
}

// StationOption describes how many stations of a given level are needed to close a resource gap.
type StationOption struct {
	Level    int                     `json:"level"`
	Count    int                     `json:"count"`
	RuneCost map[models.RuneType]int `json:"runeCost"`
}

// ResourceGapAnalysis describes the gap between current and target cap for one resource type.
type ResourceGapAnalysis struct {
	Current int             `json:"current"`
	Target  int             `json:"target"`
	Gap     int             `json:"gap"`
	Options []StationOption `json:"options"` // one entry per station level (1-6)
}

// ResourceCapResponse is the output of the resource cap calculator for a single Feat threshold.
type ResourceCapResponse struct {
	Mana     int `json:"mana"`
	Slime    int `json:"slime"`
	Darkness int `json:"darkness"`
	Combined int `json:"combined"`

	Threshold int  `json:"threshold"`
	Met       bool `json:"met"`
	// Delta is positive when the player exceeds the threshold, negative when short.
	Delta int `json:"delta"`

	// GapAnalysis is only present when the threshold is not met.
	GapAnalysis map[models.ResourceType]ResourceGapAnalysis `json:"gapAnalysis,omitempty"`
}

// ValidThresholds maps the path enum value to the actual combined storage target.
var ValidThresholds = map[string]int{
	"200k": 200000,
	"400k": 400000,
	"600k": 600000,
	"800k": 800000,
}
