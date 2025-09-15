package project

import (
	"fmt"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectLevels uint

// statusUpdateMsg is sent for each status update from the generation process
type statusUpdateMsg string

type genStatus struct {
	err error
}

const (
	ProjectInfoLevel = iota
	InfraSetupLevel
	GenerationLevel
)

type ProjectForms struct {
	spinner  spinner.Model
	quitting bool
	err      error

	statusChan chan string
	statusDesc string

	level  ProjectLevels
	prInfo project.ProjectInfo
	infras []project.Infra
}

func NewProjectForms() ProjectForms {
	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return ProjectForms{
		spinner:    s,
		quitting:   false,
		err:        nil,
		statusChan: make(chan string),
		level:      ProjectInfoLevel,
	}
}

func (pjfs ProjectForms) Init() tea.Cmd {
	return basic.SignalRouter(
		NewInfoForm(),
		basic.Next,
		nil,
	)
}

func (pjfs ProjectForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusUpdateMsg:
		pjfs.statusDesc = string(msg)
		return pjfs, tea.Batch(
			pjfs.spinner.Tick,
			waitForStatusUpdates(pjfs.statusChan),
		)
	case genStatus:
		pjfs.quitting = true
		pjfs.err = msg.err
		return pjfs, basic.SignalRouter(nil, basic.Quit, nil)
	default:
		var cmd tea.Cmd
		pjfs.spinner, cmd = pjfs.spinner.Update(msg)
		return pjfs, cmd
	}
}

func (pjfs ProjectForms) View() string {
	style := lipgloss.NewStyle().Bold(true)

	if pjfs.quitting {
		if pjfs.err != nil {
			return style.Foreground(lipgloss.Color("#FF0000")).Render("Generation failed. Reason: "+pjfs.err.Error()) + "\n"
		}
		return style.Foreground(lipgloss.Color("#76FF03")).Render("Generation completed!") + "\n"
	}

	return style.Render(fmt.Sprintf("%s %s", pjfs.spinner.View(), pjfs.statusDesc))
}

func (pjfs ProjectForms) Inject(params basic.MsgParams) basic.RouterModel {
	return pjfs
}

func (pjfs ProjectForms) Return(params basic.MsgParams) (basic.RouterModel, tea.Cmd) {
	switch pjfs.level {
	case ProjectInfoLevel:
		if !params["ok"].(bool) {
			pjfs.quitting = true
			return pjfs, basic.SignalRouter(nil, basic.Quit, nil)
		}

		pjfs.prInfo = params["project_info"].(project.ProjectInfo)
		pjfs.level++
		return pjfs, basic.SignalRouter(
			NewInfrasForms(pjfs.prInfo),
			basic.Next,
			nil,
		)
	case InfraSetupLevel:
		pjfs.infras = params["infras"].([]project.Infra)
		pjfs.level++

		proj := project.NewProject(pjfs.prInfo, pjfs.infras)

		return pjfs, tea.Batch(
			pjfs.spinner.Tick,
			genProjCmd(proj, pjfs.statusChan),
			waitForStatusUpdates(pjfs.statusChan),
		)
	}
	return pjfs, pjfs.spinner.Tick
}

func genProjCmd(proj project.Project, statusChan chan string) tea.Cmd {
	return func() tea.Msg {
		return genStatus{proj.Generate(statusChan)}
	}
}

func waitForStatusUpdates(ch chan string) tea.Cmd {
	return func() tea.Msg {
		// This will block until a message is sent on the channel.
		// When the channel is closed, `ok` will be `false`.
		status, ok := <-ch
		if !ok {
			// Channel has been closed, which means the process is done.
			// We don't need to send a message, the other command will.
			return nil
		}
		return statusUpdateMsg(status)
	}
}
