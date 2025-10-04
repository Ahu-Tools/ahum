package edge

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "edge",
	Short: "Manage asynq edges",
	Long:  `Manage asynq edges.`,
}

func init() {
	Cmd.AddCommand(addCmd)
}
