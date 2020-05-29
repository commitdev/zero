package cmd

import (
	"github.com/commitdev/zero/configs"
	"github.com/spf13/cobra"
)

var applyConfigPath string

func init() {
	applyCmd.PersistentFlags().StringVarP(&applyConfigPath, "config", "c", configs.ZeroProjectYml, "config path")

	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Execute modules to create projects, infrastructure, etc.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
