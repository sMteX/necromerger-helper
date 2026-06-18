package resourceCap

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

// ── Resource colours ─────────────────────────────────────────────────────────
// These are the single source of truth; change here to affect all occurrences.
var (
	colorMana     = lipgloss.Color("#60A5FA") // blue
	colorSlime    = lipgloss.Color("#6EE7B7") // green
	colorDarkness = lipgloss.Color("#A78BFA") // purple
)

// resourceColor returns the canonical colour for a resource type.
func resourceColor(r models.ResourceType) color.Color {
	switch r {
	case models.ResourceSlime:
		return colorSlime
	case models.ResourceDarkness:
		return colorDarkness
	default:
		return colorMana
	}
}

// resourceStyle returns a style with the resource's foreground colour.
func resourceStyle(r models.ResourceType) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(resourceColor(r))
}

// ── General palette ───────────────────────────────────────────────────────────
var (
	colorMuted  = lipgloss.Color("#6B7280")
	colorActive = lipgloss.Color("#F9FAFB")
	colorAccent = lipgloss.Color("#60A5FA")
)

// ── Styles ────────────────────────────────────────────────────────────────────
var (
	styleLabel = lipgloss.NewStyle().
			Foreground(colorMuted)

	styleLabelFocused = lipgloss.NewStyle().
				Foreground(colorAccent)

	styleValue = lipgloss.NewStyle().
			Foreground(colorActive)

	styleSectionHeader = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorAccent)

	stylePanelBorder = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#374151")).
				Padding(0, 1)

	stylePanelBorderFocused = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorAccent).
				Padding(0, 1)

	styleThresholdOption = lipgloss.NewStyle().
				Foreground(colorMuted).
				PaddingLeft(1).
				PaddingRight(1)

	styleThresholdSelected = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#1F2937")).
				Background(colorAccent).
				PaddingLeft(1).
				PaddingRight(1)

	styleResourceTab = lipgloss.NewStyle().
				Foreground(colorMuted).
				PaddingLeft(1).
				PaddingRight(1)

	styleResourceTabActive = lipgloss.NewStyle().
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1)

	styleMetYes = lipgloss.NewStyle().Bold(true).Foreground(shared.ColorGood)
	styleMetNo  = lipgloss.NewStyle().Bold(true).Foreground(shared.ColorBad)

	styleRuneIce    = lipgloss.NewStyle().Foreground(shared.ColorIce)
	styleRunePoison = lipgloss.NewStyle().Foreground(shared.ColorPoison)
	styleRuneBlood  = lipgloss.NewStyle().Foreground(shared.ColorBlood)
	styleRuneMoon   = lipgloss.NewStyle().Foreground(shared.ColorMoon)
)
