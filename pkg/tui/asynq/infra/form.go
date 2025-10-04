package infra

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Ahu-Tools/ahum/pkg/asynq/infra"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Form struct {
	form *huh.Form
	pj   project.ProjectInfo
}

func NewForm(pj project.ProjectInfo) *Form {
	cfg := infra.DefaultConfig()
	port := strconv.Itoa(cfg.Redis.Port)
	db := strconv.Itoa(cfg.Redis.DB)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your redis server host:").
				Validate(huh.ValidateNotEmpty()).
				Value(&cfg.Redis.Host).
				Key("host"),

			huh.NewInput().
				Title("Enter your redis server port:").
				Validate(func(s string) error {
					n, err := strconv.Atoi(s)
					if err == nil {
						if n > 0 && n <= 65535 {
							return nil
						}
					}
					return fmt.Errorf("invalid port number")
				}).
				Value(&port).
				Key("port"),

			huh.NewInput().
				Title("Enter your redis server username:").
				Value(&cfg.Redis.Username).
				Key("username"),

			huh.NewInput().
				Title("Enter your redis server password:").
				EchoMode(huh.EchoModePassword).
				Value(&cfg.Redis.Password).
				Key("password"),

			huh.NewInput().
				Title("Enter your redis database:").
				Value(&db).
				Validate(numValid).
				Key("db"),
		))
	return &Form{
		form: form,
		pj:   pj,
	}
}

func (f Form) Init() tea.Cmd {
	return f.form.Init()
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := f.form.Update(msg)
	f.form = form.(*huh.Form)

	switch f.form.State {
	case huh.StateAborted:
		return f, basic.SignalError(errors.New("form aborted"))
	case huh.StateCompleted:
		port, _ := strconv.Atoi(f.form.GetString("port"))
		db, _ := strconv.Atoi(f.form.GetString("db"))
		redis := infra.Redis{
			Host:     f.form.GetString("host"),
			Port:     port,
			Username: f.form.GetString("username"),
			Password: f.form.GetString("password"),
			DB:       db,
		}

		cfg := infra.Config{
			Redis: redis,
		}

		return f, basic.SignalRouter(
			nil,
			basic.Back,
			infra.NewAsynq(cfg, f.pj),
		)
	}

	//huh.form state normal
	return f, cmd
}

func (f Form) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) { return f, nil }

func (f Form) View() string {
	return f.form.View()
}

func numValid(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid number")
	}
	return nil
}
