package edge

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	"github.com/Ahu-Tools/AhuM/pkg/util"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type FormCompleted struct{}
type CreationCompleted struct{}

type CreateForm struct {
	quitting bool

	form    *huh.Form
	prjPath string
	prj     *project.Project
	edge    project.Edge
}

func NewForm(prjPath string) *CreateForm {
	edges := GetEdges()
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[Form]().
			Options(edges...).
			Title("Select the edge you want to add to your project:").
			Key("edge"),
	))

	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &CreateForm{
		form: form,

		prjPath: prjPath,
	}
}

func (c CreateForm) Init() tea.Cmd {
	return c.form.Init()
}

func (c CreateForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch c.form.State {
	case huh.StateNormal:
		model, cmd := c.form.Update(msg)
		c.form = model.(*huh.Form)
		return c, cmd
	case huh.StateAborted:
		return c, basic.SignalRouter(
			nil,
			basic.Back,
			nil,
		)
	}

	switch msg.(type) {
	case FormCompleted:
		return c, basic.SignalRouter(
			basic.NewLoader(
				spinner.Moon,
				func(statChan chan string) error {
					return c.prj.AddEdge(c.edge, statChan)
				},
			),
			basic.Next,
			nil,
		)
	case CreationCompleted:
		c.quitting = true
		return c, basic.SignalQuit()
	}

	//Select edge form is completed
	pr, err := project.LoadProject(c.prjPath)
	if err != nil {
		return c, basic.SignalError(err)
	}
	c.prj = pr

	edgesModel := c.form.Get("edge").(Form)
	edgesModel.InitProjectInfo(pr.Info)

	// Pass messages to the selected edge form
	return c, basic.SignalRouter(
		edgesModel,
		basic.Next,
		nil,
	)
}

func (c CreateForm) View() string {
	if c.form.State == huh.StateCompleted {
		c.form.View()
	}

	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(util.SuccessColor))
	return style.Render("Edge added successfully!") + "\n"
}

func (c CreateForm) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case project.Edge:
		c.edge = msg
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
