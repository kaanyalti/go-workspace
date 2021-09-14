package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "workspace",
	Short: "generates workspace to manage multiple git repos in your dev env",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HELLO WORLD")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
