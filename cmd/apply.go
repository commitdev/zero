package cmd

import (
	"log"
	"os"

	"github.com/commitdev/zero/internal/apply"
	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/spf13/cobra"
)

var applyConfigPath string
var applyEnvironments []string

func init() {
	applyCmd.PersistentFlags().StringVarP(&applyConfigPath, "config", "c", constants.ZeroProjectYml, "config path")
	applyCmd.PersistentFlags().StringSliceVarP(&applyEnvironments, "env", "e", []string{}, "environments to set up (staging, production) - specify multiple times for multiple")

	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Execute modules to create projects, infrastructure, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			rootDir = projectconfig.RootDir
		}
		apply.Apply(rootDir, applyConfigPath, applyEnvironments)
	},
}
