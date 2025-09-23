/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package version

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Manage Connect service versions",
	Long: `The 'version' command provides subcommands for managing versions of Connect services.
You can use it to add new versions to an existing service.`,
}

func init() {
	Cmd.AddCommand(addCmd)
}
