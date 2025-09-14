package project

import (
	"fmt"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type ProjectLevels uint

type genStatus struct {
	err error
}

const (
	ProjectInfoLevel = iota
	InfraListLevel
	InfraSetupLevel
	GenerationLevel
)

type ProjectForms struct {
	spinner  spinner.Model
	quitting bool
	err      error

	level        ProjectLevels
	prInfo       project.ProjectInfo
	infrasForm   *huh.Form
	infrasConfig []project.InfraConfig
	infrasJson   []project.JSONInfra
}

func NewProjectForms() ProjectForms {
	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	infras := GetInfras()
	opts := make([]huh.Option[basic.RouterModel], len(infras))
	i := 0

	for name, infra := range infras {
		opts[i] = huh.NewOption(name, infra)
		i++
	}

	infrasForm := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[basic.RouterModel]().
			Options(opts...).
			Title("Select infrastructures you want to add to your project:").
			Key("infras"),
	))
	return ProjectForms{
		spinner:  s,
		quitting: false,
		err:      nil,

		infrasForm: infrasForm,
		level:      ProjectInfoLevel,
	}
}

func (pjfs ProjectForms) Init() tea.Cmd {
	return nil
}

func (pjfs ProjectForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch pjfs.level {
	case ProjectInfoLevel:
		return pjfs, basic.SignalRouter(
			NewInfoForm(),
			basic.Next,
			nil,
		)

	case InfraListLevel:
		newForm, cmd := pjfs.infrasForm.Update(msg)
		pjfs.infrasForm = newForm.(*huh.Form)

		if pjfs.infrasForm.State == huh.StateCompleted {
			infras := pjfs.infrasForm.Get("infras").([]basic.RouterModel)
			pjfs.level++
			return pjfs, basic.SignalRouter(
				NewInfrasForms(pjfs.prInfo, infras),
				basic.Next,
				nil,
			)
		} else {
			return pjfs, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			pjfs.quitting = true
			return pjfs, basic.SignalRouter(nil, basic.Quit, nil)
		default:
			var cmd tea.Cmd
			pjfs.spinner, cmd = pjfs.spinner.Update(msg)
			return pjfs, cmd
		}
	case genStatus:
		pjfs.quitting = true
		pjfs.err = msg.err
		return pjfs, basic.SignalRouter(nil, basic.Quit, nil)
	}
	return pjfs, nil
}

func (pjfs ProjectForms) View() string {
	switch pjfs.level {
	case InfraListLevel:
		return pjfs.infrasForm.View()
	}

	var str string

	if pjfs.err != nil {
		str = pjfs.err.Error()
	}

	if pjfs.quitting {
		str += lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#76FF03")).Render("Generation completed!") + "\n"
	} else {
		str = fmt.Sprintf("%s Loading...", pjfs.spinner.View())
	}

	return str
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
		return pjfs, pjfs.infrasForm.Init()
	case InfraSetupLevel:
		pjfs.infrasConfig = params["infras_config"].([]project.InfraConfig)
		pjfs.infrasJson = params["infras_json"].([]project.JSONInfra)
		pjfs.level++

		proj := project.NewProject(pjfs.prInfo, pjfs.infrasConfig, pjfs.infrasJson)

		return pjfs, tea.Batch(pjfs.spinner.Tick, func() tea.Msg {
			return genStatus{proj.Generate()}
		})
	}
	return pjfs, pjfs.spinner.Tick
}
