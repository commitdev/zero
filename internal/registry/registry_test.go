package registry_test

import (
	"testing"

	"github.com/commitdev/zero/internal/registry"
	"github.com/stretchr/testify/assert"
)

func TestAvailableLabels(t *testing.T) {
	reg := testRegistry()

	t.Run("should be same order as declared", func(t *testing.T) {
		labels := registry.AvailableLabels(reg)
		assert.Equal(t, labels, []string{
			"EKS + Go + React + Gatsby",
			"foo",
			"bar",
			"lorem",
			"ipsum",
			"Custom",
		})
	})
}

func TestGetModulesByName(t *testing.T) {
	reg := testRegistry()
	t.Run("should return modules of specified stack", func(t *testing.T) {

		assert.Equal(t, registry.GetModulesByName(reg, "EKS + Go + React + Gatsby"),
			[]string{"module-source 1", "module-source 2"})
		assert.Equal(t, registry.GetModulesByName(reg, "lorem"), []string{"module-source 5"})
		assert.Equal(t, registry.GetModulesByName(reg, "ipsum"), []string{"module-source 6"})
		assert.Equal(t, registry.GetModulesByName(reg, "Custom"), []string{"module-source 7"})
	})
}

func testRegistry() registry.Registry {
	return registry.Registry{
		{"EKS + Go + React + Gatsby", []string{"module-source 1", "module-source 2"}},
		{"foo", []string{"module-source 3"}},
		{"bar", []string{"module-source 4"}},
		{"lorem", []string{"module-source 5"}},
		{"ipsum", []string{"module-source 6"}},
		{"Custom", []string{"module-source 7"}},
	}
}
