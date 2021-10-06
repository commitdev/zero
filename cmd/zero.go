package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero gets you to writing code quicker.",
	Long:  "Zero is an open-source CLI tool which makes it quick and easy for technical founders & developers \nto build high-quality, reliable infrastructure to launch, grow, and scale production-ready SaaS applications faster and more cost-effectively.\n https://getzero.dev\n",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if len(os.Args) > 1 {
		if err := rootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	} else { // If no arguments were provided, print the usage message.
		rootCmd.Help()
	}
}
