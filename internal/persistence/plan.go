package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/sMteX/necro-prestige-planner/internal/models"
)

// SavedPlan is the on-disk JSON representation of a plan. It is intentionally
// separate from models.Plan so we can version it independently and add metadata
// (CreatedAt, UpdatedAt) that the domain model doesn't care about.
type SavedPlan struct {
	Name      string    `json:"name"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DevourerLevel        int                         `json:"devourer_level"`
	FeatTiers            int                         `json:"feat_tiers"`
	OtherMultiplier      float64                     `json:"other_multiplier"`
	GroupBonusCount      int                         `json:"group_bonus_count"`
	LeftoverShards       int                         `json:"leftover_shards"`
	LegendaryCounts      map[models.LegendaryID]int  `json:"legendary_counts"`
	PossessedLegendaries map[models.LegendaryID]int  `json:"possessed_legendaries"`
	PossessedRunes       map[models.RuneType]int     `json:"possessed_runes"`
	ExperimentLevels     map[models.ExperimentID]int `json:"experiment_levels"`
}

// PlanMeta is a lightweight summary used to populate the load list in the UI
// without keeping full plan data in memory.
type PlanMeta struct {
	Name      string
	Notes     string
	UpdatedAt time.Time
	Path      string // absolute path to the .json file
}

// DefaultPlansDir returns ~/.config/necro-prestige-planner/plans/ (or the OS equivalent).
func DefaultPlansDir() (string, error) {
	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfg, "necro-prestige-planner", "plans"), nil
}

// Save writes the plan to path as pretty-printed JSON, creating parent directories
// as needed and stamping UpdatedAt with the current time.
func Save(path string, plan SavedPlan) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	plan.UpdatedAt = time.Now()
	data, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Load(path string) (SavedPlan, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SavedPlan{}, err
	}
	var plan SavedPlan
	return plan, json.Unmarshal(data, &plan)
}

// ListPlans reads all .json files in dir, loads their metadata, and returns them
// sorted newest-first. Returns nil (not an error) if the directory doesn't exist yet.
func ListPlans(dir string) ([]PlanMeta, error) {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		// Plans directory hasn't been created yet — treat as empty list.
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var plans []PlanMeta
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		plan, err := Load(path)
		if err != nil {
			// Skip unreadable or malformed files silently.
			continue
		}
		plans = append(plans, PlanMeta{
			Name:      plan.Name,
			Notes:     plan.Notes,
			UpdatedAt: plan.UpdatedAt,
			Path:      path,
		})
	}
	sort.Slice(plans, func(i, j int) bool {
		return plans[i].UpdatedAt.After(plans[j].UpdatedAt)
	})
	return plans, nil
}

// FromModels converts a domain Plan to a SavedPlan. If existing is non-nil its
// CreatedAt is carried forward so round-tripping a save doesn't reset the creation date.
func FromModels(plan models.Plan, existing *SavedPlan) SavedPlan {
	createdAt := time.Now()
	if existing != nil {
		createdAt = existing.CreatedAt
	}
	return SavedPlan{
		Name:                 plan.Name,
		Notes:                plan.Notes,
		CreatedAt:            createdAt,
		DevourerLevel:        plan.DevourerLevel,
		FeatTiers:            plan.FeatTiers,
		OtherMultiplier:      plan.OtherMultiplier,
		GroupBonusCount:      plan.GroupBonusCount,
		LeftoverShards:       plan.LeftoverShards,
		LegendaryCounts:      plan.LegendaryCounts,
		PossessedLegendaries: plan.PossessedLegendaries,
		PossessedRunes:       plan.PossessedRunes,
		ExperimentLevels:     plan.ExperimentLevels,
	}
}

func (s SavedPlan) ToModels() models.Plan {
	return models.Plan{
		Name:                 s.Name,
		Notes:                s.Notes,
		DevourerLevel:        s.DevourerLevel,
		FeatTiers:            s.FeatTiers,
		OtherMultiplier:      s.OtherMultiplier,
		GroupBonusCount:      s.GroupBonusCount,
		LeftoverShards:       s.LeftoverShards,
		LegendaryCounts:      s.LegendaryCounts,
		PossessedLegendaries: s.PossessedLegendaries,
		PossessedRunes:       s.PossessedRunes,
		ExperimentLevels:     s.ExperimentLevels,
	}
}

// PlanFileName converts a plan name to a safe .json filename inside dir.
// Spaces become underscores; non-alphanumeric characters (except _ and -) are dropped.
func PlanFileName(dir, name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "_")
	var out []byte
	for _, c := range []byte(s) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' || c == '-' {
			out = append(out, c)
		}
	}
	if len(out) == 0 {
		out = []byte("plan")
	}
	return filepath.Join(dir, string(out)+".json")
}
