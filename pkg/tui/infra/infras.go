package infra

import (
	"errors"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type CompletedInfra struct{}

type InfrasForms struct {
	form *huh.Form

	prInfo     project.ProjectInfo
	infras     []project.Infra
	infrasStep int
}

func NewInfrasForms(prInfo project.ProjectInfo) InfrasForms {
	opts := GetInfras()

	infrasForm := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[Form]().
			Options(opts...).
			Title("Select infrastructures you want to add to your project:").
			Key("infras"),
	))

	return InfrasForms{
		form: infrasForm,

		prInfo:     prInfo,
		infrasStep: 0,
	}
}

func (inf InfrasForms) Init() tea.Cmd {
	return inf.form.Init()
}

func (inf InfrasForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch inf.form.State {
	case huh.StateAborted:
		return inf, basic.SignalError(errors.New("form aborted"))
	case huh.StateNormal:
		newForm, cmd := inf.form.Update(msg)
		inf.form = newForm.(*huh.Form)
		return inf, cmd
	}

	//huh form completed

	if inf.infrasStep == len(inf.infras) {
		return inf, basic.SignalRouter(
			nil,
			basic.Back,
			InfrasMsg{
				Infras: inf.infras,
			},
		)
	}
	newInf, cmd := inf.goToForm()
	return newInf, cmd
}

func (inf InfrasForms) goToForm() (InfrasForms, tea.Cmd) {
	infrasModel := inf.form.Get("infras").([]Form)
	if inf.infrasStep == 0 {
		inf.infras = make([]project.Infra, len(infrasModel))
	}

	if len(infrasModel) == inf.infrasStep {
		return inf, basic.SignalRouter(
			nil,
			basic.Back,
			InfrasMsg{
				Infras: inf.infras,
			},
		)
	}

	infraModel := infrasModel[inf.infrasStep]
	infraModel.InitProjectInfo(inf.prInfo)

	return inf, basic.SignalRouter(
		infraModel,
		basic.Next,
		nil,
	)
}

func (inf InfrasForms) View() string {
	return inf.form.View()
}

func (inf InfrasForms) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case project.Infra:
		inf.infras[inf.infrasStep] = msg
		inf.infrasStep++

		return inf, func() tea.Msg {
			return CompletedInfra{}
		}
	case error:
		return inf, basic.SignalError(msg)
	}

	return inf, func() tea.Msg { return msg }
}
