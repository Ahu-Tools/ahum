/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package infra

import (
	"github.com/spf13/cobra"
)

// InfraCmd represents the infra command
var Cmd = &cobra.Command{
	Use:   "infra",
	Short: "Manage infrastructure components",
	Long:  `The infra command provides subcommands to manage different aspects of the project's infrastructure.`,
}

func init() {
	Cmd.AddCommand(genCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
