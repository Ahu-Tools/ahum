package edge

import (
	"errors"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type EdgesForms struct {
	form *huh.Form

	prInfo    project.ProjectInfo
	edges     []project.Edge
	edgesStep int
}

func NewEdgesForms(prInfo project.ProjectInfo) EdgesForms {
	opts := GetEdges()

	edgesForm := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[Form]().
			Options(opts...).
			Title("Select edges you want to add to your project:").
			Key("edges"),
	))

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
	switch ef.form.State {
	case huh.StateAborted:
		return ef, basic.SignalError(errors.New("form aborted"))
	case huh.StateNormal:
		newForm, cmd := ef.form.Update(msg)
		ef.form = newForm.(*huh.Form)
		return ef, cmd
	}

	//Form is completed
	updatedEf, cmd := ef.goToForm()
	return updatedEf, cmd
}

func (ef EdgesForms) View() string {
	return ef.form.View()
}

func (ef EdgesForms) goToForm() (tea.Model, tea.Cmd) {
	edgesModel := ef.form.Get("edges").([]Form)
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

	// Pass through messages to the current edge form
	edgeModel := edgesModel[ef.edgesStep]
	edgeModel.InitProjectInfo(ef.prInfo)

	return ef, basic.SignalRouter(
		edgeModel,
		basic.Next,
		nil,
	)
}

func (ef EdgesForms) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	if edge, ok := msg.(project.Edge); ok {
		ef.edges[ef.edgesStep] = edge
		ef.edgesStep++

		return ef, func() tea.Msg { return msg }
	}
	err := errors.New("edge form didn't completed")
	if msg, ok := msg.(error); ok {
		err = msg
	}
	return ef, basic.SignalError(err)
}
