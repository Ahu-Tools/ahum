/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package entity

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "entity",
	Short: "Manage Gin route entities",
	Long:  "The 'entity' command provides subcommands to manage entities within Gin routes, such as adding new entities.",
}

func init() {
	Cmd.AddCommand(addCmd)
}
