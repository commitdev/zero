package module

import (
	"testing"
)

func TestIsLocal(t *testing.T) {
	source := "./tests/test_data/modules"
	res := IsLocal(source)
	if !res {
		t.Errorf("Error, source %s SHOULD BE determined as local", source)
	}

	source = "https://github.com/commitdev/my-repo"
	res = IsLocal(source)
	if res {
		t.Errorf("Error, source %s SHOULD NOT BE determined as local", source)
	}
}
