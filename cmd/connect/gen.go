/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package connect

import (
	"github.com/Ahu-Tools/AhuM/pkg/connect"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Commands for manaconnectg Connect server and routes",
	Long: `The 'connect' command provides a set of tools for manaconnectg the Connect web server within your application.
You can use it to add new routes, new versions, and perform other server-related tasks.`,
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

		return c.BufGenerate(*genGuide)
	},
}

func init() {

}
