/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"rpc-demo/rpc"

	"github.com/spf13/cobra"
)

// rpcCmd represents the rpc command
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "rpc 演示程序",
	Long:  `这是一个rpc 的演示程序`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("rpc called")
	// },
}

func init() {
	rpcCmd.AddCommand(rpc.ClientCmd)
	rpcCmd.AddCommand(rpc.ServerCmd)
	rootCmd.AddCommand(rpcCmd)
}
