package resourceCap

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

func (m Model) View() tea.View {
	if m.width == 0 {
		return tea.View{Content: "Loading…"}
	}

	leftContent := m.renderLeft()
	rightContent := m.vp.View()

	leftStyle := stylePanelBorder.Width(m.leftWidth).Height(m.height - 2)
	rightStyle := stylePanelBorder.Width(m.rightWidth).Height(m.height - 2)
	if m.focusedPanel == 0 {
		leftStyle = stylePanelBorderFocused.Width(m.leftWidth).Height(m.height - 2)
	} else {
		rightStyle = stylePanelBorderFocused.Width(m.rightWidth).Height(m.height - 2)
	}

	left := leftStyle.Render(leftContent)
	right := rightStyle.Render(rightContent)

	return tea.View{
		Content:   lipgloss.JoinHorizontal(lipgloss.Top, left, right),
		AltScreen: true,
	}
}

// ─── Left panel ──────────────────────────────────────────────────────────────

func (m Model) renderLeft() string {
	var b strings.Builder

	// Threshold selector (t cycles)
	b.WriteString(styleSectionHeader.Render("Threshold") + "\n")
	b.WriteString(m.renderThresholdBar())
	b.WriteString("\n\n")

	// Resource tab bar
	b.WriteString(styleSectionHeader.Render("Resource") + "\n")
	b.WriteString(m.renderResourceBar())
	b.WriteString("\n\n")

	// Shared fields (0–10)
	b.WriteString(styleSectionHeader.Render("Shared") + "\n")
	for i := 0; i < fieldSpellSelf; i++ {
		b.WriteString(m.renderField(i))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Per-resource: header colour follows active resource
	resHeader := lipgloss.NewStyle().Bold(true).Foreground(resourceColor(m.activeResource))

	stationSectionName := map[models.ResourceType]string{
		models.ResourceMana:     "Mana Pools",
		models.ResourceSlime:    "Slime Vats",
		models.ResourceDarkness: "Dark Stores",
	}[m.activeResource]

	// Spells & Relics (self) sub-section
	b.WriteString(resHeader.Render("Spells & Relics") + "\n")
	b.WriteString(m.renderField(fieldSpellSelf))
	b.WriteString("\n")
	b.WriteString(m.renderField(fieldRelicSelf))
	b.WriteString("\n\n")

	// Pearl sub-section
	b.WriteString(resHeader.Render("Pearl") + "\n")
	b.WriteString(m.renderField(fieldPearlBonus))
	b.WriteString("\n\n")

	// Stations sub-section
	b.WriteString(resHeader.Render(stationSectionName) + "\n")
	for i := fieldStationL1; i <= fieldStationL6; i++ {
		b.WriteString(m.renderField(i))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styleLabel.Render("t=threshold  m/s/d=resource  ↑↓=navigate  tab=switch"))

	return b.String()
}

func (m Model) renderThresholdBar() string {
	parts := make([]string, len(thresholdLabels))
	for i, lbl := range thresholdLabels {
		if i == m.thresholdIdx {
			parts[i] = styleThresholdSelected.Render(lbl)
		} else {
			parts[i] = styleThresholdOption.Render(lbl)
		}
	}
	return strings.Join(parts, "")
}

func (m Model) renderResourceBar() string {
	resources := []struct {
		res   models.ResourceType
		label string
	}{
		{models.ResourceMana, "[m]ana"},
		{models.ResourceSlime, "[s]lime"},
		{models.ResourceDarkness, "[d]arkness"},
	}

	parts := make([]string, len(resources))
	for i, t := range resources {
		if t.res == m.activeResource {
			parts[i] = styleResourceTabActive.Foreground(resourceColor(t.res)).Render(t.label)
		} else {
			parts[i] = styleResourceTab.Render(t.label)
		}
	}
	return strings.Join(parts, "  ")
}

func (m Model) renderField(idx int) string {
	focused := m.focusedPanel == 0 && m.focusedField == idx
	lbl := fieldLabel(idx, m.activeResource)
	kind := fieldKindOf(idx)

	var labelStyle lipgloss.Style
	switch {
	case focused:
		labelStyle = styleLabelFocused
	case idx >= fieldSpellSelf:
		labelStyle = resourceStyle(m.activeResource)
	default:
		labelStyle = styleLabel
	}

	switch kind {
	case kindNumeric:
		var val string
		if focused {
			val = m.input.View()
		} else {
			val = styleValue.Render(m.currentNumericValueFor(idx))
		}
		return fmt.Sprintf("  %s: %s", labelStyle.Render(lbl), val)

	case kindToggle:
		checked := m.toggleValueFor(idx)
		box := "[ ]"
		if checked {
			box = styleMetYes.Render("[x]")
		}
		return fmt.Sprintf("  %s %s", box, labelStyle.Render(lbl))

	case kindSelector:
		return fmt.Sprintf("  %s: %s", labelStyle.Render(lbl), m.renderServOSelector(focused))
	}
	return ""
}

func (m Model) renderServOSelector(focused bool) string {
	parts := make([]string, len(servOLabels))
	for i, lbl := range servOLabels {
		if i == m.servOResourceIdx {
			parts[i] = styleThresholdSelected.Render(lbl)
		} else if focused {
			parts[i] = styleThresholdOption.Render(lbl)
		} else {
			parts[i] = styleLabel.Render(lbl)
		}
	}
	hint := ""
	if focused {
		hint = styleLabel.Render(" ←→")
	}
	return strings.Join(parts, " ") + hint
}

func (m Model) toggleValueFor(idx int) bool {
	switch idx {
	case fieldServOUpgraded:
		return m.servOUpgraded
	case fieldGoldenBoosts:
		return m.goldenBoosts
	case fieldSkinBase:
		switch m.activeResource {
		case models.ResourceMana:
			return m.skinWizard
		case models.ResourceSlime:
			return m.skinOozing
		case models.ResourceDarkness:
			return m.skinSid
		}
	case fieldSkinMult:
		switch m.activeResource {
		case models.ResourceMana:
			return m.skinSanta
		case models.ResourceSlime:
			return m.skinBirthday
		case models.ResourceDarkness:
			return m.skinWitch
		}
	case fieldSkinGood:
		return m.skinGood
	case fieldSkinRoyalty:
		return m.skinRoyalty
	}
	return false
}

// ─── Right panel ─────────────────────────────────────────────────────────────

func (m Model) renderOutput() string {
	var b strings.Builder

	threshold := thresholds[m.thresholdIdx]
	combined := m.result.Mana + m.result.Slime + m.result.Darkness
	delta := combined - threshold

	// Resource caps summary
	b.WriteString(styleSectionHeader.Render("Resource caps") + "\n\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n",
		resourceStyle(models.ResourceMana).Render("Mana:"),
		styleValue.Render(shared.FormatNumberLong(m.result.Mana))))
	b.WriteString(fmt.Sprintf("  %s  %s\n",
		resourceStyle(models.ResourceSlime).Render("Slime:"),
		styleValue.Render(shared.FormatNumberLong(m.result.Slime))))
	b.WriteString(fmt.Sprintf("  %s  %s\n",
		resourceStyle(models.ResourceDarkness).Render("Darkness:"),
		styleValue.Render(shared.FormatNumberLong(m.result.Darkness))))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  Combined:  %s\n", styleValue.Render(shared.FormatNumberLong(combined))))

	// Threshold with inline gap delta
	var deltaStyled string
	if delta >= 0 {
		deltaStyled = styleMetYes.Render("+" + shared.FormatNumberLong(delta))
	} else {
		deltaStyled = styleMetNo.Render(shared.FormatNumberLong(delta))
	}
	b.WriteString(fmt.Sprintf("  Threshold: %s  %s\n",
		styleValue.Render(thresholdLabels[m.thresholdIdx]),
		deltaStyled))

	b.WriteString("\n")
	b.WriteString(styleSectionHeader.Render("Gap analysis") + "\n\n")

	// Compute targets for all resources
	manaT, slimeT, darkT := calculator.ResourceTargets(threshold, nil)
	targets := map[models.ResourceType]int{
		models.ResourceMana:     manaT,
		models.ResourceSlime:    slimeT,
		models.ResourceDarkness: darkT,
	}
	caps := map[models.ResourceType]int{
		models.ResourceMana:     m.result.Mana,
		models.ResourceSlime:    m.result.Slime,
		models.ResourceDarkness: m.result.Darkness,
	}
	multis := map[models.ResourceType]float64{
		models.ResourceMana:     m.result.ManaMulti,
		models.ResourceSlime:    m.result.SlimeMulti,
		models.ResourceDarkness: m.result.DarkMulti,
	}

	// Three-column layout
	colW := m.rightWidth / 3
	if colW < 22 {
		colW = 22
	}
	colStyle := lipgloss.NewStyle().Width(colW)

	col1 := m.renderResourceColumn(models.ResourceMana, targets, caps, multis)
	col2 := m.renderResourceColumn(models.ResourceSlime, targets, caps, multis)
	col3 := m.renderResourceColumn(models.ResourceDarkness, targets, caps, multis)

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top,
		colStyle.Render(col1),
		colStyle.Render(col2),
		colStyle.Render(col3),
	))

	// Effective multipliers (single line to save space)
	b.WriteString("\n")
	b.WriteString(styleSectionHeader.Render("Multipliers") + "\n")
	b.WriteString(fmt.Sprintf("  %s %.4f   %s %.4f   %s %.4f\n",
		resourceStyle(models.ResourceMana).Render("Mana"),
		m.result.ManaMulti,
		resourceStyle(models.ResourceSlime).Render("Slime"),
		m.result.SlimeMulti,
		resourceStyle(models.ResourceDarkness).Render("Dark"),
		m.result.DarkMulti,
	))

	return b.String()
}

// renderResourceColumn renders one resource's gap analysis as a vertical string.
func (m Model) renderResourceColumn(
	res models.ResourceType,
	targets, caps map[models.ResourceType]int,
	multis map[models.ResourceType]float64,
) string {
	var b strings.Builder

	cur := caps[res]
	tgt := targets[res]
	gap := tgt - cur

	rs := resourceStyle(res)
	label := map[models.ResourceType]string{
		models.ResourceMana:     "Mana",
		models.ResourceSlime:    "Slime",
		models.ResourceDarkness: "Darkness",
	}[res]

	b.WriteString(rs.Bold(true).Render(label) + "\n")
	b.WriteString(fmt.Sprintf("cur  %s\n", styleValue.Render(shared.FormatNumberLong(cur))))
	b.WriteString(fmt.Sprintf("tgt  %s\n", styleValue.Render(shared.FormatNumberLong(tgt))))

	if gap <= 0 {
		b.WriteString(fmt.Sprintf("gap  %s\n", styleMetYes.Render(shared.FormatNumberLong(gap))))
		b.WriteString("\n")
		b.WriteString(styleMetYes.Render("Already met") + "\n")
	} else {
		b.WriteString(fmt.Sprintf("gap  %s\n", styleMetNo.Render(shared.FormatNumberLong(gap))))
		b.WriteString("\n")
		opts := stationOptionsFor(res, gap, multis[res])
		for _, opt := range opts {
			b.WriteString(fmt.Sprintf("L%d %3dx  %s\n",
				opt.Level, opt.Count, renderRuneCost(opt.RuneCost)))
		}
	}

	return b.String()
}

// stationOptionsFor calls the right calculator helper per resource type.
func stationOptionsFor(res models.ResourceType, gap int, multi float64) []calculator.StationOptionResult {
	switch res {
	case models.ResourceMana:
		return calculator.ManaStationOptions(gap, multi)
	case models.ResourceSlime:
		return calculator.SlimeStationOptions(gap, multi)
	case models.ResourceDarkness:
		return calculator.DarknessStationOptions(gap, multi)
	}
	return nil
}

// renderRuneCost renders rune costs as coloured numbers only (no text labels).
// Each amount is right-aligned in a 6-char field (supports up to 5-digit costs)
// so columns stay aligned across all station levels.
func renderRuneCost(costs calculator.RuneCosts) string {
	runeStyles := map[models.RuneType]lipgloss.Style{
		models.RuneIce:    styleRuneIce,
		models.RunePoison: styleRunePoison,
		models.RuneBlood:  styleRuneBlood,
		models.RuneMoon:   styleRuneMoon,
	}
	order := []models.RuneType{models.RuneIce, models.RunePoison, models.RuneBlood, models.RuneMoon}
	var parts []string
	for _, rt := range order {
		amt, ok := costs[rt]
		if !ok || amt == 0 {
			continue
		}
		// Right-align in 6 chars BEFORE applying color so ANSI codes don't skew width.
		padded := fmt.Sprintf("%6s", shared.FormatNumberLong(amt))
		parts = append(parts, runeStyles[rt].Render(padded))
	}
	return strings.Join(parts, " ")
}
