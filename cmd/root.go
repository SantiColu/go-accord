package cmd

import (
	"fmt"
	"os"

	"github.com/SantiColu/go-accord/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "accord",
	Short: "accord is a self-hosted discord-like cli chat... and useless",
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
