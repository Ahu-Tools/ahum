package infra

import (
	"errors"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	"github.com/Ahu-Tools/AhuM/pkg/util"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type SelectedInfra struct{}
type AbortedSelectionInfra struct{}
type FormCompleted struct{}
type CreationCompleted struct{}

type CreateForm struct {
	quitting bool

	form  *huh.Form
	prj   *project.Project
	infra project.Infra
}

func NewForm(prjPath string) (*CreateForm, error) {
	prj, err := project.LoadProject(prjPath)
	if err != nil {
		return nil, err
	}

	infras := GetInfras(prj.Info)
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[basic.RouterModel]().
			Options(infras...).
			Title("Select the infrastructure you want to add to your project:").
			Key("infra"),
	))

	form.SubmitCmd = func() tea.Msg { return SelectedInfra{} }
	form.CancelCmd = func() tea.Msg { return AbortedSelectionInfra{} }

	return &CreateForm{
		form: form,

		prj: prj,
	}, nil
}

func (c CreateForm) Init() tea.Cmd {
	return c.form.Init()
}

func (c CreateForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case SelectedInfra:
		infrasModel := c.form.Get("infra").(basic.RouterModel)

		return c, basic.SignalRouter(
			infrasModel,
			basic.Next,
			nil,
		)
	case AbortedSelectionInfra:
		return c, basic.SignalError(errors.New("infra selection aborted"))
	case FormCompleted:
		return c, basic.SignalRouter(
			basic.NewLoader(
				spinner.Moon,
				func(statChan chan string) error {
					return c.prj.AddInfra(c.infra, statChan)
				},
			),
			basic.Next,
			nil,
		)
	case CreationCompleted:
		c.quitting = true
		return c, basic.SignalQuit()
	}

	model, cmd := c.form.Update(msg)
	c.form = model.(*huh.Form)
	return c, cmd
}

func (c CreateForm) View() string {
	if !c.quitting {
		return c.form.View()
	} else {
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(util.SuccessColor))
		return style.Render("Infra added successfully!") + "\n"
	}
}

func (c CreateForm) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case project.Infra:
		c.infra = msg
		return c, func() tea.Msg { return FormCompleted{} }
	case basic.LoaderResultMsg:
		if msg.Err != nil {
			return c, basic.SignalError(msg.Err)
		}
		return c, func() tea.Msg { return CreationCompleted{} }
	case error:
		return c, basic.SignalError(msg)
	}

	return c, func() tea.Msg { return msg }
}
