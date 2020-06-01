package cmd

import (
	"fmt"
	"log"

	"github.com/commitdev/zero/configs"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var applyConfigPath string
var applyEnvironments []string

func init() {
	applyCmd.PersistentFlags().StringVarP(&applyConfigPath, "config", "c", configs.ZeroProjectYml, "config path")
	applyCmd.PersistentFlags().StringSliceVarP(&applyEnvironments, "env", "e", []string{}, "environments to set up (staging, production) - specify multiple times for multiple")

	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Execute modules to create projects, infrastructure, etc.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(applyEnvironments) == 0 {
			fmt.Println(`Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.`)
			applyEnvironments = promptEnvironments()
		}

		// Strict for now, we can brainstorm how much we want to support custom environments later
		for _, env := range applyEnvironments {
			if env != "staging" && env != "production" {
				exit.Fatal("The currently supported environments are \"staging\" and \"production\"")
			}
		}

		// @TODO : Pass environments to make commands
	},
}

// promptEnvironments Prompts the user for the environments to apply against and returns a slice of strings representing the environments
func promptEnvironments() []string {
	items := map[string][]string{
		"Staging ":                    []string{"staging"},
		"Production":                  []string{"production"},
		"Both Staging and Production": []string{"staging", "production"},
	}

	var labels []string
	for label := range items {
		labels = append(labels, label)
	}

	providerPrompt := promptui.Select{
		Label: "Environments",
		Items: labels,
	}
	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}
	return items[providerResult]
}
