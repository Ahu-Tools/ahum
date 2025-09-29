/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package connect

import (
	"github.com/Ahu-Tools/ahum/pkg/connect"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Connect related files (e.g., protobuf, Go stubs)",
	Long: `The 'gen' command generates various files required for Connect, including protobuf definitions and Go stubs.
It uses the project configuration to determine what needs to be generated.`,
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
