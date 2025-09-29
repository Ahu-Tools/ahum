package edge

import (
	"errors"

	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type NextFormMsg struct{}
type AbortedFormMsg struct{}

type EdgesForms struct {
	form *huh.Form

	prInfo    project.ProjectInfo
	edges     []project.Edge
	edgesStep int
}

func NewEdgesForms(prInfo project.ProjectInfo) EdgesForms {
	opts := GetEdges(prInfo)

	edgesForm := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[basic.RouterModel]().
			Options(opts...).
			Title("Select edges you want to add to your project:").
			Key("edges"),
	))

	edgesForm.CancelCmd = func() tea.Msg { return AbortedFormMsg{} }
	edgesForm.SubmitCmd = func() tea.Msg { return NextFormMsg{} }

	return EdgesForms{
		form: edgesForm,

		prInfo:    prInfo,
		edgesStep: 0,
	}
}

func (ef EdgesForms) Init() tea.Cmd {
	return ef.form.Init()
}

func (ef EdgesForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case AbortedFormMsg:
		return ef, basic.SignalError(errors.New("form aborted"))
	case NextFormMsg:
		ef, cmd := ef.goToForm()
		return ef, cmd
	}

	newForm, cmd := ef.form.Update(msg)
	ef.form = newForm.(*huh.Form)
	return ef, cmd
}

func (ef EdgesForms) View() string {
	return ef.form.View()
}

func (ef EdgesForms) goToForm() (EdgesForms, tea.Cmd) {
	edgesModel := ef.form.Get("edges").([]basic.RouterModel)
	if ef.edgesStep == 0 {
		ef.edges = make([]project.Edge, len(edgesModel))
	}

	if ef.edgesStep == len(edgesModel) {
		return ef, basic.SignalRouter(
			nil,
			basic.Back,
			EdgesMsg{
				Ok:    true,
				Edges: ef.edges,
			},
		)
	}

	return ef, basic.SignalRouter(
		edgesModel[ef.edgesStep],
		basic.Next,
		nil,
	)
}

func (ef EdgesForms) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case project.Edge:
		ef.edges[ef.edgesStep] = msg
		ef.edgesStep++

		return ef, func() tea.Msg { return NextFormMsg{} }
	case error:
		return ef, basic.SignalError(msg)
	}
	return ef, func() tea.Msg { return msg }
}
