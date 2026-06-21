package shared

import (
	"slices"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type TabModel struct {
	Cursor int
	Fields []InputField
}

func NewTabModel(size int) TabModel {
	return TabModel{
		Fields: make([]InputField, size),
	}
}

func (t *TabModel) InitializeInputs() {
	for i, field := range t.Fields {
		t.Fields[i].Input = field.CreateInput()
	}
}

func (t *TabModel) CurrentField() *InputField {
	return &t.Fields[t.Cursor]
}

func (t *TabModel) CurrentInput() *textinput.Model {
	return &t.CurrentField().Input
}

func (t *TabModel) HandleNonKeyMsg(msg tea.Msg) (cmd tea.Cmd) {
	t.CurrentField().Input, cmd = t.CurrentInput().Update(msg)
	return
}

func (t *TabModel) HandleUpKey(min int) tea.Cmd {
	if t.Cursor > min {
		t.CurrentInput().Blur()
		t.Cursor--
		return t.CurrentInput().Focus()
	}
	return nil
}

func (t *TabModel) HandleUpKeyFn(f func(int) bool) tea.Cmd {
	if f(t.Cursor) {
		t.CurrentInput().Blur()
		t.Cursor--
		return t.CurrentInput().Focus()
	}
	return nil
}

func (t *TabModel) HandleDownKey(max int) tea.Cmd {
	if t.Cursor < max {
		t.CurrentInput().Blur()
		t.Cursor++
		return t.CurrentInput().Focus()
	}
	return nil
}

func (t *TabModel) HandleDownKeyFn(f func(int) bool) tea.Cmd {
	if f(t.Cursor) {
		t.CurrentInput().Blur()
		t.Cursor++
		return t.CurrentInput().Focus()
	}
	return nil
}

func (t *TabModel) HandleStepKeys(key string) (newVal string, changed bool) {
	field := t.CurrentField()
	if len(field.Options) > 0 {
		cur := t.CurrentInput().Value()
		idx := slices.Index(field.Options, cur)
		if idx < 0 {
			idx = 0
		}
		if key == "left" && idx > 0 {
			idx--
		} else if key == "right" && idx < len(field.Options)-1 {
			idx++
		} else {
			return "", false
		}
		newVal = field.Options[idx]
		t.CurrentInput().SetValue(newVal)
		return newVal, true
	} else if field.Step > 0 {
		cur, err := strconv.Atoi(t.CurrentInput().Value())
		if err != nil {
			return "", false
		}
		if key == "left" {
			cur -= field.Step
		} else {
			cur += field.Step
		}
		newVal = strconv.Itoa(cur)
		if field.Validate != nil {
			if err := field.Validate(newVal); err != nil {
				return "", false
			}
		}
		t.CurrentInput().SetValue(newVal)
		return newVal, true
	}
	return "", false
}

func (t *TabModel) HandleInputKeyMsg(msg tea.KeyPressMsg) (cmd tea.Cmd, changed bool) {
	t.CurrentField().Input, cmd = t.CurrentInput().Update(msg)
	if t.CurrentInput().Err == nil {
		return cmd, true
	}
	return cmd, false
}
