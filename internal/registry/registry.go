package registry

type Registry []Stack
type Stack struct {
	Name          string
	ModuleSources []string
}

func GetRegistry() Registry {
	return Registry{
		// TODO: better place to store these options as configuration file or any source
		{
			"EKS + Go + React + Gatsby",
			[]string{
				"github.com/commitdev/zero-aws-eks-stack",
				"github.com/commitdev/zero-deployable-landing-page",
				"github.com/commitdev/zero-deployable-backend",
				"github.com/commitdev/zero-deployable-react-frontend",
			},
		},
		{
			"Custom", []string{},
		},
	}
}

func GetModulesByName(registry Registry, name string) []string {
	for _, v := range registry {
		if v.Name == name {
			return v.ModuleSources
		}
	}
	return []string{}
}

func AvailableLabels(registry Registry) []string {
	labels := make([]string, len(registry))
	i := 0
	for _, stack := range registry {
		labels[i] = stack.Name
		i++
	}
	return labels
}
