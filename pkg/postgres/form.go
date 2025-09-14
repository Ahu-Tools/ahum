package postgres

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type PostrgesForm struct {
	config     PostgresConfig
	jsonConfig *PostgresJSONConfig
	form       *huh.Form
}

func NewPostgresForm() PostrgesForm {
	jsonConfig := DefaultPostgresJSONConfig()
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter your postgres username:").
			Validate(noEmptyStr).
			Value(&jsonConfig.User).
			Key("user"),

		huh.NewInput().
			Title("Enter your postgres password:").
			EchoMode(huh.EchoModePassword).
			Placeholder("Your password won't be shown").
			Validate(noEmptyStr).
			Value(&jsonConfig.Password).
			Key("password"),

		huh.NewInput().
			Title("Enter the database name for your project:").
			Validate(noEmptyStr).
			Value(&jsonConfig.DbName).
			Key("db_name"),

		huh.NewInput().
			Title("Enter the hostname of your postgres database:").
			Placeholder("Default is "+jsonConfig.Host).
			Validate(noEmptyStr).
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
			Validate(noEmptyStr).
			Value(&jsonConfig.SSLMode).
			Key("sslmode"),
	))

	return PostrgesForm{
		jsonConfig: jsonConfig,
		form:       form,
	}
}

func (pf PostrgesForm) Init() tea.Cmd {
	return tea.Batch(pf.form.Init(), func() tea.Msg {
		return tea.KeyMsg{Type: tea.KeyEnter}
	})
}

func (pf PostrgesForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if pf.form.State == huh.StateCompleted {
		return pf, basic.SignalRouter(
			nil,
			basic.Back,
			basic.MsgParams{
				"ok":          true,
				"config":      pf.config,
				"config_json": *pf.jsonConfig,
			},
		)
	} else if pf.form.State == huh.StateAborted {
		return pf, basic.SignalRouter(
			nil,
			basic.Back,
			basic.MsgParams{
				"ok": false,
			},
		)
	}
	form, cmd := pf.form.Update(msg)
	pf.form = form.(*huh.Form)

	return pf, cmd
}

func (pf PostrgesForm) View() string {
	return pf.form.View()
}

func (pf PostrgesForm) Inject(params basic.MsgParams) basic.RouterModel {
	projectInfo := params["project_info"].(project.ProjectInfo)
	pf.config = *NewPostgresConfig(projectInfo)

	return pf
}

func (pf PostrgesForm) Return(params basic.MsgParams) (basic.RouterModel, tea.Cmd) {
	return pf, nil
}

func noEmptyStr(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("Empty input is not allowed!")
	}
	return nil
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
