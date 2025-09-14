package basic

import (
	"github.com/Ahu-Tools/AhuM/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
)

type MsgParams map[string]interface{}
type RouteSignal uint

const (
	Quit RouteSignal = iota
	Back
	Next
	BackAndNext
)

type RouterModel interface {
	tea.Model
	Inject(params MsgParams) RouterModel            //Called on initialisation
	Return(params MsgParams) (RouterModel, tea.Cmd) //Call on return
}

type Router struct {
	main     RouterModel
	quitting bool
	models   *util.Stack[RouterModel]
}

type RouteMsg struct {
	model  RouterModel
	signal RouteSignal
	params MsgParams
}

func SignalRouter(model RouterModel, signal RouteSignal, params MsgParams) func() tea.Msg {
	return func() tea.Msg {
		return RouteMsg{
			model,
			signal,
			params,
		}
	}
}

func NewRouter(init RouterModel) Router {
	stack := &util.Stack[RouterModel]{}
	return Router{
		main:   init,
		models: stack,
	}
}

func (m Router) Init() tea.Cmd {
	return m.main.Init()
}

func (m Router) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	case RouteMsg:
		switch msg.signal {
		case Quit:
			return m, tea.Quit

		case Back:
			oldModel, ok := m.models.Pop()
			if !ok {
				return m.main, tea.Quit
			}
			m.main = oldModel
			var cmd tea.Cmd
			m.main, cmd = m.main.Return(msg.params)
			return m, cmd

		case Next:
			m.models.Push(m.main)
			m.main = msg.model
			m.main = m.main.Inject(msg.params)
			cmd := m.main.Init()
			updatedModel, cmd := m.main.Update(cmd())
			m.main = updatedModel.(RouterModel)
			return m, cmd

		case BackAndNext:
			m.main = msg.model
			m.main = m.main.Inject(msg.params)
			cmd := m.main.Init()
			updatedModel, cmd := m.main.Update(cmd())
			m.main = updatedModel.(RouterModel)
			return m, cmd
		}
	}

	updatedModel, cmd := m.main.Update(msg)
	m.main = updatedModel.(RouterModel)
	return m, cmd
}

func (m Router) View() string {
	str := m.main.View()
	if m.quitting {
		str += "\n\nquitting\n"
	}
	return str
}
