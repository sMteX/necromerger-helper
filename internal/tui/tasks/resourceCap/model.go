package resourceCap

import (
	"strconv"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necromerger-helper/internal/calculator"
	"github.com/sMteX/necromerger-helper/internal/models"
)

// servOResources is the cycle order for the Serv-O resource selector.
// No "none" — Serv-O always has a resource assigned.
var servOResources = []models.ResourceType{models.ResourceMana, models.ResourceSlime, models.ResourceDarkness}
var servOLabels = []string{"Mana", "Slime", "Darkness"}

// threshold options in cycle order (t key cycles through them).
var thresholds = []int{200_000, 400_000, 600_000, 800_000}
var thresholdLabels = []string{"200k", "400k", "600k", "800k"}

// Model is the root bubbletea model for the TUI.
type Model struct {
	width, height int
	leftWidth     int
	rightWidth    int

	// Which panel has keyboard focus: 0 = left (inputs), 1 = right (output)
	focusedPanel int

	// Currently active resource tab (for per-resource fields)
	activeResource models.ResourceType

	// Selected threshold index (into thresholds slice)
	thresholdIdx int

	// Focused field index within the left panel
	focusedField int

	// Shared toggles / selectors
	servOResourceIdx int // index into servOResources
	servOUpgraded    bool
	goldenBoosts     bool
	skinWizard       bool
	skinOozing       bool
	skinSid          bool
	skinSanta        bool
	skinBirthday     bool
	skinWitch        bool
	skinGood         bool
	skinRoyalty      bool
	capExp1          int
	capExp2          int

	// Shared spell/relic all-resources values (single value, not per-tab)
	spellAll int
	relicAll int

	// Per-resource station counts [mana/slime/darkness][L1..L6]
	stations [3][6]int

	// Per-resource self spells and relics
	spellSelf [3]int
	relicSelf [3]int

	// Per-resource pearl bonus as integer % (e.g. 6 means 6%)
	pearlBonus [3]int

	// Single textinput reused across numeric fields
	input textinput.Model

	// Right-panel viewport
	vp viewport.Model

	// Last calculated result (recalculated on every change)
	result calculator.ResourceCapResult
}

func New() Model {
	inp := textinput.New()
	inp.SetWidth(8)
	inp.Prompt = ""
	inp.CharLimit = 6

	m := Model{
		activeResource:   models.ResourceMana,
		servOResourceIdx: 0, // defaults to Mana
		thresholdIdx:     0,
		focusedField:     0,
		input:            inp,
	}

	m.recalculate()
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

// resourceIdx returns 0/1/2 for mana/slime/darkness.
func resourceIdx(r models.ResourceType) int {
	switch r {
	case models.ResourceSlime:
		return 1
	case models.ResourceDarkness:
		return 2
	default:
		return 0
	}
}

func (m *Model) recalculate() {
	m.result = calculator.CalculateResourceCaps(m.buildCalculatorInput())
}

func (m Model) buildCalculatorInput() calculator.ResourceCapInput {
	return calculator.ResourceCapInput{
		ManaPools:  m.stations[0],
		SlimeVats:  m.stations[1],
		DarkStores: m.stations[2],

		ServOResource: servOResources[m.servOResourceIdx],
		ServOUpgraded: m.servOUpgraded,

		WizardSkin: m.skinWizard,
		OozingSkin: m.skinOozing,
		SidSkin:    m.skinSid,

		SantaSkin:    m.skinSanta,
		BirthdaySkin: m.skinBirthday,
		WitchSkin:    m.skinWitch,
		GoodSkin:     m.skinGood,
		RoyaltySkin:  m.skinRoyalty,

		GoldenBoosts: m.goldenBoosts,

		ManaSpell:         m.spellSelf[0],
		SlimeSpell:        m.spellSelf[1],
		DarknessSpell:     m.spellSelf[2],
		AllResourcesSpell: m.spellAll,

		ManaRelic:         m.relicSelf[0],
		SlimeRelic:        m.relicSelf[1],
		DarknessRelic:     m.relicSelf[2],
		AllResourcesRelic: m.relicAll,

		CapacityExp1: m.capExp1,
		PearlBonus: map[models.ResourceType]float64{
			models.ResourceMana:     float64(m.pearlBonus[0]) / 100.0,
			models.ResourceSlime:    float64(m.pearlBonus[1]) / 100.0,
			models.ResourceDarkness: float64(m.pearlBonus[2]) / 100.0,
		},
		CapacityExp2: m.capExp2,
	}
}

// currentNumericValue returns the string representation of the focused numeric field.
func (m Model) currentNumericValue() string {
	return m.currentNumericValueFor(m.focusedField)
}

// currentNumericValueFor returns the display string for any numeric field.
func (m Model) currentNumericValueFor(idx int) string {
	ri := resourceIdx(m.activeResource)
	switch idx {
	case fieldCapExp1:
		return strconv.Itoa(m.capExp1)
	case fieldCapExp2:
		return strconv.Itoa(m.capExp2)
	case fieldSpellAll:
		return strconv.Itoa(m.spellAll)
	case fieldRelicAll:
		return strconv.Itoa(m.relicAll)
	case fieldStationL1:
		return strconv.Itoa(m.stations[ri][0])
	case fieldStationL2:
		return strconv.Itoa(m.stations[ri][1])
	case fieldStationL3:
		return strconv.Itoa(m.stations[ri][2])
	case fieldStationL4:
		return strconv.Itoa(m.stations[ri][3])
	case fieldStationL5:
		return strconv.Itoa(m.stations[ri][4])
	case fieldStationL6:
		return strconv.Itoa(m.stations[ri][5])
	case fieldSpellSelf:
		return strconv.Itoa(m.spellSelf[ri])
	case fieldRelicSelf:
		return strconv.Itoa(m.relicSelf[ri])
	case fieldPearlBonus:
		return strconv.Itoa(m.pearlBonus[ri])
	}
	return ""
}

// commitNumericInput parses the textinput value and writes it back to model state.
func (m *Model) commitNumericInput() {
	ri := resourceIdx(m.activeResource)
	raw := m.input.Value()

	asInt := func() int {
		v, err := strconv.Atoi(raw)
		if err != nil || v < 0 {
			return 0
		}
		return v
	}

	switch m.focusedField {
	case fieldCapExp1:
		v := asInt()
		if v > 9 {
			v = 9
		}
		m.capExp1 = v
	case fieldCapExp2:
		v := asInt()
		if v > 9 {
			v = 9
		}
		m.capExp2 = v
	case fieldSpellAll:
		v := asInt()
		if v > 2 {
			v = 2
		}
		m.spellAll = v
	case fieldRelicAll:
		v := asInt()
		if v > 5 {
			v = 5
		}
		m.relicAll = v
	case fieldStationL1:
		m.stations[ri][0] = asInt()
	case fieldStationL2:
		m.stations[ri][1] = asInt()
	case fieldStationL3:
		m.stations[ri][2] = asInt()
	case fieldStationL4:
		m.stations[ri][3] = asInt()
	case fieldStationL5:
		m.stations[ri][4] = asInt()
	case fieldStationL6:
		m.stations[ri][5] = asInt()
	case fieldSpellSelf:
		v := asInt()
		if v > 5 {
			v = 5
		}
		m.spellSelf[ri] = v
	case fieldRelicSelf:
		v := asInt()
		if v > 10 {
			v = 10
		}
		m.relicSelf[ri] = v
	case fieldPearlBonus:
		m.pearlBonus[ri] = asInt()
	}
}
