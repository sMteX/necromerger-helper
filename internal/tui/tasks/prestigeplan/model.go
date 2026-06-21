package prestigeplan

import (
	"log"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/persistence"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/planmenu"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/base"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/experiments"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/legendaries"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/tabs/runes"
)

type planTab int8

const (
	planTabBase planTab = iota
	planTabLegendaries
	planTabRunes
	planTabExperiments
)

// clearSaveStatusMsg is sent by savePlan's timer goroutine to clear the
// "Saved" toast from the header after a short delay.
type clearSaveStatusMsg struct{}

type Model struct {
	selectedTab               planTab
	windowHeight, windowWidth int

	baseTab        *base.Model
	legendariesTab *legendaries.Model
	runesTab       *runes.Model
	experimentsTab *experiments.Model

	// result is allocated once and shared with the tab sub-models via pointer.
	// recalculate() writes through it so tabs always read fresh data.
	result *calculator.PrestigePlanResult

	planName   string
	planNotes  string
	planPath   string // empty string means the plan has never been saved
	planDirty  bool   // true = unsaved changes exist; shown as "*" in the header
	saveStatus string // short-lived toast text (e.g. "Saved"); cleared by timer
	// loadedPlan holds the last SavedPlan read from or written to disk.
	// It exists solely to preserve CreatedAt when re-saving — the domain model
	// (models.Plan) doesn't carry metadata timestamps.
	loadedPlan *persistence.SavedPlan

	plansDir string          // resolved once in New(); used when listing/saving plans
	menu     *planmenu.Model // non-nil while the plan menu modal is open
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) recalculate() {
	*m.result = calculator.Calculate(m.assemblePlan())
}

func (m *Model) assemblePlan() models.Plan {
	return models.Plan{
		Name:                 m.planName,
		Notes:                m.planNotes,
		DevourerLevel:        m.baseTab.DevourerLevel,
		FeatTiers:            m.baseTab.FeatTiers,
		OtherMultiplier:      m.baseTab.OtherMultiplier,
		GroupBonusCount:      m.baseTab.GroupBonusCount,
		LeftoverShards:       m.baseTab.LeftoverShards,
		LegendaryCounts:      m.legendariesTab.LegendaryCounts,
		PossessedLegendaries: m.legendariesTab.PossessedLegendaries,
		PossessedRunes:       m.runesTab.PossessedRunes,
		ExperimentLevels:     m.experimentsTab.ExperimentLevels,
	}
}

// savePlan writes the current plan to disk, updates plan metadata, and returns
// a command that clears the "Saved" toast after 2 seconds. The 2-second sleep
// runs inside the command function, which bubbletea executes in a goroutine —
// it does not block the TUI event loop.
func (m *Model) savePlan(name, notes, path string) tea.Cmd {
	m.planName = name
	m.planNotes = notes
	saved := persistence.FromModels(m.assemblePlan(), m.loadedPlan)
	if err := persistence.Save(path, saved); err != nil {
		log.Printf("save error: %v", err)
		return nil
	}
	m.planPath = path
	m.planDirty = false
	m.loadedPlan = &saved
	m.saveStatus = "Saved"
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		return clearSaveStatusMsg{}
	}
}

// loadPlan reads a plan from disk and pushes its values into all four tab
// sub-models, then recalculates and re-focuses the current input.
func (m *Model) loadPlan(path string) tea.Cmd {
	saved, err := persistence.Load(path)
	if err != nil {
		log.Printf("load error: %v", err)
		return nil
	}
	plan := saved.ToModels()
	m.planName = plan.Name
	m.planNotes = plan.Notes
	m.planPath = path
	m.planDirty = false
	m.loadedPlan = &saved
	m.baseTab.LoadFrom(plan)
	m.legendariesTab.LoadFrom(plan)
	m.runesTab.LoadFrom(plan)
	m.experimentsTab.LoadFrom(plan)
	m.recalculate()
	return m.currentInput().Focus()
}

// newPlan resets the entire model to a blank state. A fresh result pointer is
// allocated so all tab sub-models share the same new pointer — the old one is
// abandoned to the GC.
func (m *Model) newPlan() tea.Cmd {
	resultPtr := new(calculator.PrestigePlanResult)
	m.result = resultPtr
	m.baseTab = base.NewModel()
	m.legendariesTab = legendaries.NewModel(resultPtr)
	m.runesTab = runes.NewModel(resultPtr)
	m.experimentsTab = experiments.NewModel()
	m.planName = ""
	m.planNotes = ""
	m.planPath = ""
	m.planDirty = false
	m.loadedPlan = nil
	m.recalculate()
	return m.currentInput().Focus()
}

func New() *Model {
	resultPtr := new(calculator.PrestigePlanResult)
	plansDir, _ := persistence.DefaultPlansDir()
	m := &Model{
		selectedTab:    planTabExperiments,
		result:         resultPtr,
		baseTab:        base.NewModel(),
		legendariesTab: legendaries.NewModel(resultPtr),
		runesTab:       runes.NewModel(resultPtr),
		experimentsTab: experiments.NewModel(),
		plansDir:       plansDir,
	}
	m.recalculate()
	return m
}

// currentInput returns the current input model for the currently selected tab
func (m *Model) currentInput() *textinput.Model {
	switch m.selectedTab {
	case planTabBase:
		return m.baseTab.CurrentInput()
	case planTabLegendaries:
		return m.legendariesTab.CurrentInput()
	case planTabRunes:
		return m.runesTab.CurrentInput()
	case planTabExperiments:
		return m.experimentsTab.CurrentInput()
	}
	return nil
}
