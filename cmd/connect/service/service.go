/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package service

import (
	"fmt"

	"github.com/Ahu-Tools/AhuM/cmd/connect/service/version"
	"github.com/spf13/cobra"
)

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Commands for managing Gin services",
	Long: `The 'service' command provides subcommands for managing services in your Gin application.
You can use it to add new versions to your API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
}

func init() {
	ServiceCmd.AddCommand(addCmd)
	ServiceCmd.AddCommand(version.Cmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
