package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero gets you to writing code quicker.",
	Long:  "Zero is an open-source developer platform CLI tool which makes it quick and easy for technical founders & developers \nto build quality and reliable infrastructure to launch, grow and scale high-quality SaaS applications faster and more cost-effectively.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
