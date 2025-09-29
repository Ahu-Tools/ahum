/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package create

import (
	"fmt"
	"os"

	"github.com/Ahu-Tools/ahum/pkg/tui/service"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var serviceCreatePath string

// serviceCreateCmd represents the service create command
var CreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new service in an Ahu project",
	Long:    `Create a new service in an already initialised Ahu project.`,
	PreRunE: preRunE,
	RunE:    runE,
}

func init() {
	CreateCmd.Flags().StringVarP(&serviceCreatePath, "path", "p", ".", "Path to the project")
}

func preRunE(cmd *cobra.Command, args []string) error {
	fileInfo, err := os.Stat(serviceCreatePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path '%s' does not exist", serviceCreatePath)
		}
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("path '%s' is not a directory", serviceCreatePath)
	}
	return nil
}

func runE(cmd *cobra.Command, args []string) error {
	model := service.NewServiceForm(serviceCreatePath)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
