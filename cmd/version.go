package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion = "SNAPSHOT"
	appBuild   = "SNAPSHOT"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of zero",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %v\n", appVersion)
		fmt.Printf("build: %v\n", appBuild)
	},
}
