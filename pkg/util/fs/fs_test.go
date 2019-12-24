package fs

import (
	"testing"
)

var prependBaseDirTests = []struct {
	in     string
	prefix string
	out    string
}{
	{"../../dir/file.ext", "prefix", "../../prefix/dir/file.ext"},
	{"../opps/../../dir/file.ext", "prefix", "../../prefix/dir/file.ext"},
	{"../opps/../../dir/file.ext", "", "../../dir/file.ext"},
	{"dir/file.ext", "prefix", "prefix/dir/file.ext"},
	{"./dir/file.ext", "prefix", "prefix/dir/file.ext"},
}

func TestPrependBaseDir(t *testing.T) {
	for _, tt := range prependBaseDirTests {
		t.Run(tt.in, func(t *testing.T) {
			out := PrependBaseDir(tt.in, tt.prefix)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}
