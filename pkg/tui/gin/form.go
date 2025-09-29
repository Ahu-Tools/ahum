package gin

import (
	"errors"

	"github.com/Ahu-Tools/ahum/pkg/gin"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type DoneFormMsg struct{}
type AbortedFormMsg struct{}

type Form struct {
	form *huh.Form

	pj project.ProjectInfo
}

func NewForm(pj project.ProjectInfo) *Form {
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

	form.SubmitCmd = func() tea.Msg { return DoneFormMsg{} }
	form.CancelCmd = func() tea.Msg { return AbortedFormMsg{} }

	return &Form{
		form: form,
		pj:   pj,
	}
}

func (f Form) Init() tea.Cmd {
	return f.form.Init()
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case DoneFormMsg:
		host := f.form.Get("host").(string)
		port := f.form.Get("port").(string)
		return f, basic.SignalRouter(
			nil,
			basic.Back,
			gin.NewGin(&f.pj, gin.GinConfig{Server: gin.GinServer{Host: host, Port: port}}),
		)
	case AbortedFormMsg:
		return f, basic.SignalError(errors.New("form aborted"))
	}

	model, cmd := f.form.Update(msg)
	f.form = model.(*huh.Form)

	return f, cmd
}

func (f Form) View() string {
	return f.form.View()
}

func (f Form) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) {
	return f, nil
}
