package resourceCap

import (
	"fmt"

	"github.com/sMteX/necromerger-helper/internal/models"
)

// Field indices.
// Shared fields (0–10) are always visible regardless of active resource.
// Per-resource fields (11–19) reflect the currently selected resource tab.
const (
	// Shared
	fieldServOResource = 0
	fieldServOUpgraded = 1
	fieldGoldenBoosts  = 2
	fieldSkinBase      = 3 // wizard / oozing / sid depending on active resource
	fieldSkinMult      = 4 // santa / birthday / witch depending on active resource
	fieldSkinGood      = 5
	fieldSkinRoyalty   = 6
	fieldCapExp1       = 7
	fieldCapExp2       = 8
	fieldSpellAll      = 9  // all-resources spell (same value across tabs)
	fieldRelicAll      = 10 // all-resources relic (same value across tabs)

	// Per-resource — ordered to match visual render order so ↑↓ navigation is sequential
	fieldSpellSelf  = 11
	fieldRelicSelf  = 12
	fieldPearlBonus = 13
	fieldStationL1  = 14
	fieldStationL2  = 15
	fieldStationL3  = 16
	fieldStationL4  = 17
	fieldStationL5  = 18
	fieldStationL6  = 19

	fieldCount = 20
)

type fieldKind int

const (
	kindNumeric fieldKind = iota
	kindToggle
	kindSelector // inline selector (Serv-O resource)
)

type fieldMeta struct {
	label string
	kind  fieldKind
}

// sharedFields describes the 11 shared fields (indices 0–10).
var sharedFields = []fieldMeta{
	fieldServOResource: {label: "Serv-O resource", kind: kindSelector},
	fieldServOUpgraded: {label: "Serv-O upgraded", kind: kindToggle},
	fieldGoldenBoosts:  {label: "Golden boosts", kind: kindToggle},
	fieldSkinBase:      {label: "Skin (+2k base)", kind: kindToggle}, // label overridden in fieldLabel
	fieldSkinMult:      {label: "Skin (+5% self)", kind: kindToggle}, // label overridden in fieldLabel
	fieldSkinGood:      {label: "Good skin (+2% all)", kind: kindToggle},
	fieldSkinRoyalty:   {label: "Royalty skin (+5% all)", kind: kindToggle},
	fieldCapExp1:       {label: "Capacity Exp I level (0-9)", kind: kindNumeric},
	fieldCapExp2:       {label: "Capacity Exp II level (0-9)", kind: kindNumeric},
	fieldSpellAll:      {label: "All-res spell level (0-2)", kind: kindNumeric},
	fieldRelicAll:      {label: "All-res relic level (0-5)", kind: kindNumeric},
}

// perResourceFields describes the 9 per-resource fields (indices 11–19).
// Indexed by idx - fieldSpellSelf. Station labels are overridden in fieldLabel().
var perResourceFields = []fieldMeta{
	0: {label: "Spell level (0-5)", kind: kindNumeric},  // fieldSpellSelf
	1: {label: "Relic level (0-10)", kind: kindNumeric}, // fieldRelicSelf
	2: {label: "Pearl bonus %", kind: kindNumeric},      // fieldPearlBonus
	3: {label: "Station L1", kind: kindNumeric},         // fieldStationL1
	4: {label: "Station L2", kind: kindNumeric},
	5: {label: "Station L3", kind: kindNumeric},
	6: {label: "Station L4", kind: kindNumeric},
	7: {label: "Station L5", kind: kindNumeric},
	8: {label: "Station L6", kind: kindNumeric}, // fieldStationL6
}

// fieldLabel returns the display label for a field, with resource-specific overrides.
func fieldLabel(idx int, res models.ResourceType) string {
	if idx < fieldSpellSelf {
		switch idx {
		case fieldSkinBase:
			switch res {
			case models.ResourceMana:
				return "Wizard skin (+2k)"
			case models.ResourceSlime:
				return "Oozing skin (+2k)"
			case models.ResourceDarkness:
				return "Sid skin (+2k)"
			}
		case fieldSkinMult:
			switch res {
			case models.ResourceMana:
				return "Santa skin (+5%)"
			case models.ResourceSlime:
				return "Birthday skin (+5%)"
			case models.ResourceDarkness:
				return "Witch skin (+5%)"
			}
		}
		return sharedFields[idx].label
	}

	// Per-resource station names (dynamic label)
	if idx >= fieldStationL1 && idx <= fieldStationL6 {
		level := idx - fieldStationL1 + 1
		switch res {
		case models.ResourceMana:
			return fmt.Sprintf("Mana Pool L%d", level)
		case models.ResourceSlime:
			return fmt.Sprintf("Slime Vat L%d", level)
		case models.ResourceDarkness:
			return fmt.Sprintf("Dark Store L%d", level)
		}
	}

	return perResourceFields[idx-fieldSpellSelf].label
}

// fieldKindOf returns the input kind for a field index.
func fieldKindOf(idx int) fieldKind {
	if idx < fieldSpellSelf {
		return sharedFields[idx].kind
	}
	return perResourceFields[idx-fieldSpellSelf].kind
}
