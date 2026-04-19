# Charm TUI API Reference

Researched from module cache source. Versions: bubbletea v2.0.6, lipgloss v2.0.3, bubbles v2.1.0.
Import paths: `charm.land/bubbletea/v2`, `charm.land/lipgloss/v2`, `charm.land/bubbles/v2`.

---

## Bubbletea v2

### Model interface
```go
type Model interface {
    Init() Cmd
    Update(Msg) (Model, Cmd)
    View() View   // returns View struct, NOT a string — this is the key v2 change
}
```

### View struct (v2 change)
```go
tea.View{
    Content:         string,        // ANSI/lipgloss-rendered string goes here
    WindowTitle:     string,
    AltScreen:       bool,
    MouseMode:       MouseMode,     // MouseModeNone | CellMotion | AllMotion
    ReportFocus:     bool,
    BackgroundColor: color.Color,
    ForegroundColor: color.Color,
}
```

### Key events (v2 change: split into press/release)
```go
// In Update, type-switch on KeyPressMsg for normal use (KeyReleaseMsg exists but rarely needed)
case tea.KeyPressMsg:
    switch msg.String() {
    case "ctrl+c", "q": return m, tea.Quit()
    case "tab":         // switch focus
    case "1", "2":      // direct key binds
    case "up", "down":  // arrow keys
    }
```

### WindowSizeMsg
```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
```

### Commands
```go
tea.Quit()
tea.Batch(cmd1, cmd2)
tea.Sequence(cmd1, cmd2)   // sequential
tea.Tick(duration, func(time.Time) Msg)
```

### Program
```go
p := tea.NewProgram(model,
    tea.WithContext(ctx),
    // No WithAltScreen option — set AltScreen in View struct instead
)
p.Run()
```

---

## Lipgloss v2

### Style basics
```go
s := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#ff0000")).  // hex
    Background(lipgloss.ANSIColor(235)).    // ANSI256
    Width(40).
    Height(20).
    Padding(1, 2).          // top/bottom, left/right
    Margin(0, 1).
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#444"))

output := s.Render("text")
```

### Colors
```go
lipgloss.Color("#rrggbb")       // hex
lipgloss.Color("205")           // ANSI256 index as string
lipgloss.ANSIColor(n)           // ANSI256 int
lipgloss.RGBColor{R, G, B}
lipgloss.NoColor{}
// Named: lipgloss.Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Bright* variants
```

### Border styles
```go
lipgloss.NormalBorder()
lipgloss.RoundedBorder()
lipgloss.ThickBorder()
lipgloss.DoubleBorder()
lipgloss.HiddenBorder()    // invisible but keeps layout space
lipgloss.ASCIIBorder()
```

### Layout
```go
// Side by side — position is lipgloss.Top / Center / Bottom, or 0.0–1.0
lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

// Stacked — position is lipgloss.Left / Center / Right, or 0.0–1.0
lipgloss.JoinVertical(lipgloss.Left, top, bottom)

// Place in a sized box
lipgloss.Place(width, height, hAlign, vAlign, content)
```

---

## Bubbles v2

### textinput
```go
input := textinput.New()
input.SetValue("42")
input.SetWidth(10)
input.SetPrompt("")
input.CharLimit = 6
input.Validate = func(s string) error { ... }   // live validation

cmd := input.Focus()
input.Blur()
input.Focused() bool
input.Value() string

// In Update, forward messages when focused:
input, cmd = input.Update(msg)
```

### viewport
```go
vp := viewport.New(viewport.WithWidth(w), viewport.WithHeight(h))
vp.SetContent("multiline\nstring")
vp.SetContentLines([]string{...})
vp.SoftWrap = true
vp.MouseWheelEnabled = true

// In Update:
vp, cmd = vp.Update(msg)
// In View:
vp.View()
```

### table
```go
cols := []table.Column{
    {Title: "Level", Width: 7},
    {Title: "Count", Width: 7},
    {Title: "Rune cost", Width: 30},
}
rows := []table.Row{{"L1", "17", "170 Ice  85 Poison"}}

t := table.New(
    table.WithColumns(cols),
    table.WithRows(rows),
    table.WithHeight(10),
    table.WithFocused(true),
    table.WithStyles(table.DefaultStyles()),
)
t.SelectedRow() table.Row
t.SetRows(rows)

t.SetStyles(table.Styles{
    Header:   lipgloss.NewStyle()...,
    Cell:     lipgloss.NewStyle()...,
    Selected: lipgloss.NewStyle()...,
})
```

### list
```go
// Items must implement:
type Item interface { FilterValue() string }

l := list.New(items, delegate, width, height)
l.Title = "title"
l.SelectedItem() list.Item
l.SetItems(newItems)
```

---

## No built-in select/picker

Bubbles has no fixed-option selector component. For small sets (threshold: 200k/400k/600k/800k,
resource type: Mana/Slime/Darkness), use a hand-rolled inline selector: render options as a
styled string with the active one highlighted, driven by direct key binds (e.g. `1-4` for
thresholds, `m/s/d` or `1-3` for resource type).

---

## Two-panel layout pattern

```go
func (m Model) View() tea.View {
    left  := lipgloss.NewStyle().Width(m.leftWidth).Height(m.height).
                 Border(lipgloss.RoundedBorder()).Render(m.renderInputs())
    right := lipgloss.NewStyle().Width(m.rightWidth).Height(m.height).
                 Border(lipgloss.RoundedBorder()).Render(m.viewport.View())
    return tea.View{
        Content: lipgloss.JoinHorizontal(lipgloss.Top, left, right),
    }
}
```