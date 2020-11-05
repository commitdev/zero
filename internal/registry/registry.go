package registry

type Registry []Stack
type Stack struct {
	Name          string
	ModuleSources []string
}

func GetRegistry(path string) Registry {
	return Registry{
		// TODO: better place to store these options as configuration file or any source
		{
			"EKS + Go + React + Gatsby",
			[]string{
				path + "/zero-aws-eks-stack",
				path + "/zero-deployable-landing-page",
				path + "/zero-deployable-backend",
				path + "/zero-deployable-react-frontend",
			},
		},
		{
			"EKS + NodeJS + React + Gatsby",
			[]string{
				path + "/zero-aws-eks-stack",
				path + "/zero-deployable-landing-page",
				path + "/zero-deployable-node-backend",
				path + "/zero-deployable-react-frontend",
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
