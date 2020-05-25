package cmd

import (
	"github.com/commitdev/zero/internal/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(zeroApi)
}

var zeroApi = &cobra.Command{
	Use:   "ui",
	Short: "Run zero Api",
	Run: func(cmd *cobra.Command, args []string) {
		api.Commit0Api()
	},
}
