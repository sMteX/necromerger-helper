package calculator

import (
	"testing"

	"github.com/sMteX/necromerger-helper/internal/models"
)

func TestCalculateTotalRunes(t *testing.T) {
	plan := models.Plan{
		LegendaryCounts: map[models.LegendaryID]int{
			models.Lich: 1, // Target: 1 Lich
		},
		PossessedRunes: map[models.RuneType]int{
			models.RuneIce: 100, // Have 100 Ice
		},
	}

	// 1 Lich = 31 * 20 Ice = 620 Ice
	// Total: 620 Ice
	// Needed: 620 - 100 = 520 Ice
	total, needed := CalculateTotalRunes(plan)

	if total[models.RuneIce] != 620 {
		t.Errorf("Expected 620 total ice, got %d", total[models.RuneIce])
	}
	if needed[models.RuneIce] != 520 {
		t.Errorf("Expected 520 needed ice, got %d", needed[models.RuneIce])
	}

	// Add Cursed to target
	plan.LegendaryCounts[models.TheCursed] = 1
	// 1 Cursed requires 1 Lich and 1 Reaper
	// Total targets: 2 Liches, 1 Reaper, 1 Cursed
	// Total station-produced: 2 Liches, 1 Reaper
	// 2 Liches = 2 * 620 Ice = 1240 Ice
	// 1 Reaper = 31 * (50 Ice + 20 Moon) = 1550 Ice, 620 Moon
	// Total Ice: 1240 + 1550 = 2790
	// Total Moon: 620
	total, needed = CalculateTotalRunes(plan)

	if total[models.RuneIce] != 2790 {
		t.Errorf("Expected 2790 total ice, got %d", total[models.RuneIce])
	}
	if total[models.RuneMoon] != 620 {
		t.Errorf("Expected 620 total moon, got %d", total[models.RuneMoon])
	}

	// Add possessed Lich
	plan.PossessedLegendaries = map[models.LegendaryID]int{
		models.Lich: 1,
	}
	// TotalNeeded: 2 Liches, 1 Reaper
	// TotalHave: 1 Lich
	// ToBuy: 1 Lich, 1 Reaper
	// Needed Ice: (1*620 + 1*1550) - 100 = 2170 - 100 = 2070
	_, needed = CalculateTotalRunes(plan)
	if needed[models.RuneIce] != 2070 {
		t.Errorf("Expected 2070 needed ice, got %d", needed[models.RuneIce])
	}
}
