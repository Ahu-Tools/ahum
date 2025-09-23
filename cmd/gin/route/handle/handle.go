/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package handle

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "handle",
	Short: "Manage Gin route handlers",
	Long:  "The 'handle' command provides subcommands to manage handlers within Gin routes, such as adding new handler methods.",
}

func init() {
	Cmd.AddCommand(addCmd)
}
