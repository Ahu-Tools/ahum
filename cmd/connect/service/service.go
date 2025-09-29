/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package service

import (
	"fmt"

	"github.com/Ahu-Tools/ahum/cmd/connect/service/method"
	"github.com/Ahu-Tools/ahum/cmd/connect/service/version"
	"github.com/spf13/cobra"
)

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage Connect services",
	Long: `The 'service' command provides subcommands for managing Connect services.
You can use it to add new services, versions, and methods.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
}

func init() {
	ServiceCmd.AddCommand(addCmd)
	ServiceCmd.AddCommand(version.Cmd)
	ServiceCmd.AddCommand(method.Cmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
