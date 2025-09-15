/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package service

import (
	"github.com/Ahu-Tools/AhuM/cmd/service/create"
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage services in an Ahu project",
	Long:  `Manage services in an already initialised Ahu project.`,
}

func init() {
	ServiceCmd.AddCommand(create.CreateCmd)
}
