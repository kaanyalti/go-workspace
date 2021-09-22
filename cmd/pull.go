package cmd

import (
	_ "embed"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull directories",
	Long:  "test",
	Run: func(cmd *cobra.Command, args []string) {
		var s storage.Storer
		git.Clone()
	},
}
