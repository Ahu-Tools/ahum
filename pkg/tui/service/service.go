package service

import (
	"fmt"

	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/service"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type statusUpdate struct {
	statusDesc string
	err        error
	done       bool
}

type ServiceForm struct {
	loading    spinner.Model
	statusDesc string
	quitting   bool

	err        error
	statusChan chan string

	project *project.Project
	form    *huh.Form
	svcData *service.ServiceData
}

func NewServiceForm(projectPath string) ServiceForm {
	project, err := project.LoadProject(projectPath)
	svcData := service.NewServiceData("", "")

	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter your service name:").
			Validate(huh.ValidateNotEmpty()).
			Value(&svcData.Name).
			Key("name"),
		huh.NewInput().
			Title("Enter your service package name:").
			Validate(huh.ValidateNotEmpty()).
			Value(&svcData.PackageName).
			Key("package_name"),
	))

	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return ServiceForm{
		loading:    s,
		statusDesc: "Loading...",
		err:        err,
		quitting:   false,

		statusChan: make(chan string),

		project: project,
		form:    form,
		svcData: svcData,
	}
}

func (sf ServiceForm) Init() tea.Cmd {
	if sf.err != nil {
		return func() tea.Msg { return statusUpdate{err: sf.err, done: true} }
	}
	return sf.form.Init()
}

func (sf ServiceForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusUpdate:
		if msg.done {
			sf.quitting = true
			sf.err = msg.err
			sf.statusDesc = msg.statusDesc
			return sf, tea.Quit
		}

		sf.statusDesc = msg.statusDesc
		return sf, tea.Batch(
			sf.loading.Tick,
			waitForStatusUpdates(sf.statusChan),
		)
	case spinner.TickMsg:
		return sf, sf.loading.Tick
	}

	switch sf.form.State {
	case huh.StateCompleted:
		return sf, tea.Batch(
			sf.loading.Tick,
			waitForStatusUpdates(sf.statusChan),
			sf.Submit(),
		)
	case huh.StateAborted:
		sf.err = fmt.Errorf("User aborted")
		sf.quitting = true
		return sf, tea.Quit
	}

	form, cmd := sf.form.Update(msg)
	sf.form = form.(*huh.Form)

	return sf, cmd
}

func (sf ServiceForm) View() string {
	if sf.form.State == huh.StateNormal {
		return sf.form.View()
	}

	style := lipgloss.NewStyle().Bold(true)

	if sf.quitting {
		if sf.err != nil {
			return style.Foreground(lipgloss.Color("#FF0000")).Render("Generation failed. Reason: "+sf.err.Error()) + "\n"
		}
		return style.Foreground(lipgloss.Color("#76FF03")).Render("Generation completed!") + "\n"
	}

	return style.Render(fmt.Sprintf("%s %s", sf.loading.View(), sf.statusDesc))
}

func (sf *ServiceForm) Submit() tea.Cmd {
	return func() tea.Msg {
		svc := service.NewService(sf.project, *sf.svcData)
		err := svc.Generate(sf.statusChan)
		return statusUpdate{err: err, done: true}
	}
}

func waitForStatusUpdates(statusChan chan string) tea.Cmd {
	return func() tea.Msg {
		status := <-statusChan
		return statusUpdate{statusDesc: status, done: false}
	}
}
