package project

import (
	"log"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type initModel struct {
	form       *huh.Form
	Proj       project.Project
	Aborted    bool
	spinner    spinner.Model
	generating bool
	err        error
}

type projectGeneratedMsg struct{ err error }

func NewInitModel() initModel {
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter the name of your new project:").
			Validate(validateName).
			Key("name"),

		huh.NewInput().
			Title("Enter the package name of your new project:").
			Validate(validateName).
			Key("packageName"),

		huh.NewInput().
			Title("Enter the root path of your new project:").
			Placeholder("Default is (.)").
			Validate(validatePath).
			Key("rootPath"),

		huh.NewMultiSelect[project.Edge]().
			Options(
				huh.NewOption("Connect", project.CONNECT),
				huh.NewOption("Gin", project.GIN),
			).
			Title("Choose edges:").
			Validate(validateEdges).
			Key("edges"),

		huh.NewMultiSelect[project.Database]().
			Options(
				huh.NewOption("PostgreSQL", project.POSTGRES),
			).
			Title("Choose databases:").
			Validate(validateDbs).
			Key("databases"),
	))
	s := spinner.New()
	s.Spinner = spinner.Dot
	return initModel{
		form:    form,
		Aborted: false,
		spinner: s,
	}
}

func (m initModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Aborted {
		return m, tea.Quit
	}

	if m.generating {
		switch msg := msg.(type) {
		case projectGeneratedMsg:
			if msg.err != nil {
				log.Fatal(msg.err)
				m.err = msg.err
			}
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		m.generating = true
		m.Proj.Name = m.form.GetString("name")
		m.Proj.PackageName = m.form.GetString("packageName")

		m.Proj.RootPath = m.form.GetString("rootPath")
		if m.Proj.RootPath == "" {
			m.Proj.RootPath = "."
		}

		m.Proj.Edges = m.form.Get("edges").([]project.Edge)
		m.Proj.Dbs = m.form.Get("databases").([]project.Database)

		return m, tea.Batch(m.spinner.Tick, func() tea.Msg {
			err := m.Proj.Generate()
			return projectGeneratedMsg{err}
		})
	}

	if m.form.State == huh.StateAborted {
		m.Aborted = true
	}

	return m, cmd
}

func (m initModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}
	if m.Aborted {
		return "Aborted"
	}
	if m.generating {
		return m.spinner.View() + " Generating project..."
	}
	return m.form.View()
}
