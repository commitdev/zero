package cmd

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate"
	"github.com/commitdev/commit0/internal/util"
	"github.com/spf13/cobra"
)

var configPath string

func init() {
	generate2Cmd.PersistentFlags().StringVarP(&configPath, "config", "c", util.CommitYml, "config path")

	rootCmd.AddCommand(generateCmd)
}

var generate2Cmd = &cobra.Command{
	Use:   "generate2",
	Short: "Generate idl & application folders",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadGeneratorConfig(configPath)
		// cfg.Print()

		generate.Generate(cfg)
	},
}
