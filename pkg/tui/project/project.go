package project

import (
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	"github.com/Ahu-Tools/ahum/pkg/tui/edge"
	"github.com/Ahu-Tools/ahum/pkg/tui/infra"
	"github.com/Ahu-Tools/ahum/pkg/util"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectLevel uint

const (
	ProjectInfoLevel = iota
	EdgeSetupLevel
	InfraSetupLevel
	GenerationLevel
	CompletedLevel
)

type ProjectForms struct {
	prInfo project.ProjectInfo
	edges  []project.Edge
	infras []project.Infra
}

func NewProjectForms() ProjectForms {
	return ProjectForms{}
}

func (pjfs ProjectForms) Init() tea.Cmd {
	return levelCmd(ProjectInfoLevel)
}

func (pjfs ProjectForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	level, ok := msg.(ProjectLevel)
	if !ok {
		return pjfs, nil
	}

	switch level {
	case ProjectInfoLevel:
		return pjfs, basic.SignalRouter(
			NewInfoForm(),
			basic.Next,
			nil,
		)

	case EdgeSetupLevel:
		return pjfs, basic.SignalRouter(
			edge.NewEdgesForms(pjfs.prInfo),
			basic.Next,
			nil,
		)
	case InfraSetupLevel:
		return pjfs, basic.SignalRouter(
			infra.NewInfrasForms(pjfs.prInfo),
			basic.Next,
			nil,
		)

	case GenerationLevel:
		prj := project.NewProject(pjfs.prInfo, pjfs.infras, pjfs.edges)
		return pjfs, basic.SignalRouter(
			basic.NewLoader(
				spinner.Moon,
				prj.Generate,
			),
			basic.Next,
			nil,
		)

	case CompletedLevel:
		return pjfs, basic.SignalQuit()
	}
	return pjfs, nil
}

func (pjfs ProjectForms) View() string {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(util.SuccessColor))

	return style.Render("Generation completed!") + "\n"
}

func (pjfs ProjectForms) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case ProjectInfoMsg:
		pjfs.prInfo = msg.ProjectInfo
		return pjfs, levelCmd(EdgeSetupLevel)

	case edge.EdgesMsg:
		pjfs.edges = msg.Edges
		return pjfs, levelCmd(InfraSetupLevel)

	case infra.InfrasMsg:
		pjfs.infras = msg.Infras
		return pjfs, levelCmd(GenerationLevel)
	case basic.LoaderResultMsg:
		if msg.Err != nil {
			return pjfs, basic.SignalError(msg.Err)
		}
		return pjfs, levelCmd(CompletedLevel)
	case error:
		return pjfs, basic.SignalError(msg)
	}

	return pjfs, func() tea.Msg {
		return msg
	}
}

func levelCmd(pl ProjectLevel) tea.Cmd {
	return func() tea.Msg {
		return pl
	}
}
