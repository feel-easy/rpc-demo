package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rpcdemo",
	Short: "关于rpc的演示程序",
	Long:  `关于rpc的一些演示程序`,
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

func Execute() {
	rootCmd.Execute()
}
