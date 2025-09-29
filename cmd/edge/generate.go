/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package edge

import (
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	"github.com/Ahu-Tools/ahum/pkg/tui/edge"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new edge",
	Long: `Use the generate command to create a new edge in your project.
This command will guide you through the process of defining the edge's properties.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		form, err := edge.NewForm(viper.GetString("projectPath"))
		if err != nil {
			return err
		}
		model := basic.NewRouter(form)

		p := tea.NewProgram(model)
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
