package redis

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/redis"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Form struct {
	form *huh.Form
	pj   project.ProjectInfo
}

func NewForm(pj project.ProjectInfo) *Form {
	cfg := redis.DefaultConfig()
	port := strconv.Itoa(cfg.Port)
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter your redis server host:").
			Validate(huh.ValidateNotEmpty()).
			Value(&cfg.Host).
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
			Value(&cfg.Username).
			Key("username"),

		huh.NewInput().
			Title("Enter your redis server password:").
			EchoMode(huh.EchoModePassword).
			Value(&cfg.Password).
			Key("password"),
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
		cfg := redis.NewConfig(
			f.form.GetString("host"),
			port,
			f.form.GetString("username"),
			f.form.GetString("password"),
		)

		return f, basic.SignalRouter(
			nil,
			basic.Back,
			redis.NewRedis(f.pj, *cfg),
		)
	}

	//huh.form state normal
	return f, cmd
}

func (f Form) Return(msg tea.Msg) (basic.RouterModel, tea.Cmd) { return f, nil }

func (f Form) View() string {
	return f.form.View()
}
