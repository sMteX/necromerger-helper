package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necro-prestige-planner/internal/tui/tasks/prestigePlan"
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

	resourceCapModel  resourceCap.Model
	prestigePlanModel *prestigePlan.Model
}

func (m *AppModel) Init() tea.Cmd {
	// TODO: remove, temp
	m.currentTask = func() *taskType {
		t := taskPrestigePlan
		return &t
	}()
	m.resourceCapModel = resourceCap.New()
	m.prestigePlanModel = prestigePlan.New()
	return nil
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.currentTask == nil {
		return m.handleUpdateMainMenu(msg)
	}
	switch *m.currentTask {
	case taskResourceCap:
		newModel, cmd := m.resourceCapModel.Update(msg)
		m.resourceCapModel = newModel.(resourceCap.Model)
		return m, cmd
	case taskPrestigePlan:
		newModel, cmd := m.prestigePlanModel.Update(msg)
		m.prestigePlanModel = newModel.(*prestigePlan.Model)
		return m, cmd
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
		return m.prestigePlanModel.View()
	default:
		panic("unknown task type")
	}
}

func New() *AppModel {
	return &AppModel{}
}
