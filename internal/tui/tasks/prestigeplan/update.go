package prestigeplan

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/persistence"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigeplan/planmenu"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// When the plan menu is open, it gets all messages exclusively. Once the
	// modal sets a result, we clear it and handle the action.
	if m.menu != nil {
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		if result := m.menu.Result(); result != nil {
			m.menu = nil
			return m, tea.Batch(cmd, m.handleMenuResult(result))
		}
		return m, cmd
	}

	// Timer fired — clear the "Saved" toast from the header.
	if _, ok := msg.(clearSaveStatusMsg); ok {
		m.saveStatus = ""
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// bubbletea only sends WindowSizeMsg once, before the user picks a task,
		// so the prestige planner never gets Init() called by the framework.
		// Use the first WindowSizeMsg as the init signal to activate the cursor field.
		// TODO: not sure here, if this should work after we put back the resource cap planner too
		if !m.currentInput().Focused() {
			return m, m.currentInput().Focus()
		}
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "f12":
			m.debugDump()
			return m, nil
		case "x":
			return m, m.openMenu(false)
		case "ctrl+s":
			return m, m.quickSave()
		case "f1", "f2", "f3", "f4":
			// blur the previous input
			m.currentInput().Blur()
			switch msg.String() {
			case "f1":
				m.selectedTab = planTabBase
			case "f2":
				m.selectedTab = planTabLegendaries
			case "f3":
				m.selectedTab = planTabRunes
			case "f4":
				m.selectedTab = planTabExperiments
			}
			// focus the new input - will refer to something else after switching the tab
			return m, m.currentInput().Focus()
		}
	}
	return m.handleTabUpdates(msg)
}

// openMenu blurs the current input, fetches the latest plan list from disk, and
// opens the plan menu modal. startAtSave skips the top-level menu and goes
// directly to the name editor (used by ctrl+s when no path exists yet).
func (m *Model) openMenu(startAtSave bool) tea.Cmd {
	m.currentInput().Blur()
	plans, _ := persistence.ListPlans(m.plansDir)
	m.menu = planmenu.New(planmenu.Config{
		PlansDir:    m.plansDir,
		Plans:       plans,
		Dirty:       m.planDirty,
		StartAtSave: startAtSave,
		Name:        m.planName,
		Notes:       m.planNotes,
	})
	if startAtSave {
		// Focus is handled inside planmenu.New when StartAtSave is true —
		// no cmd needed from here.
		return nil
	}
	return nil
}

// quickSave saves immediately if the plan already has a path (i.e. it has been
// named before), otherwise opens the name editor to get one.
func (m *Model) quickSave() tea.Cmd {
	if m.planPath != "" {
		return m.savePlan(m.planName, m.planNotes, m.planPath)
	}
	return m.openMenu(true)
}

func (m *Model) handleMenuResult(result *planmenu.Action) tea.Cmd {
	switch result.Type {
	case planmenu.ActionNew:
		return m.newPlan()
	case planmenu.ActionLoad:
		return m.loadPlan(result.Path)
	case planmenu.ActionSave:
		// Batch the save (which returns a timer cmd) with re-focusing the input.
		return tea.Batch(m.savePlan(result.Name, result.Notes, result.Path), m.currentInput().Focus())
	default:
		// Modal was dismissed without action — just restore focus.
		return m.currentInput().Focus()
	}
}

func (m *Model) debugDump() {
	log.Println("===== DEBUG DUMP =====")
	log.Printf("window: %dx%d  tab: %d", m.windowWidth, m.windowHeight, m.selectedTab)
	log.Println("--- plan ---")
	log.Printf("  DevourerLevel=%d  FeatTiers=%d  OtherMultiplier=%.4f  GroupBonusCount=%d  LeftoverShards=%d",
		m.baseTab.DevourerLevel, m.baseTab.FeatTiers, m.baseTab.OtherMultiplier, m.baseTab.GroupBonusCount, m.baseTab.LeftoverShards)
	log.Println("  LegendaryCounts:")
	for k, v := range m.legendariesTab.LegendaryCounts {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  PossessedLegendaries:")
	for k, v := range m.legendariesTab.PossessedLegendaries {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  ExperimentLevels:")
	for k, v := range m.experimentsTab.ExperimentLevels {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  PossessedRunes:")
	for k, v := range m.runesTab.PossessedRunes {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("--- result ---")
	log.Printf("  BaseShards=%d  TotalShards=%d  ExperimentCost=%d  NetShards=%d",
		m.result.BaseShards, m.result.TotalShards, m.result.ExperimentCost, m.result.NetShards)
	log.Printf("  FeatMultiplier=%.6f  LegendMultiplier=%.6f  OtherMultiplier=%.6f",
		m.result.FeatMultiplier, m.result.LegendMultiplier, m.result.OtherMultiplier)
	log.Printf("  RuneTotal=%v", m.result.RuneTotal)
	log.Printf("  RuneNeeded=%v", m.result.RuneNeeded)
	log.Println("  LegendaryBonuses:")
	for k, v := range m.result.LegendaryBonuses {
		log.Printf("    %v = %.4f", k, v)
	}
	log.Println("  LegendaryGroupBonuses:")
	for k, v := range m.result.LegendaryGroupBonuses {
		log.Printf("    %v = %.4f", k, v)
	}
	log.Println("--- rendered summary values ---")
	const summaryW = 35
	fw := shared.Styles.SubContainer.GetHorizontalFrameSize()
	labelW := summaryW - fw - 13
	log.Printf("  summaryWidth=%d subContainerFW=%d => labelWidth=%d valueWidth=13", summaryW, fw, labelW)
	base := shared.FormatNumberLong(m.result.BaseShards)
	total := shared.FormatNumberLong(m.result.TotalShards)
	spent := shared.FormatNumberLong(-m.result.ExperimentCost)
	net := shared.FormatNumberLong(m.result.NetShards)
	log.Printf("  Base=%q (%d chars)", base, len([]rune(base)))
	log.Printf("  Total=%q (%d chars)", total, len([]rune(total)))
	log.Printf("  Spent=%q (%d chars)", spent, len([]rune(spent)))
	log.Printf("  Net=%q (%d chars)", net, len([]rune(net)))
	// Render the net row as view.go does and measure its visual width.
	valueStyle := lipgloss.NewStyle().Width(13).AlignHorizontal(lipgloss.Right)
	labelStyle := lipgloss.NewStyle().Width(labelW)
	netColor := shared.Colors.Good
	if m.result.NetShards <= 0 {
		netColor = shared.Colors.Bad
	}
	netStyle := lipgloss.NewStyle().Width(13).AlignHorizontal(lipgloss.Right).Foreground(netColor)
	netLabel := labelStyle.Render("Net")
	netVal := netStyle.Render(net)
	netRow := lipgloss.JoinHorizontal(lipgloss.Top, netLabel, netVal)
	log.Printf("  net label visual width: %d", lipgloss.Width(netLabel))
	log.Printf("  net value visual width: %d", lipgloss.Width(netVal))
	log.Printf("  net row visual width: %d (expected %d)", lipgloss.Width(netRow), labelW+13)
	spentVal := valueStyle.Foreground(shared.Colors.Bad).Render(spent)
	log.Printf("  spent value visual width: %d", lipgloss.Width(spentVal))
	log.Println("======================")
}

func (m *Model) handleTabUpdates(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.selectedTab {
	case planTabBase:
		m.baseTab, cmd = m.baseTab.Update(msg)
	case planTabLegendaries:
		m.legendariesTab, cmd = m.legendariesTab.Update(msg)
	case planTabRunes:
		m.runesTab, cmd = m.runesTab.Update(msg)
	case planTabExperiments:
		m.experimentsTab, cmd = m.experimentsTab.Update(msg)
	}
	m.recalculate()
	if key, isKey := msg.(tea.KeyPressMsg); isKey {
		switch key.String() {
		case "up", "down", "tab", "shift+tab":
			// Pure navigation — cursor moves between fields but no value changes.
		default:
			// Any other key (character input, backspace, left/right step) may
			// have changed a value, so mark the plan dirty.
			m.planDirty = true
		}
	}
	return m, cmd
}
