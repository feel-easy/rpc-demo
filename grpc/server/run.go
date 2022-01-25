package server

import (
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {

}
