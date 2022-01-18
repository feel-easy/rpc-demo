package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rpc",
	Short: "rpc",
	Long:  `rpc`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.Execute()
}
