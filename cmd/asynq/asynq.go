package asynq

import (
	"github.com/Ahu-Tools/ahum/cmd/asynq/edge"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "asynq",
	Short: "short description",
	Long:  "long description",
}

func init() {
	Cmd.AddCommand(edge.Cmd)
}
