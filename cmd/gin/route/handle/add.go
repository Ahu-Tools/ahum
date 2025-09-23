/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package handle

import (
	"github.com/Ahu-Tools/AhuM/pkg/gin"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [version_name] [entity_name] [method_name]",
	Short: "Add a new handler method to a specific Gin route entity",
	Long:  "The 'add' command creates a new handler method for a specified entity within a Gin route version, generating the necessary files and configurations.",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		pj, err := project.LoadProject(viper.GetString("projectPath"))
		if err != nil {
			return err
		}

		g := gin.LoadGinFromProject(*pj)
		genGuide, err := pj.GetEdgeGenGuide(g)
		if err != nil {
			return err
		}

		return g.AddHandler(args[0], args[1], args[2], *genGuide)
	},
}
