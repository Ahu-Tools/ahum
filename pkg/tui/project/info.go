package project

import (
	"errors"
	"runtime"
	"strings"

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
	// Get the full version string, e.g., "go1.22.5"
	fullVersionString := runtime.Version()

	// Remove the "go" prefix to get just the number part
	versionNumber := strings.TrimPrefix(fullVersionString, "go")

	info := project.NewProjectInfo("", versionNumber, ".")
	form := huh.NewForm(huh.NewGroup(

		huh.NewInput().
			Title("Enter the package name of your new project:").
			Validate(huh.ValidateNotEmpty()).
			Value(&info.PackageName).
			Key("packageName"),

		huh.NewInput().
			Title("Enter the Go version of your new project:").
			Validate(huh.ValidateNotEmpty()).
			Value(&info.GoVersion).
			Key("goVersion"),

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
	switch m.form.State {
	case huh.StateAborted:
		return m, basic.SignalError(errors.New("form aborted"))
	case huh.StateNormal:
		newForm, cmd := m.form.Update(msg)
		m.form = newForm.(*huh.Form)
		return m, cmd
	}

	return m, basic.SignalRouter(
		nil,
		basic.Back,
		ProjectInfoMsg{
			ProjectInfo: *m.info,
		},
	)
}

func (m InfoForm) View() string {
	return m.form.View()
}

func (info InfoForm) Inject(msg tea.Msg) basic.RouterModel {
	return info
}

func (info InfoForm) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	return info, nil
}
