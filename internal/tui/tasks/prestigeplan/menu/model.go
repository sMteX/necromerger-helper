package menu

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
	"github.com/sMteX/necro-prestige-planner/internal/persistence"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
)

// menuState drives which screen is currently shown inside the modal.
type menuState int

const (
	stateMenu       menuState = iota // top-level menu (New / Load / Save as / Close)
	stateList                        // scrollable list of saved plans (Load flow)
	stateNameEditor                  // name + notes inputs (Save As flow)
	stateConfirm                     // "discard unsaved changes?" dialog
)

type menuItem int

const (
	menuNew menuItem = iota
	menuLoad
	menuSaveAs
	menuClose
	menuItemCount // used to clamp cursor movement
)

type confirmOption int

const (
	confirmYes confirmOption = iota
	confirmNo
)

// pendingAction remembers what the user wanted to do when the confirm dialog
// interrupted them (e.g. they chose Load while there were unsaved changes).
type pendingAction int

const (
	pendingNew pendingAction = iota
	pendingLoad
)

// ActionType is the result the modal signals back to the parent after closing.
type ActionType int

const (
	ActionNone ActionType = iota // dismissed, nothing to do
	ActionNew                    // start a fresh plan
	ActionLoad                   // load the plan at Path
	ActionSave                   // save current plan with Name/Notes to Path
)

// Action is set on Model.result when the user finishes interacting with the
// modal. The parent polls Result() each update cycle.
type Action struct {
	Type  ActionType
	Path  string // ActionLoad: path to load; ActionSave: path to save to
	Name  string // ActionSave
	Notes string // ActionSave
}

type Config struct {
	PlansDir    string
	Plans       []persistence.PlanMeta
	Dirty       bool   // whether the current plan has unsaved changes
	StartAtSave bool   // skip the menu and go straight to the name editor
	Name        string // pre-fill the name input with the current plan name
	Notes       string // pre-fill the notes input with the current plan notes
}

type Model struct {
	state      menuState
	cfg        Config
	menuCursor menuItem
	listCursor int

	nameInput  textinput.Model
	notesInput textinput.Model
	nameFocus  bool   // true = name input is active, false = notes input
	saveTitle  string // "Save plan" or "Save plan as" depending on the flow

	confirmCursor  confirmOption
	pendingConfirm struct {
		action pendingAction
		path   string // relevant only for pendingLoad
	}

	// result is nil while the modal is open. Once set, the parent should read
	// it and close the modal.
	result *Action
}

func New(cfg Config) *Model {
	// Apply a dim colour to placeholder text so it's visually distinct from real input.
	placeholderStyle := lipgloss.NewStyle().Foreground(shared.Colors.Dim)

	nameInput := textinput.New()
	nameInput.Prompt = ""
	nameInput.CharLimit = 64
	nameInput.Placeholder = "Plan name"
	nameInput.SetWidth(26)
	nameInput.SetValue(cfg.Name)
	styles := nameInput.Styles()
	styles.Focused.Placeholder = placeholderStyle
	styles.Blurred.Placeholder = placeholderStyle
	nameInput.SetStyles(styles)

	notesInput := textinput.New()
	notesInput.Prompt = ""
	notesInput.CharLimit = 128
	notesInput.Placeholder = "Notes (optional)"
	notesInput.SetWidth(26)
	notesInput.SetValue(cfg.Notes)
	styles = notesInput.Styles()
	styles.Focused.Placeholder = placeholderStyle
	styles.Blurred.Placeholder = placeholderStyle
	notesInput.SetStyles(styles)

	m := &Model{
		cfg:        cfg,
		nameInput:  nameInput,
		notesInput: notesInput,
		nameFocus:  true,
		saveTitle:  "Save plan",
	}

	if cfg.StartAtSave {
		// Opened via ctrl+s with no existing save path — skip the top menu and
		// go straight to the name editor. Focus is handled here rather than in
		// Update because New() is called once and Focus() returns a tea.Cmd we
		// don't need to thread back (the cursor blink starts on the next tick).
		m.state = stateNameEditor
		m.nameInput.Focus()
	} else {
		m.state = stateMenu
	}

	return m
}

func (m *Model) Result() *Action {
	return m.result
}
