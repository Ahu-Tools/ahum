package edge

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var EdgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "A command for managing edges",
	Long:  `The edge command provides tools for managing and generating edges within your project.`,
}

func init() {
	EdgeCmd.AddCommand(genCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// edgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// edgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	EdgeCmd.PersistentFlags().StringP("path", "p", ".", "project root path")
	viper.BindPFlag("projectPath", EdgeCmd.PersistentFlags().Lookup("path"))
}
