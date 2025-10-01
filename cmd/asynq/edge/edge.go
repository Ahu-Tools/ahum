package edge

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "edge",
	Short: "short description",
	Long:  "long description",
}

func init() {
	Cmd.AddCommand(addCmd)
}
