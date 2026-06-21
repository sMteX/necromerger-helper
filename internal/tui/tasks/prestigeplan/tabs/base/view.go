package base

import (
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

func (m *Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		m.renderInput(fieldDevourerLevel),
		m.renderInput(fieldFeatTiers),
		m.renderInput(fieldOtherMultiplier),
		m.renderInput(fieldGroupBonusCount),
		m.renderInput(fieldLeftoverShards),
	)
}

func (m *Model) renderInput(i fieldIndex) string {
	labelStyle := lipgloss.NewStyle().Width(30)
	valueStyle := lipgloss.NewStyle().Width(10)

	field := m.Fields[i]
	if m.Cursor == int(i) {
		if field.Step > 0 || len(field.Options) > 0 {
			return labelStyle.Foreground(shared.Colors.Good).Render(field.Label) + valueStyle.Foreground(shared.Colors.Good).AlignHorizontal(lipgloss.Left).Render(shared.PadRight("< "+field.Input.Value()+" >", valueStyle.GetWidth()))
		}
		return labelStyle.Foreground(shared.Colors.Good).Render(field.Label) + "  " + field.Input.View()
	}
	return labelStyle.Render(field.Label) + valueStyle.Render(shared.PadRight("  "+field.Input.Value(), valueStyle.GetWidth()))
}

func (m *Model) GetHelpItems() []string {
	return []string{
		shared.Styles.Help.Render("↑ / ↓  navigate"),
		shared.Styles.Help.Render("← / →  change value"),
	}
}
