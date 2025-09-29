/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package version

import (
	"github.com/Ahu-Tools/ahum/pkg/connect"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [version_name] [service_name]",
	Short: "Add a new version to a Connect service",
	Long: `The 'add' command creates a new version for a specified Connect service.
This will generate the necessary files and update the project configuration.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pj, err := project.LoadProject(viper.GetString("projectPath"))
		if err != nil {
			return err
		}

		c := connect.LoadConnectFromProject(*pj)
		genGuide, err := pj.GetEdgeGenGuide(c)
		if err != nil {
			return err
		}

		return c.AddVersion(args[0], args[1], *genGuide)
	},
}
