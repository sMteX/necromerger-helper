package shared

import (
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
)

type InputField struct {
	Label          string   // not going to be used every time, but it's convenient to have it here
	Step           int      // Step == 0 means the field is text-only; Step > 0 enables ←/→ increment/decrement.
	Options        []string // if len(Options) > 0, Step is ignored and the field is a select-like
	Width          int
	CharacterLimit int
	Validate       func(string) error
	InitialValue   string
	Input          textinput.Model
}

func InputValidationIntInRange(min, max int) func(string) error {
	return func(s string) error {
		v, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("must be a whole number")
		}
		if v < min || v > max {
			return fmt.Errorf("must be between %d and %d", min, max)
		}
		return nil
	}
}

func (f *InputField) CreateInput() textinput.Model {
	input := textinput.New()
	input.Prompt = ""
	input.CharLimit = f.CharacterLimit
	input.Validate = f.Validate
	input.SetVirtualCursor(true)
	input.SetWidth(f.Width)
	input.SetValue(f.InitialValue)
	return input
}
