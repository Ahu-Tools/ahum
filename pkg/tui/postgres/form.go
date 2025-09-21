package postgres

import (
	"errors"
	"strconv"

	"github.com/Ahu-Tools/AhuM/pkg/postgres"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type DoneFormMsg struct{}
type AbortedFormMsg struct{}

type PostrgesForm struct {
	jsonConfig *postgres.PostgresConfig
	pjInfo     project.ProjectInfo
	form       *huh.Form
}

func NewPostgresForm(p project.ProjectInfo) *PostrgesForm {
	jsonConfig := postgres.DefaultPostgresConfig()
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter your postgres username:").
			Validate(huh.ValidateNotEmpty()).
			Value(&jsonConfig.User).
			Key("user"),

		huh.NewInput().
			Title("Enter your postgres password:").
			EchoMode(huh.EchoModePassword).
			Placeholder("Your password won't be shown").
			Validate(huh.ValidateNotEmpty()).
			Value(&jsonConfig.Password).
			Key("password"),

		huh.NewInput().
			Title("Enter the database name for your project:").
			Validate(huh.ValidateNotEmpty()).
			Value(&jsonConfig.DbName).
			Key("db_name"),

		huh.NewInput().
			Title("Enter the hostname of your postgres database:").
			Placeholder("Default is "+jsonConfig.Host).
			Validate(huh.ValidateNotEmpty()).
			Value(&jsonConfig.Host).
			Key("host"),

		huh.NewInput().
			Title("Enter the port which your postgres instance is listening to:").
			Placeholder("Default is "+jsonConfig.Port).
			Validate(checkPortRange).
			Value(&jsonConfig.Port).
			Key("port"),

		huh.NewInput().
			Title("Enter the sslmode of your postgres connection:").
			Placeholder("Default is "+jsonConfig.SSLMode).
			Validate(huh.ValidateNotEmpty()).
			Value(&jsonConfig.SSLMode).
			Key("sslmode"),
	))

	form.SubmitCmd = func() tea.Msg { return DoneFormMsg{} }
	form.CancelCmd = func() tea.Msg { return AbortedFormMsg{} }

	return &PostrgesForm{
		jsonConfig: jsonConfig,
		form:       form,

		pjInfo: p,
	}
}

func (pf PostrgesForm) Init() tea.Cmd {
	return pf.form.Init()
}

func (pf PostrgesForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case DoneFormMsg:
		return pf, basic.SignalRouter(
			nil,
			basic.Back,
			postgres.NewPostgres(pf.pjInfo, *pf.jsonConfig),
		)
	case AbortedFormMsg:
		return pf, basic.SignalError(errors.New("form aborted"))
	}

	newForm, cmd := pf.form.Update(msg)
	pf.form = newForm.(*huh.Form)
	return pf, cmd
}

func (pf PostrgesForm) View() string {
	return pf.form.View()
}

func (pf PostrgesForm) Return(params tea.Msg) (basic.RouterModel, tea.Cmd) {
	return pf, nil
}

func checkPortRange(portStr string) error {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return errors.New("port must be a positive integer number")
	}

	if port > 65535 || port <= 0 {
		return errors.New("provided port is not in the valid range 0-65535")
	}

	return nil
}
