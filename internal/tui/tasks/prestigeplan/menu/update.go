package menu

import (
	tea "charm.land/bubbletea/v2"
)

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	keyMsg, isKey := msg.(tea.KeyPressMsg)
	if !isKey {
		// Non-key messages are cursor blink ticks. Only the name editor needs
		// them; other states have no active textinput.
		if m.state == stateNameEditor {
			var cmd tea.Cmd
			if m.nameFocus {
				m.nameInput, cmd = m.nameInput.Update(msg)
			} else {
				m.notesInput, cmd = m.notesInput.Update(msg)
			}
			return m, cmd
		}
		return m, nil
	}
	return m.handleKey(keyMsg)
}

func (m *Model) handleKey(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch m.state {
	case stateMenu:
		return m.handleMenuKey(msg)
	case stateList:
		return m.handleListKey(msg)
	case stateNameEditor:
		return m.handleNameEditorKey(msg)
	case stateConfirm:
		return m.handleConfirmKey(msg)
	}
	return m, nil
}

func (m *Model) handleMenuKey(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.menuCursor > 0 {
			m.menuCursor--
		}
	case "down":
		if m.menuCursor < menuItemCount-1 {
			m.menuCursor++
		}
	case "enter":
		return m.selectMenuItem()
	case "x", "esc":
		// x mirrors the lazydocker convention of toggling the menu with the
		// same key that opened it. Both x and esc dismiss without any action.
		m.result = &Action{Type: ActionNone}
	}
	return m, nil
}

func (m *Model) selectMenuItem() (*Model, tea.Cmd) {
	switch m.menuCursor {
	case menuNew:
		if m.cfg.Dirty {
			// Unsaved changes — ask the user to confirm before discarding.
			m.pendingConfirm.action = pendingNew
			m.state = stateConfirm
			m.confirmCursor = confirmNo // default to "No" to prevent accidental loss
		} else {
			m.result = &Action{Type: ActionNew}
		}
	case menuLoad:
		if len(m.cfg.Plans) > 0 {
			m.state = stateList
			m.listCursor = 0
		}
		// If there are no saved plans, stay on the menu (nothing to show).
	case menuSaveAs:
		m.saveTitle = "Save plan as"
		m.state = stateNameEditor
		return m, m.nameInput.Focus()
	case menuClose:
		m.result = &Action{Type: ActionNone}
	}
	return m, nil
}

func (m *Model) handleListKey(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.listCursor > 0 {
			m.listCursor--
		}
	case "down":
		if m.listCursor < len(m.cfg.Plans)-1 {
			m.listCursor++
		}
	case "enter":
		selected := m.cfg.Plans[m.listCursor]
		if m.cfg.Dirty {
			// Unsaved changes — confirm before discarding. Store the chosen
			// path so the confirm handler knows where to load from.
			m.pendingConfirm.action = pendingLoad
			m.pendingConfirm.path = selected.Path
			m.state = stateConfirm
			m.confirmCursor = confirmNo
		} else {
			m.result = &Action{Type: ActionLoad, Path: selected.Path}
		}
	case "esc":
		m.state = stateMenu
	}
	return m, nil
}

func (m *Model) handleNameEditorKey(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "tab", "shift+tab":
		// Toggle focus between the name and notes inputs.
		m.nameFocus = !m.nameFocus
		if m.nameFocus {
			m.notesInput.Blur()
			return m, m.nameInput.Focus()
		}
		m.nameInput.Blur()
		return m, m.notesInput.Focus()
	case "enter":
		name := m.nameInput.Value()
		if name == "" {
			name = "Unnamed plan"
		}
		m.result = &Action{
			Type:  ActionSave,
			Path:  m.cfg.PlansDir + "/" + planFileName(name),
			Name:  name,
			Notes: m.notesInput.Value(),
		}
		return m, nil
	case "esc":
		// Cancel the save and go back to the top menu, leaving the inputs as-is.
		m.nameInput.Blur()
		m.notesInput.Blur()
		m.nameFocus = true
		m.state = stateMenu
		return m, nil
	}

	// Any other key goes to whichever textinput is currently focused.
	var cmd tea.Cmd
	if m.nameFocus {
		m.nameInput, cmd = m.nameInput.Update(msg)
	} else {
		m.notesInput, cmd = m.notesInput.Update(msg)
	}
	return m, cmd
}

func (m *Model) handleConfirmKey(msg tea.KeyPressMsg) (*Model, tea.Cmd) {
	switch msg.String() {
	case "up", "down":
		// Toggle between Yes and No — only two options so direction doesn't matter.
		if m.confirmCursor == confirmYes {
			m.confirmCursor = confirmNo
		} else {
			m.confirmCursor = confirmYes
		}
	case "enter":
		if m.confirmCursor == confirmYes {
			// User confirmed the discard — fire the originally-intended action.
			switch m.pendingConfirm.action {
			case pendingNew:
				m.result = &Action{Type: ActionNew}
			case pendingLoad:
				m.result = &Action{Type: ActionLoad, Path: m.pendingConfirm.path}
			}
		} else {
			// User chose No — go back to the menu without doing anything.
			m.state = stateMenu
		}
	case "esc":
		m.state = stateMenu
	}
	return m, nil
}

// planFileName converts a name to a safe filename (no directory, .json extension).
// Spaces and underscores become underscores; everything else non-alphanumeric is dropped.
func planFileName(name string) string {
	out := make([]byte, 0, len(name))
	for _, c := range []byte(name) {
		switch {
		case c >= 'a' && c <= 'z':
			out = append(out, c)
		case c >= 'A' && c <= 'Z':
			out = append(out, c+32) // to lower
		case c >= '0' && c <= '9':
			out = append(out, c)
		case c == ' ' || c == '_':
			out = append(out, '_')
		case c == '-':
			out = append(out, '-')
		}
	}
	if len(out) == 0 {
		return "plan"
	}
	return string(out) + ".json"
}
