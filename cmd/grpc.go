/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	grpc_client "rpc-demo/grpc/client"
	grpc_server "rpc-demo/grpc/server"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc 的演示程序",
	Long:  `这是一个 grpc 的演示程序`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("grpc called")
	// },
}

func init() {
	grpcCmd.AddCommand(grpc_client.ClientCmd)
	grpcCmd.AddCommand(grpc_server.ServerCmd)
	rootCmd.AddCommand(grpcCmd)
}
