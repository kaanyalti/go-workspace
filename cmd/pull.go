package cmd

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

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
		fmt.Println("Initializing workspace")
		tmpl, err := template.New("initialize").Parse(initFile)
		if err != nil {
			log.Fatalln(err)
		}
		file, err := os.Create(".workspace.yaml")
		defer file.Close()

		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(file, nil)
	},
}
