package cmd

import (
	"github.com/commitdev/commit0/internal/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commit0api)
}

var commit0api = &cobra.Command{
	Use:   "api",
	Short: "Run Commit0 Api",
	Run: func(cmd *cobra.Command, args []string) {
		api.Commit0Api()
	},
}
