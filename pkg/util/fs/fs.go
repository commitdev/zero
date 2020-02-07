package fs

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

// CreateDirs creates directories from the given directory path arguments.
func CreateDirs(dirPaths ...string) error {
	for _, path := range dirPaths {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	return nil
}

// FileExists checks whether the given path exists and belongs to a file.
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%v: is a directory, expected file", path)
	}

	return true, nil
}

// PrependPath prepends a path with prefix while disregarding back directories ../
func ReplacePath(p, old, new string) string {
	return path.Clean(strings.Replace(p, old, new, 1))
}

// PrependPath prepends a path with prefix while disregarding back directories ../
func PrependPath(filepath string, prefix string) string {
	re := regexp.MustCompile(`(\.\.\/)+`)
	cleanPath := path.Clean(filepath)
	baseDir := re.FindString(cleanPath)
	if baseDir == "" {
		return path.Join(prefix, cleanPath)
	}
	return strings.Replace(cleanPath, baseDir, path.Join(baseDir, prefix)+"/", 1)
}
