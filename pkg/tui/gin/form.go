package gin

import (
	"errors"

	"github.com/Ahu-Tools/AhuM/pkg/gin"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Form struct {
	form *huh.Form

	pj project.ProjectInfo
}

func NewForm() *Form {
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter your gin server host:").
			Validate(huh.ValidateNotEmpty()).
			Key("host"),

		huh.NewInput().
			Title("Enter your gin server port:").
			Validate(huh.ValidateNotEmpty()).
			Key("port"),
	))

	return &Form{
		form: form,
	}
}

func (f *Form) InitProjectInfo(pj project.ProjectInfo) {
	f.pj = pj
}

func (f Form) Init() tea.Cmd {
	return f.form.Init()
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := f.form.Update(msg)
	f.form = model.(*huh.Form)

	switch f.form.State {
	case huh.StateCompleted:
		host := f.form.Get("host").(string)
		port := f.form.Get("port").(string)
		return f, basic.SignalRouter(
			nil,
			basic.Back,
			gin.NewGin(&f.pj, gin.GinConfig{Server: gin.GinServer{Host: host, Port: port}}),
		)
	case huh.StateAborted:
		return f, basic.SignalError(errors.New("form aborted"))
	}

	return f, cmd
}

func (f Form) View() string {
	return f.form.View()
}

func (f Form) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	return f, nil
}
