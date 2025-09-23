/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package method

import (
	"github.com/Ahu-Tools/AhuM/pkg/connect"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [method_name] [service_name] [version_name]",
	Short: "short descriptions",
	Long:  "long descriptions",
	Args:  cobra.ExactArgs(3),
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

		return c.AddMethod(args[0], args[1], args[2], *genGuide)
	},
}

func init() {

}
