package start

import (
	"github.com/mittacy/ego-layout/cmd/start/http"
	"github.com/mittacy/ego-layout/cmd/start/job"
	"github.com/mittacy/ego-layout/cmd/start/task"
	"github.com/spf13/cobra"
)

// Cmd api server
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start server command",
	Long:  "start server command",
	Run:   run,
}

func init() {
	Cmd.AddCommand(http.Cmd)
	Cmd.AddCommand(job.Cmd)
	Cmd.AddCommand(task.Cmd)
}

func run(cmd *cobra.Command, args []string) {
}
