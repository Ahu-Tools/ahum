package edge

import (
	"fmt"

	asynqedge "github.com/Ahu-Tools/ahum/pkg/asynq/edge"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "short description",
	Long:  "long description",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("invalid option choosed")
		}

		p, err := project.LoadProject(viper.GetString("projectPath"))
		if err != nil {
			return err
		}

		edges := p.GetEdgesByName()
		edge := edges[asynqedge.Name].(*asynqedge.Asynq)
		edgeGuide, err := p.GetEdgeGenGuide(edge)
		if err != nil {
			return err
		}

		switch args[0] {
		case "module":
			if len(args) < 3 {
				return fmt.Errorf("no version or module name specified")
			}
			err = edge.AddModule(args[1], args[2], *edgeGuide)

		case "task":
			if len(args) < 4 {
				return fmt.Errorf("no task or module or version name specified")
			}
			err = edge.AddTaskHandler(args[1], args[2], args[3], *edgeGuide)
		}
		return err
	},
}

func init() {

}
