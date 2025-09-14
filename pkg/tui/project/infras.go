package project

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tea "github.com/charmbracelet/bubbletea"
)

type InfrasForms struct {
	prInfo       project.ProjectInfo
	infras       []basic.RouterModel
	infrasConfig []project.InfraConfig
	infrasJson   []project.JSONInfra
	infrasStep   int
}

func NewInfrasForms(prInfo project.ProjectInfo, infras []basic.RouterModel) InfrasForms {
	return InfrasForms{
		prInfo:       prInfo,
		infras:       infras,
		infrasStep:   0,
		infrasConfig: make([]project.InfraConfig, len(infras)),
		infrasJson:   make([]project.JSONInfra, len(infras)),
	}
}

func (pjfs InfrasForms) Init() tea.Cmd {
	return func() tea.Msg { return "hello" }
}

func (pjfs InfrasForms) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Pass through messages to the current infrastructure form
	return pjfs, basic.SignalRouter(
		pjfs.infras[pjfs.infrasStep],
		basic.Next,
		basic.MsgParams{
			"project_info": pjfs.prInfo,
		},
	)
}

func (pjfs InfrasForms) View() string {
	return ""
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
