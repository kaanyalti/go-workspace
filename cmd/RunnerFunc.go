package cmd

import "github.com/spf13/cobra"

type RunnerFunc func() func(cmd *cobra.Command, args []string)
