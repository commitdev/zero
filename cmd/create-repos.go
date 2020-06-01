package cmd

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"os"
	"os/exec"
)

// this is being developed with the following assumptions:
// 1. the initializeRepositories function and its deps will be moved into cmd/create.go eventually
// 2. i didn't implement cobra.Command.  assuming this will get used by the Create command eventually
// 3. create.go will handle parsing the repo names and github credentials and pass them into initializeRepostiories().
// 4. configs are just hardcoded here for testing purposes
// 5. if organizationName is set, create an org owned repo.  if not, create a personal repo.

// takes a list of directories containing modules to create a repos and and do initial commit for
// can be converted to a private function if moved into create.go
func InitializeRepositories(moduleDirs []string, remoteRepository string, organizationName string, githubApiKey string) {

	for _, moduleDir := range moduleDirs {
		// if organizationName is not set, create a personal github repo,  else create an org owned repo.
		if organizationName == "" {
			if err := createPersonalRepository(moduleDir, githubApiKey); err != nil {
				fmt.Printf("error creating repository: %s\n", err.Error())
				continue
			}
		} else {
			if err := createOrganizationOwnedRepository(moduleDir, githubApiKey, organizationName); err != nil {
				fmt.Printf("error creating repository: %s\n", err.Error())
				continue
			}
		}

		if err := doInitialCommit(moduleDir, remoteRepository); err != nil {
			fmt.Printf("error initializing repository: %s\n", err.Error())
		}
	}

}

const createPersonalRepositoryMutation = `mutation ($repoName: String!, $repoDescription: String!) {
		createRepository(
			input: {
				name:$repoName, 
				visibility: PRIVATE, 
				description: $repoDescription
			}) 
		{
			clientMutationId
		}
	}`

func createPersonalRepository(moduleDir string, githubApiKey string) error {

	fmt.Printf("Creating repository for module: %s\n", moduleDir)

	// create client and mutation
	client := graphql.NewClient("https://api.github.com/graphql")
	req := graphql.NewRequest(createPersonalRepositoryMutation)
	req.Var("repoName", moduleDir)
	req.Var("repoDescription", fmt.Sprintf("Repository for %s", moduleDir))

	// add auth token
	var bearer = fmt.Sprintf("Bearer %s", githubApiKey)
	req.Header.Add("Authorization", bearer)

	ctx := context.Background()
	if err := client.Run(ctx, req, nil); err != nil {
		return err
	}

	fmt.Printf("Repository successfully created for module: %s\n", moduleDir)

	return nil
}

const createOrganizationRepositoryMutation = `mutation ($repoName: String!, $repoDescription: String!, $ownerId: String!) {
		createRepository(
			input: {
				name:$repoName, 
				visibility: PRIVATE, 
				description: $repoDescription
				ownerId: $ownerId
			}) 
		{
			clientMutationId
		}
	}`

const getOrganizationQuery = `query ($organizationName: String!) {
		organization(login: $organizationName) {
			id
		}
	}`

type organizationQueryResponse struct {
	Organization struct {
		Id string
	}
}

func createOrganizationOwnedRepository(moduleDir string, githubApiKey string, organizationName string) error {

	fmt.Printf("Creating org owned repository for module: %s\n", moduleDir)

	// create client and organization query
	var bearer = fmt.Sprintf("Bearer %s", githubApiKey)
	client := graphql.NewClient("https://api.github.com/graphql")
	orgIdReq := graphql.NewRequest(getOrganizationQuery)
	orgIdReq.Var("organizationName", organizationName)
	orgIdReq.Header.Add("Authorization", bearer)

	var orgIdResp organizationQueryResponse
	ctx := context.Background()
	if err := client.Run(ctx, orgIdReq, &orgIdResp); err != nil {
		return err
	}
	organizationId := orgIdResp.Organization.Id

	// create mutation and run it
	req := graphql.NewRequest(createOrganizationRepositoryMutation)
	req.Var("repoName", moduleDir)
	req.Var("repoDescription", fmt.Sprintf("Repository for %s", moduleDir))
	req.Var("ownerId", organizationId)
	req.Header.Add("Authorization", bearer)
	if err := client.Run(ctx, req, nil); err != nil {
		return err
	}

	fmt.Printf("Repository successfully created for module: %s\n", moduleDir)

	return nil
}

type InitialCommands struct {
	description string
	command     string
	args        []string
}

// do initial commit to a repository
func doInitialCommit(moduleDir string, remoteRepository string) error {
	fmt.Printf("Initializing repository for module: %s\n", moduleDir)

	remoteOrigin := fmt.Sprintf("%s/%s.git", remoteRepository, moduleDir)
	commands := []InitialCommands{
		{
			description: "git init",
			command:     "git",
			args:        []string{"init"},
		},
		{
			description: "git add .",
			command:     "git",
			args:        []string{"add", "."},
		},
		{
			description: "git commit -m \"initial commit by zero\"",
			command:     "git",
			args:        []string{"commit", "-m", "initial commit by zero"},
		},
		{
			description: fmt.Sprintf("git remote add origin %s", remoteOrigin),
			command:     "git",
			args:        []string{"remote", "add", "origin", remoteOrigin},
		},
		{
			description: "git push -u origin master",
			command:     "git",
			args:        []string{"push", "-u", "origin", "master"},
		},
	}

	// create README.md. may not need this.
	if err := createReadme(moduleDir); err != nil {
		return err
	}

	for _, command := range commands {
		fmt.Printf(">> %s\n", command.description)

		cmd := exec.Command(command.command, command.args...)
		cmd.Dir = "./" + moduleDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("ERROR: failed to run %s: %s\n", command.description, err.Error())
			// this is a partial failure.  some commands may have exec'ed successfully.
			break
		} else {
			response := string(out)
			if len(response) > 0 {
				fmt.Println(response)
			}
		}
	}

	fmt.Printf("Repository successfully initialized for module: %s\n", moduleDir)

	return nil
}

func createReadme(moduleDir string) error {
	fmt.Println(">> create readme file")

	readmeFilename := fmt.Sprintf("./%s/README.md", moduleDir)
	readmeFile, err := os.Create(readmeFilename)
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	readmeContent := "# repository created using zero"
	cmd := exec.Command("echo", readmeContent)
	cmd.Stdout = readmeFile
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
