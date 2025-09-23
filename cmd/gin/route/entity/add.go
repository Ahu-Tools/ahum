/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package entity

import (
	"github.com/Ahu-Tools/AhuM/pkg/gin"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [version_name] [entity_name]",
	Short: "short descriptions",
	Long:  "long descriptions",
	Args:  cobra.ExactArgs(2),
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

		return g.AddEntity(args[0], args[1], *genGuide)
	},
}
