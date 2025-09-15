package project

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type InfrasForms struct {
	form *huh.Form

	prInfo       project.ProjectInfo
	infras       []basic.RouterModel
	infrasConfig []project.InfraConfig
	infrasJson   []project.JSONInfra
	infrasStep   int
}

func NewInfrasForms(prInfo project.ProjectInfo) InfrasForms {
	allInfras := GetInfras()
	opts := make([]huh.Option[basic.RouterModel], len(allInfras))
	i := 0

	for name, infra := range allInfras {
		opts[i] = huh.NewOption(name, infra)
		i++
	}

	infrasForm := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[basic.RouterModel]().
			Options(opts...).
			Title("Select infrastructures you want to add to your project:").
			Key("infras"),
	))

	return InfrasForms{
		form: infrasForm,

		prInfo:     prInfo,
		infrasStep: 0,
	}
}

func (pjfs InfrasForms) Init() tea.Cmd {
	return pjfs.form.Init()
}

func (pjfs InfrasForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newForm, cmd := pjfs.form.Update(msg)
	pjfs.form = newForm.(*huh.Form)

	if pjfs.form.State == huh.StateCompleted {
		pjfs.infras = pjfs.form.Get("infras").([]basic.RouterModel)

		pjfs.infrasConfig = make([]project.InfraConfig, len(pjfs.infras))
		pjfs.infrasJson = make([]project.JSONInfra, len(pjfs.infras))

		if len(pjfs.infras) == 0 {
			return pjfs, basic.SignalRouter(
				nil,
				basic.Back,
				basic.MsgParams{
					"infras_config": pjfs.infrasConfig,
					"infras_json":   pjfs.infrasJson,
				},
			)
		}
		// Pass through messages to the current infrastructure form
		return pjfs, basic.SignalRouter(
			pjfs.infras[pjfs.infrasStep],
			basic.Next,
			basic.MsgParams{
				"project_info": pjfs.prInfo,
			},
		)
	}

	return pjfs, cmd
}

func (pjfs InfrasForms) View() string {
	return pjfs.form.View()
}

func (pjfs InfrasForms) Inject(params basic.MsgParams) basic.RouterModel {
	return pjfs
}

func (pjfs InfrasForms) Return(params basic.MsgParams) (basic.RouterModel, tea.Cmd) {
	pjfs.infrasConfig[pjfs.infrasStep] = params["config"].(project.InfraConfig)
	pjfs.infrasJson[pjfs.infrasStep] = params["config_json"].(project.JSONInfra)
	pjfs.infrasStep++

	if pjfs.infrasStep == len(pjfs.infras) {
		return pjfs, basic.SignalRouter(
			nil,
			basic.Back,
			basic.MsgParams{
				"infras_config": pjfs.infrasConfig,
				"infras_json":   pjfs.infrasJson,
			},
		)
	}

	return pjfs, basic.SignalRouter(
		pjfs.infras[pjfs.infrasStep],
		basic.Next,
		basic.MsgParams{
			"project_info": pjfs.prInfo,
		},
	)
}
