package fs

import (
	"testing"
)

var replacePathTest = []struct {
	path string
	old  string
	new  string
	out  string
}{
	{"../../dir/file.ext", "../../dir", "output", "output/file.ext"},
	{"dir/file.ext", "dir", "output", "output/file.ext"},
}

func TestReplacePath(t *testing.T) {
	for _, tt := range replacePathTest {
		t.Run(tt.path, func(t *testing.T) {
			out := ReplacePath(tt.path, tt.old, tt.new)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}

var prependPathTests = []struct {
	in     string
	prefix string
	out    string
}{
	{"../../dir/file.ext", "prefix", "../../prefix/dir/file.ext"},
	{"../opps/../../dir/file.ext", "prefix", "../../prefix/dir/file.ext"},
	{"../opps/../../dir/file.ext", "", "../../dir/file.ext"},
	{"dir/file.ext", "prefix", "prefix/dir/file.ext"},
	{"dir/file.ext", "../prefix", "../prefix/dir/file.ext"},
	{"./dir/file.ext", "prefix", "prefix/dir/file.ext"},
}

func TestPrependPath(t *testing.T) {
	for _, tt := range prependPathTests {
		t.Run(tt.in, func(t *testing.T) {
			out := PrependPath(tt.in, tt.prefix)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}
