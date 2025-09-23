package basic

import (
	"github.com/Ahu-Tools/AhuM/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RouteSignal uint

const (
	Quit RouteSignal = iota
	Err
	Back
	Next
	BackAndNext
)

type RouterModel interface {
	tea.Model
	Return(msg tea.Msg) (RouterModel, tea.Cmd) //Call on return
}

type Router struct {
	main     RouterModel
	quitting bool
	models   *util.Stack[RouterModel]
	err      error
}

type RouteMsg struct {
	model  RouterModel
	signal RouteSignal
	msg    tea.Msg
	err    error
}

func SignalRouter(model RouterModel, signal RouteSignal, msg tea.Msg) func() tea.Msg {
	return func() tea.Msg {
		return RouteMsg{
			model,
			signal,
			msg,
			nil,
		}
	}
}

func SignalQuit() func() tea.Msg {
	return func() tea.Msg {
		return RouteMsg{
			nil,
			Quit,
			nil,
			nil,
		}
	}
}

func SignalError(err error) func() tea.Msg {
	return func() tea.Msg {
		return RouteMsg{
			nil,
			Err,
			nil,
			err,
		}
	}
}

func NewRouter(init RouterModel) Router {
	stack := &util.Stack[RouterModel]{}
	return Router{
		main:   init,
		models: stack,
		err:    nil,
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
			m.quitting = true
			return m, tea.Quit

		case Err:
			oldModel, ok := m.models.Pop()
			if !ok {
				m.quitting = true
				m.err = msg.err
				return m, tea.Quit
			}
			m.main = oldModel
			var cmd tea.Cmd
			m.main, cmd = m.main.Return(msg.err)
			return m, cmd

		case Back:
			oldModel, ok := m.models.Pop()
			if !ok {
				return m.main, tea.Quit
			}
			m.main = oldModel
			var cmd tea.Cmd
			m.main, cmd = m.main.Return(msg.msg)
			return m, cmd

		case Next:
			m.models.Push(m.main)
			m.main = msg.model
			cmd := m.main.Init()
			if msg.msg == nil {
				return m, cmd
			}
			return m, tea.Batch(cmd, func() tea.Msg { return msg.msg })

		case BackAndNext:
			m.main = msg.model
			cmd := m.main.Init()
			if msg.msg == nil {
				return m, cmd
			}
			return m, tea.Batch(cmd, func() tea.Msg { return msg.msg })
		}
	}

	updatedModel, cmd := m.main.Update(msg)
	m.main = updatedModel.(RouterModel)
	return m, cmd
}

func (m Router) View() string {
	str := m.main.View()
	if m.err != nil {
		str = lipgloss.NewStyle().Foreground(lipgloss.Color(util.ErrorColor)).Render("Error: " + m.err.Error())
		if m.quitting {
			str += "\nquitting"
		}
	}

	return str
}
