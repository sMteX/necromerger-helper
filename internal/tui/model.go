package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sMteX/necromerger-helper/internal/tui/tasks/prestigeplan"
	"github.com/sMteX/necromerger-helper/internal/tui/tasks/resourceCap"
	"github.com/sMteX/necromerger-helper/internal/tui/tasks/stationefficiency"
)

type taskType int8

const (
	taskResourceCap taskType = iota
	taskPrestigePlan
	taskStationEfficiency
	taskTypeCount
)

type AppModel struct {
	windowWidth, windowHeight int

	cursor      int
	currentTask *taskType

	resourceCapModel  resourceCap.Model
	prestigePlanModel *prestigeplan.Model
	stationEfficiencyModel *stationefficiency.Model
}

func (m *AppModel) Init() tea.Cmd {
	m.resourceCapModel = resourceCap.New()
	m.prestigePlanModel = prestigeplan.New()
	m.stationEfficiencyModel = stationefficiency.New()
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
		m.prestigePlanModel = newModel.(*prestigeplan.Model)
		return m, cmd
	case taskStationEfficiency:
		newModel, cmd := m.stationEfficiencyModel.Update(msg)
		m.stationEfficiencyModel = newModel.(*stationefficiency.Model)
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
	case taskStationEfficiency:
		return m.stationEfficiencyModel.View()
	default:
		panic("unknown task type")
	}
}

func New() *AppModel {
	return &AppModel{}
}
