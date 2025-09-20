/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package gin

import (
	"fmt"

	"github.com/Ahu-Tools/AhuM/cmd/gin/route"
	"github.com/spf13/cobra"
)

// ginCmd represents the gin command
var GinCmd = &cobra.Command{
	Use:   "gin",
	Short: "Commands for managing Gin server and routes",
	Long: `The 'gin' command provides a set of tools for managing the Gin web server within your application.
You can use it to add new routes, new versions, and perform other server-related tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gin called")
	},
}

func init() {
	GinCmd.AddCommand(route.RouteCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
