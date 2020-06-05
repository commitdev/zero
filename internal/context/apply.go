package context

import (
	"fmt"
	"log"
	"path"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/manifoldco/promptui"
)

// Apply will load the context/configuration to be used by the apply command
func Apply(applyEnvironments []string, applyConfigPath string) *projectconfig.ZeroProjectConfig {
	if len(applyEnvironments) == 0 {
		fmt.Println(`Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.`)
		applyEnvironments = promptEnvironments()
	}

	validateEnvironments(applyEnvironments)

	if applyConfigPath == "" {
		exit.Fatal("config path cannot be empty!")
	}

	configPath := path.Join(projectconfig.RootDir, applyConfigPath)
	projectConfig := projectconfig.LoadConfig(configPath)
	return projectConfig
}

// promptEnvironments Prompts the user for the environments to apply against and returns a slice of strings representing the environments
func promptEnvironments() []string {
	items := map[string][]string{
		"Staging ":                    {"staging"},
		"Production":                  {"production"},
		"Both Staging and Production": {"staging", "production"},
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

func validateEnvironments(applyEnvironments []string) {
	// Strict for now, we can brainstorm how much we want to support custom environments later
	for _, env := range applyEnvironments {
		if env != "staging" && env != "production" {
			exit.Fatal("The currently supported environments are \"staging\" and \"production\"")
		}
	}
}
