/*
Copyright © 2025 Sina Sadeghi sina.sadeghi83@gmail.com
*/
package entity

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "entity",
	Short: "short descriptions",
	Long:  "long descriptions",
}

func init() {
	Cmd.AddCommand(addCmd)
}
