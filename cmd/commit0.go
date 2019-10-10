package cmd

import (
	"fmt"
	"github.com/commitdev/commit0/templator"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"os"
)

var Templator *templator.Templator

var rootCmd = &cobra.Command{
	Use:   "commit0",
	Short: "Commit0 is a moduler service generator.",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Init() {
	templates := packr.New("templates", "../templates")
	Templator = templator.NewTemplator(templates)
}

func Execute() {
	Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
