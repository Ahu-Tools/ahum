/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package connect

import (
	"github.com/Ahu-Tools/ahum/cmd/connect/service"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var Cmd = &cobra.Command{
	Use:   "connect",
	Short: "Manage Connect server and routes",
	Long: `The 'connect' command provides a set of tools for managing the Connect server within your application.
You can use it to add new routes, new versions, and perform other server-related tasks.`,
}

func init() {
	Cmd.AddCommand(service.ServiceCmd)
	Cmd.AddCommand(genCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
