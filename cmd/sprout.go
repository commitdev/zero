package cmd

import (
	"fmt"
	"github.com/commitdev/sprout/templator"
	"github.com/spf13/cobra"
	"os"
)

var Templator *templator.Templator

var rootCmd = &cobra.Command{
	Use:   "sprout",
	Short: "Sprout a moduler service generator.",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute(templates *templator.Templator) {
	Templator = templates
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
