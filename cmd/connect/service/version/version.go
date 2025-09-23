/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package version

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Commands for manaconnectg Connect server and routes",
	Long: `The 'connect' command provides a set of tools for manaconnectg the Connect web server within your application.
You can use it to add new routes, new versions, and perform other server-related tasks.`,
}

func init() {
	Cmd.AddCommand(addCmd)
}
