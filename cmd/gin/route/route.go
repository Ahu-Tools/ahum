/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package route

import (
	"fmt"

	"github.com/Ahu-Tools/AhuM/cmd/gin/route/entity"
	"github.com/Ahu-Tools/AhuM/cmd/gin/route/handle"
	"github.com/Ahu-Tools/AhuM/cmd/gin/route/version"
	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var RouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Commands for managing Gin routes",
	Long: `The 'route' command provides subcommands for managing routes in your Gin application.
You can use it to add new versions to your API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("route called")
	},
}

func init() {
	RouteCmd.AddCommand(version.VersionCmd)
	RouteCmd.AddCommand(entity.Cmd)
	RouteCmd.AddCommand(handle.Cmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
