/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package version

import (
	"github.com/Ahu-Tools/AhuM/pkg/gin"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [version]",
	Short: "Add a new version to the Gin server",
	Long: `Add a new version to the Gin server.
This will create a new version directory and a registrar file.
You must provide the version name as an argument.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pj, err := project.LoadProject(viper.GetString("projectPath"))
		if err != nil {
			return err
		}

		g, err := gin.LoadGinFromProject(*pj)
		if err != nil {
			return err
		}

		genGuide, err := pj.GetEdgeGenGuide(g)
		if err != nil {
			return err
		}

		return g.AddVersion(args[0], *genGuide)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
