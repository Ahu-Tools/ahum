package project

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type InfoForm struct {
	form *huh.Form
	info *project.ProjectInfo
}

func NewInfoForm() InfoForm {
	info := project.NewProjectInfo("", "", ".")
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter the name of your new project:").
			Validate(validateName).
			Value(&info.Name).
			Key("name"),

		huh.NewInput().
			Title("Enter the package name of your new project:").
			Validate(validateName).
			Value(&info.PackageName).
			Key("packageName"),

		huh.NewInput().
			Title("Enter the root path of your new project:").
			Placeholder("Default is (.)").
			Validate(validatePath).
			Value(&info.RootPath).
			Key("rootPath"),
	))

	return InfoForm{
		form,
		info,
	}
}

func (m InfoForm) Init() tea.Cmd {
	return m.form.Init()
}

func (m InfoForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		return m, basic.SignalRouter(
			nil,
			basic.Back,
			basic.MsgParams{
				"ok":           true,
				"project_info": *m.info,
			},
		)
	} else if m.form.State == huh.StateAborted {
		return m, basic.SignalRouter(
			nil,
			basic.Back,
			basic.MsgParams{
				"ok": false,
			},
		)
	}

	return m, cmd
}

func (m InfoForm) View() string {
	return m.form.View()
}

func (info InfoForm) Inject(params basic.MsgParams) basic.RouterModel {
	return info
}

func (info InfoForm) Return(params basic.MsgParams) (basic.RouterModel, tea.Cmd) {
	return info, nil
}
