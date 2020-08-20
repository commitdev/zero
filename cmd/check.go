package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/tabwriter"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

type requirement struct {
	name       string
	command    string
	args       []string
	minVersion string
	regexStr   string
	docsURL    string
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

func printErrors(errors []commandError) {
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

// getSemver uses the regular expression from the requirement to parse the
// output of a command and extract the version from it. Returns the version
// or an error if the version string could not be parsed.
func getSemver(req requirement, out []byte) (*semver.Version, error) {
	re := regexp.MustCompile(req.regexStr)
	v := re.FindStringSubmatch(string(out))
	if len(v) < 4 {
		return nil, &commandError{
			req.command,
			"Could not find version number in output",
			fmt.Sprintf("Try running %s %s locally and checking it works.", req.command, strings.Join(req.args, " ")),
		}
	}

	// Default patch version number to 0 if it doesn't exist
	if v[3] == "" {
		v[3] = "0"
	}

	versionString := fmt.Sprintf("%s.%s.%s", v[1], v[2], v[3])
	version, err := semver.NewVersion(versionString)
	if err != nil {
		return version, err
	}
	return version, nil
}

// checkSemver validates that the version of a tool meets the minimum required
// version listed in your requirement. Returns a boolean.
// For more information on parsing semver, see semver.org
// If your tool doesn't do full semver then you may need to add custom logic
// to support it.
func checkSemver(req requirement, actualVersion *semver.Version) bool {
	requiredVersion := semver.New(req.minVersion)
	return actualVersion.LessThan(*requiredVersion)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print the check number of zero",
	Run: func(cmd *cobra.Command, args []string) {
		// Add any new requirements to this slice.
		required := []requirement{
			{
				name:       "AWS CLI\t\t",
				command:    "aws",
				args:       []string{"--version"},
				regexStr:   `aws-cli\/(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`,
				minVersion: "1.16.0",
				docsURL:    "https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html",
			},
			{
				name:       "Kubectl\t\t",
				command:    "kubectl",
				args:       []string{"version", "--client=true", "--short"},
				regexStr:   `Client Version: v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`,
				minVersion: "1.15.2",
				docsURL:    "https://kubernetes.io/docs/tasks/tools/install-kubectl/",
			},
			{
				name:       "Terraform\t",
				command:    "terraform",
				args:       []string{"version"},
				regexStr:   `Terraform v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`,
				minVersion: "0.13.0",
				docsURL:    "https://www.terraform.io/downloads.html",
			},
			{
				name:       "jq\t\t",
				command:    "jq",
				args:       []string{"--version"},
				regexStr:   `jq-(0|[1-9]\d*)\.(0|[1-9]\d*)\-?(0|[1-9]\d*)?`,
				minVersion: "1.5.0",
				docsURL:    "https://stedolan.github.io/jq/download/",
			},
			{
				name:       "Git\t\t",
				command:    "git",
				args:       []string{"version"},
				regexStr:   `^git version (0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`,
				minVersion: "2.17.1",
				docsURL:    "https://git-scm.com/book/en/v2/Getting-Started-Installing-Git",
			},
			{
				name:       "Wget\t\t",
				command:    "wget",
				args:       []string{"--version"},
				regexStr:   `^GNU Wget (0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`,
				minVersion: "1.14.0",
				docsURL:    "https://www.gnu.org/software/wget/",
			},
		}

		// Store and errors from the commands we run.
		errors := []commandError{}

		fmt.Println("Checking Zero Requirements...")
		for _, r := range required {
			fmt.Printf("%s", r.name)
			// In future we could parse the stderr and stdout separately, but for now it's nice to see
			// the full output on a failure.
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
			version, err := getSemver(r, out)
			if err != nil {
				cerr := commandError{
					r.command,
					err.Error(),
					r.docsURL,
				}
				errors = append(errors, cerr)
				fmt.Printf("\033[0;31mFAIL\033[0m\t\t%s\n", version)
				continue
			}
			if checkSemver(r, version) {
				cerr := commandError{
					r.command,
					fmt.Sprintf("Version does not meet required. Want: %s; Got: %s", r.minVersion, version),
					r.docsURL,
				}
				errors = append(errors, cerr)
				fmt.Printf("\033[0;31mFAIL\033[0m\t\t%s\n", version)
			} else {
				fmt.Printf("\033[0;32mPASS\033[0m\t\t%s\n", version)
			}
		}

		if len(errors) > 0 {
			printErrors(errors)
			os.Exit(1)
		}

		fmt.Println()
	},
}
