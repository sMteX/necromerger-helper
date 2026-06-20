package prestigePlan

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// bubbletea only sends WindowSizeMsg once, before the user picks a task,
		// so the prestige planner never gets Init() called by the framework.
		// Use the first WindowSizeMsg as the init signal to activate the cursor field.
		if !m.currentInput().Focused() {
			return m, m.currentInput().Focus()
		}
		return m, nil
	case tea.KeyPressMsg:
		return m.handleKey(msg)
	}
	// Non-key messages (cursor blink ticks) must reach the active textinput.
	var cmd tea.Cmd
	m.fields[fieldIndex(m.cursor)].input, cmd = m.currentInput().Update(msg)
	return m, cmd
}

func (m *Model) debugDump() {
	log.Println("===== DEBUG DUMP =====")
	log.Printf("window: %dx%d  tab: %d  cursor: %d", m.windowWidth, m.windowHeight, m.selectedTab, m.cursor)
	log.Println("--- plan ---")
	log.Printf("  DevourerLevel=%d  FeatTiers=%d  OtherMultiplier=%.4f  GroupBonusCount=%d  LeftoverShards=%d",
		m.plan.DevourerLevel, m.plan.FeatTiers, m.plan.OtherMultiplier, m.plan.GroupBonusCount, m.plan.LeftoverShards)
	log.Println("  LegendaryCounts:")
	for k, v := range m.plan.LegendaryCounts {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  PossessedLegendaries:")
	for k, v := range m.plan.PossessedLegendaries {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  ExperimentLevels:")
	for k, v := range m.plan.ExperimentLevels {
		log.Printf("    %v = %d", k, v)
	}
	log.Println("  PossessedRunes:")
	for k, v := range m.plan.PossessedRunes {
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

func (m *Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "f12":
		m.debugDump()
		return m, nil
	case "f1", "f2", "f3", "f4":
		// blur the previous input
		m.currentInput().Blur()
		switch msg.String() {
		case "f1":
			m.selectedTab = planTabBase
			m.cursor = int(fieldBaseDevourerLevel)
		case "f2":
			m.selectedTab = planTabLegendaries
			m.cursor = int(fieldLegendariesLichHave)
		case "f3":
			m.selectedTab = planTabRunes
			m.cursor = int(fieldRunesIce)
		case "f4":
			m.selectedTab = planTabExperiments
			m.cursor = int(fieldExperimentsSeasoning1)
		}
		// focus the new input
		return m, m.currentInput().Focus()
	}

	switch m.selectedTab {
	case planTabBase:
		return m.handleBaseTabKey(msg)
	case planTabLegendaries:
		return m.handleLegendariesTabKey(msg)
	case planTabRunes:
		return m.handleRunesTabKey(msg)
	case planTabExperiments:
		return m.handleExperimentsTabKey(msg)
	}
	return m, nil
}
