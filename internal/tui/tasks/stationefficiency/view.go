package stationefficiency

import (
	"fmt"
	"image/color"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necromerger-helper/internal/models"
	"github.com/sMteX/necromerger-helper/internal/tui/shared"
)

func (m *Model) View() tea.View {
	fw, fh := shared.Styles.MainContainer.GetFrameSize()

	// elements with more or less fixed size - will be subtracted from the available space to figure out the main content
	header := shared.Styles.Header.Render("Station efficiency")

	mainContentHeight := m.windowHeight - lipgloss.Height(header) - fh
	// empirically chosen
	inputPanelWidth := 38

	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.renderInputPanel(inputPanelWidth, mainContentHeight), // 35 = min width, 16 = min height
			// TODO: min main content size
			m.renderOutputs(m.windowWidth-fw-inputPanelWidth, mainContentHeight),
		),
	)
	content = shared.Styles.MainContainer.
		Width(m.windowWidth).
		Height(m.windowHeight).
		Render(content)

	return tea.View{
		Content:   content,
		AltScreen: true,
	}
}

func (m *Model) renderInputPanel(summaryWidth, summaryHeight int) string {
	return shared.Styles.SubContainer.
		Width(summaryWidth).
		Height(summaryHeight).
		Render(lipgloss.JoinVertical(lipgloss.Left,
			m.renderInput(fieldChampion),
			m.renderInput(fieldManaCap),
			m.renderInput(fieldSlimeCap),
			m.renderInput(fieldDarknessCap),
			m.renderInput(fieldSpeedPercent),
		))
}

func (m *Model) renderInput(i fieldIndex) string {
	labelStyle := lipgloss.NewStyle().Width(20)
	valueStyle := lipgloss.NewStyle().Width(15)

	field := m.Fields[i]
	if m.Cursor == int(i) {
		if field.Step > 0 || len(field.Options) > 0 {
			return labelStyle.Foreground(shared.Colors.Good).Render(field.Label) + valueStyle.Foreground(shared.Colors.Good).AlignHorizontal(lipgloss.Left).Render(shared.PadRight("< "+field.Input.Value()+" >", valueStyle.GetWidth()))
		}
		return labelStyle.Foreground(shared.Colors.Good).Render(field.Label) + "  " + field.Input.View()
	}
	return labelStyle.Render(field.Label) + valueStyle.Render(shared.PadRight("  "+field.Input.Value(), valueStyle.GetWidth()))
}

var resourceToColorMap = map[models.ResourceType]color.Color{
	models.ResourceMana:     shared.Colors.Mana,
	models.ResourceSlime:    shared.Colors.Slime,
	models.ResourceDarkness: shared.Colors.Darkness,
}

func (m *Model) renderOutputs(maxWidth, maxHeight int) string {
	levelColumn := lipgloss.NewStyle().Width(5).AlignHorizontal(lipgloss.Right).MarginRight(1)
	stationColumn := lipgloss.NewStyle().MarginRight(2) //.Width(17)
	valueColumn := lipgloss.NewStyle().Width(5)

	lines := []string{}
	if m.result != nil {
		// first find out the longest station name being used
		maxLength := 0
		for _, station := range m.result.Stations {
			if len(station.StationName) > maxLength {
				maxLength = len(station.StationName)
			}
		}
		stationColumn = stationColumn.Width(maxLength)
		for i, station := range m.result.Stations {
			// don't really care that far back, maybe not even past top 5
			if i > 10 {
				break
			}
			level := ""
			if station.Hacked {
				level = levelColumn.Foreground(shared.Colors.Good).Render(fmt.Sprintf("L%d", station.Level))
			} else {
				level = levelColumn.Foreground(shared.Colors.Bad).Render(fmt.Sprintf("L%d", station.Level))
			}
			lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top,
				level,
				stationColumn.Foreground(resourceToColorMap[station.Resource]).Render(station.StationName),
				valueColumn.Render(fmt.Sprintf("%.2f", station.ExpectedSummons)),
			))
		}
	}
	return shared.Styles.SubContainer.Width(maxWidth).Height(maxHeight).Render(
		lipgloss.JoinVertical(lipgloss.Left, lines...),
	)
}
