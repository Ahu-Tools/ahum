package basic

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoaderResultMsg struct {
	Err error
}

type LoaderTask func(chan string) error

type StatusUpdate struct {
	done bool
	desc string
	err  error
}

type Loader struct {
	spinner  spinner.Model
	statChan chan string

	desc string
	task LoaderTask
}

func NewLoader(spinType spinner.Spinner, task LoaderTask) Loader {
	spin := spinner.New()
	spin.Spinner = spinType
	spin.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Loader{
		spinner:  spin,
		statChan: make(chan string),

		desc: "Loading",
		task: task,
	}
}

func (l Loader) Init() tea.Cmd {
	return tea.Batch(
		l.spinner.Tick,
		l.loadTask,
	)
}

func (l Loader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatusUpdate:
		if msg.done {
			return l, SignalRouter(
				nil,
				Back,
				LoaderResultMsg{
					Err: msg.err,
				},
			)
		} else {
			l.desc = msg.desc
			return l, l.updateStatus
		}
	default:
		spin, cmd := l.spinner.Update(msg)
		l.spinner = spin
		return l, cmd
	}
}

func (l Loader) View() string {
	style := lipgloss.NewStyle().Bold(true)

	return style.Render(fmt.Sprintf("%s %s", l.spinner.View(), l.desc))
}

func (l Loader) Return(msg tea.Msg) (RouterModel, tea.Cmd) {
	return l, l.Init()
}

func (l Loader) updateStatus() tea.Msg {
	desc, ok := <-l.statChan

	if !ok {
		return nil
	}

	return StatusUpdate{
		done: false,
		desc: desc,
		err:  nil,
	}
}

func (l Loader) loadTask() tea.Msg {
	err := l.task(l.statChan)
	return StatusUpdate{
		done: true,
		err:  err,
	}
}
