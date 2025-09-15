/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package cmd

import (
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	tproj "github.com/Ahu-Tools/AhuM/pkg/tui/project"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise an Ahu project",
	Long:  `Create folders and go files related to the main architecture of the Ahu project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		model := tproj.NewProjectForms()
		router := basic.NewRouter(model)
		p := tea.NewProgram(router)
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
