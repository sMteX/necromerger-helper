package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/resourceCap"
)

type taskType int8

const (
	taskResourceCap taskType = iota
	taskPrestigePlan
)

type AppModel struct {
	windowWidth, windowHeight int

	cursor      int
	currentTask *taskType

	resourceCapModel resourceCap.Model
}

func (m *AppModel) Init() tea.Cmd {
	m.resourceCapModel = resourceCap.New()
	return nil
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.currentTask == nil {
		return m.handleUpdateMainMenu(msg)
	}
	switch *m.currentTask {
	case taskResourceCap:
		return m.resourceCapModel.Update(msg)
	case taskPrestigePlan:
		// TODO: implement me
		return m, nil
	default:
		panic("unknown task type")
	}
}

func (m *AppModel) View() tea.View {
	if m.currentTask == nil {
		return m.renderMainMenu()
	}
	switch *m.currentTask {
	case taskResourceCap:
		return m.resourceCapModel.View()
	case taskPrestigePlan:
		// TODO: implement me
		return m.renderMainMenu()
	default:
		panic("unknown task type")
	}
}

func New() *AppModel {
	return &AppModel{}
}
