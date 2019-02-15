package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blockchain cli",
	Short: "blockchain cli",
	Long:  `blockchain cli`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute ... execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
