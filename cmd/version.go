package cmd

import (
	"fmt"

	"github.com/commitdev/zero/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of zero",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %v\n", version.AppVersion)
		fmt.Printf("build: %v\n", version.AppBuild)
	},
}
