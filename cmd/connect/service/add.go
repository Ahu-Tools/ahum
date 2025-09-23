/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package service

import (
	"github.com/Ahu-Tools/AhuM/pkg/connect"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [service_name]",
	Short: "short descriptions",
	Long:  "long descriptions",
	Args:  cobra.ExactArgs(1),
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

		return c.AddService(args[0], *genGuide)
	},
}

func init() {

}
