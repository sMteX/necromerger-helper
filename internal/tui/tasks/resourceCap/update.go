package resourceCap

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necromerger-helper/internal/models"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layoutPanels()
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		if m.focusedPanel == 1 {
			return m.updateRightPanel(msg)
		}
		return m.updateLeftPanel(msg)
	}

	// Forward to textinput for any other message types when a numeric field is active
	if m.focusedPanel == 0 && fieldKindOf(m.focusedField) == kindNumeric {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		m.commitNumericInput()
		m.recalculate()
		m.refreshViewport()
		return m, cmd
	}

	return m, nil
}

func (m *Model) layoutPanels() {
	// Width()/Height() set the total box size in lipgloss v2 (frame included),
	// so we assign the full terminal dimensions and let lipgloss handle the border/padding.
	m.leftWidth = m.width / 3
	if m.leftWidth < 34 {
		m.leftWidth = 34
	}
	m.rightWidth = m.width - m.leftWidth
	if m.rightWidth < 40 {
		m.rightWidth = 40
	}

	// Viewport fills the content area of the right panel: total height minus
	// the top+bottom border (2 chars; no top/bottom padding in stylePanelBorder).
	vpHeight := m.height - 2
	if vpHeight < 1 {
		vpHeight = 1
	}
	m.vp = viewport.New(viewport.WithWidth(m.rightWidth), viewport.WithHeight(vpHeight))
	m.vp.SoftWrap = true
	m.refreshViewport()
}

func (m *Model) refreshViewport() {
	m.vp.SetContent(m.renderOutput())
}

func (m Model) updateLeftPanel(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	kind := fieldKindOf(m.focusedField)

	switch msg.String() {
	case "tab", "shift+tab":
		m.focusedPanel = 1
		return m, nil

	// Threshold: t cycles forward through 200k → 400k → 600k → 800k → 200k
	case "t":
		m.thresholdIdx = (m.thresholdIdx + 1) % len(thresholds)
		m.recalculate()
		m.refreshViewport()
		return m, nil

	// Resource tab selection
	case "m":
		m.switchResource(models.ResourceMana)
		return m, nil
	case "s":
		m.switchResource(models.ResourceSlime)
		return m, nil
	case "d":
		m.switchResource(models.ResourceDarkness)
		return m, nil

	// Navigate fields
	case "up", "k":
		m.moveField(-1)
		return m, nil
	case "down", "j":
		m.moveField(1)
		return m, nil

	// Serv-O resource: left/right arrows cycle through resources
	case "right":
		if m.focusedField == fieldServOResource {
			m.servOResourceIdx = (m.servOResourceIdx + 1) % len(servOResources)
			m.recalculate()
			m.refreshViewport()
			return m, nil
		}
		// For numeric fields, fall through to textinput forwarding below

	case "left":
		if m.focusedField == fieldServOResource {
			m.servOResourceIdx = (m.servOResourceIdx - 1 + len(servOResources)) % len(servOResources)
			m.recalculate()
			m.refreshViewport()
			return m, nil
		}
		// For numeric fields, fall through to textinput forwarding below

	// Toggle / selector: enter or space activates
	case "enter", "space":
		if kind == kindToggle {
			m.toggleField()
			m.recalculate()
			m.refreshViewport()
		} else if kind == kindSelector {
			// enter/space also cycles selector forward (same as right arrow)
			m.servOResourceIdx = (m.servOResourceIdx + 1) % len(servOResources)
			m.recalculate()
			m.refreshViewport()
		}
		return m, nil
	}

	// Numeric field: forward key to textinput
	if kind == kindNumeric {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		m.commitNumericInput()
		m.recalculate()
		m.refreshViewport()
		return m, cmd
	}

	return m, nil
}

func (m Model) updateRightPanel(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab", "shift+tab":
		m.focusedPanel = 0
		m.focusInputToField()
		return m, nil
	}
	var cmd tea.Cmd
	m.vp, cmd = m.vp.Update(msg)
	return m, cmd
}

func (m *Model) moveField(delta int) {
	if fieldKindOf(m.focusedField) == kindNumeric {
		m.commitNumericInput()
	}
	m.focusedField = (m.focusedField + delta + fieldCount) % fieldCount
	m.focusInputToField()
}

func (m *Model) focusInputToField() {
	if fieldKindOf(m.focusedField) == kindNumeric {
		m.input.SetValue(m.currentNumericValue())
		m.input.Focus()
	} else {
		m.input.Blur()
	}
}

func (m *Model) switchResource(r models.ResourceType) {
	if fieldKindOf(m.focusedField) == kindNumeric {
		m.commitNumericInput()
	}
	m.activeResource = r
	m.focusInputToField()
	m.recalculate()
	m.refreshViewport()
}

func (m *Model) toggleField() {
	switch m.focusedField {
	case fieldServOUpgraded:
		m.servOUpgraded = !m.servOUpgraded
	case fieldGoldenBoosts:
		m.goldenBoosts = !m.goldenBoosts
	case fieldSkinBase:
		switch m.activeResource {
		case models.ResourceMana:
			m.skinWizard = !m.skinWizard
		case models.ResourceSlime:
			m.skinOozing = !m.skinOozing
		case models.ResourceDarkness:
			m.skinSid = !m.skinSid
		}
	case fieldSkinMult:
		switch m.activeResource {
		case models.ResourceMana:
			m.skinSanta = !m.skinSanta
		case models.ResourceSlime:
			m.skinBirthday = !m.skinBirthday
		case models.ResourceDarkness:
			m.skinWitch = !m.skinWitch
		}
	case fieldSkinGood:
		m.skinGood = !m.skinGood
	case fieldSkinRoyalty:
		m.skinRoyalty = !m.skinRoyalty
	}
}
