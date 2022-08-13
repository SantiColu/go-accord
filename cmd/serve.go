package cmd

import (
	"github.com/SantiColu/go-accord/pkg/server"
	"github.com/spf13/cobra"
)

var Port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the server that will host the chat",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(Port)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&Port, "port", "p", "", "Port to run on")
	serveCmd.MarkFlagRequired("port")
	rootCmd.AddCommand(serveCmd)
}
