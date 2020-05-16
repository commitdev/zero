package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

type requirement struct {
	name    string
	command string
	args    []string
	docsURL string
	checker func([]byte) (string, error)
}

type versionError struct {
	errorText string
}

type commandError struct {
	Command    string
	ErrorText  string
	Suggestion string
}

func (e *versionError) Error() string {
	return fmt.Sprintf("%s", e.errorText)
}

func (e *commandError) Error() string {
	return fmt.Sprintf("%s", e.ErrorText)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print the check number of commit0",
	Run: func(cmd *cobra.Command, args []string) {
		// Add any new requirements to this slice.
		required := []requirement{
			{
				name:    "AWS CLI\t\t",
				command: "aws",
				args:    []string{"--version"},
				docsURL: "",
				checker: func(output []byte) (string, error) {
					ver := ""
					re := regexp.MustCompile(`aws-cli/([0-9]+)\.([0-9]+)\.([0-9]+)`)
					m := re.FindStringSubmatch(string(output))
					major, err := strconv.ParseInt(m[1], 0, 64)
					if err != nil {
						return ver, err
					}
					minor, err := strconv.ParseInt(m[2], 0, 64)
					if err != nil {
						return ver, err
					}
					patch, err := strconv.ParseInt(m[3], 0, 64)
					if err != nil {
						return ver, err
					}

					ver = fmt.Sprintf("%d.%d.%d", major, minor, patch)

					if major < 1 || (major == 1 && minor < 16) {
						return ver, &versionError{"Requires 1.16 or greater."}
					}

					return ver, err
				}},
			{
				name:    "Kubectl\t\t",
				command: "kubectl",
				args:    []string{"version", "--client=true"},
				docsURL: "https://kubernetes.io/docs/tasks/tools/install-kubectl/",
				checker: func(output []byte) (string, error) {
					ver := ""
					re := regexp.MustCompile(`version\.Info{Major:"([0-9]+)", Minor:"([0-9]+)"`)
					m := re.FindStringSubmatch(string(output))
					major, err := strconv.ParseInt(m[1], 0, 64)
					if err != nil {
						return ver, err
					}
					minor, err := strconv.ParseInt(m[2], 0, 64)
					if err != nil {
						return ver, err
					}

					ver = fmt.Sprintf("%d.%d", major, minor)

					if major < 1 || (major == 1 && minor < 12) {
						return ver, &versionError{"Requires 2.12 or greater."}
					}

					return ver, err
				},
			},
			{
				name:    "Terraform\t",
				command: "terraform",
				args:    []string{"version"},
				docsURL: "https://www.terraform.io/downloads.html",
				checker: func(output []byte) (string, error) {
					ver := ""
					re := regexp.MustCompile(`Terraform v([0-9]+)\.([0-9]+)\.([0-9]+)`)
					m := re.FindStringSubmatch(string(output))
					major, err := strconv.ParseInt(m[1], 0, 64)
					if err != nil {
						return ver, err
					}
					minor, err := strconv.ParseInt(m[2], 0, 64)
					if err != nil {
						return ver, err
					}
					patch, err := strconv.ParseInt(m[3], 0, 64)
					if err != nil {
						return ver, err
					}

					ver = fmt.Sprintf("%d.%d.%d", major, minor, patch)

					if major < 0 || (major == 0 && minor < 12) {
						return ver, &versionError{"Zero requires terraform 0.12 or greater."}
					}

					return ver, err
				},
			},
			{
				name:    "jq\t\t",
				command: "jq",
				args:    []string{"--version"},
				docsURL: "https://stedolan.github.io/jq/download/",
				checker: func(output []byte) (string, error) {
					ver := ""
					re := regexp.MustCompile(`jq-([0-9]+)\.([0-9]+)-`)
					m := re.FindStringSubmatch(string(output))
					major, err := strconv.ParseInt(m[1], 0, 64)
					if err != nil {
						return ver, err
					}
					minor, err := strconv.ParseInt(m[2], 0, 64)
					if err != nil {
						return ver, err
					}

					ver = fmt.Sprintf("%d.%d", major, minor)

					if major < 1 || (major == 1 && minor < 5) {
						return ver, &versionError{"Requires jq version 1.15 or greater."}
					}

					return ver, err
				}},
			{
				name:    "Git\t\t",
				command: "git",
				args:    []string{"version"},
				docsURL: "https://git-scm.com/book/en/v2/Getting-Started-Installing-Git",
				checker: func(output []byte) (string, error) {
					ver := ""
					re := regexp.MustCompile(`git version ([0-9]+)\.([0-9]+)\.([0-9]+)`)
					m := re.FindStringSubmatch(string(output))
					major, err := strconv.ParseInt(m[1], 0, 64)
					if err != nil {
						return ver, err
					}
					minor, err := strconv.ParseInt(m[2], 0, 64)
					if err != nil {
						return ver, err
					}
					patch, err := strconv.ParseInt(m[3], 0, 64)
					if err != nil {
						return ver, err
					}

					ver = fmt.Sprintf("%d.%d.%d", major, minor, patch)

					if major < 2 || (major == 2 && minor < 12) {
						return ver, &versionError{"Zero requires git version 2.12 or greater."}
					}

					return ver, err
				}},
		}

		// Store and errors from the commands we run.
		errors := []commandError{}

		fmt.Println("Checking Zero Requirements...")
		for _, r := range required {
			fmt.Printf("%s", r.name)
			out, err := exec.Command(r.command, r.args...).CombinedOutput()
			if err != nil {
				cerr := commandError{
					fmt.Sprintf("%s %s", r.command, strings.Join(r.args, " ")),
					err.Error(),
					r.docsURL,
				}
				errors = append(errors, cerr)
				fmt.Printf("\033[0;31mFAIL\033[0m\t\t%s\n", "-")
				continue
			}
			version, err := r.checker(out)
			if err != nil {
				cerr := commandError{
					r.command,
					err.Error(),
					r.docsURL,
				}
				errors = append(errors, cerr)
				fmt.Printf("\033[0;31mFAIL\033[0m\t\t%s\n", version)
			} else {
				fmt.Printf("\033[0;32mPASS\033[0m\t\t%s\n", version)
			}
		}

		if len(errors) > 0 {
			// initialize tabwriter
			w := new(tabwriter.Writer)

			// minwidth, tabwidth, padding, padchar, flags
			w.Init(os.Stdout, 10, 12, 2, ' ', 0)

			defer w.Flush()

			fmt.Fprintf(w, "\n %s\t%s\t%s\t", "Command", "Error", "Info")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t", "---------", "---------", "---------")

			for _, e := range errors {
				fmt.Fprintf(w, "\n%s\t%s\t%s\t", e.Command, e.ErrorText, e.Suggestion)
			}
		}
		fmt.Println()
	},
}
