/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package method

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "method",
	Short: "Manage Connect service methods",
	Long: `The 'method' command provides subcommands for managing methods within Connect services.
You can use it to add new methods to an existing service.`,
}

func init() {
	Cmd.AddCommand(addCmd)
}
