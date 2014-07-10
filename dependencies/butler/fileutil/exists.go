package fileutil

import (
	"path/filepath"
	"os"
	"strings"
)

// `Exists` returns whether the given file or directory exists or not.
// (Grabbed off of Stack Overflow)
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// `StrictExists` returns whether the given path exists at prefix,
// where the value of path is case-sensitive.
func CaseSensitiveExists(path, prefix string) (bool, error) {
	path = filepath.FromSlash(path)
	// pathDepth = number of separators + 1
	pathDepth := strings.Count(filepath.Clean(path), string(filepath.Separator)) + 1
	// glob = prefix + path.replace(filenames, "*")
	glob := filepath.Clean(prefix + strings.Repeat(string(filepath.Separator)+"*", pathDepth))
	matches, err := filepath.Glob(glob)
	if err != nil {
		return false, err
	}
	expectedFilename := filepath.Join(prefix, path)
	for _, match := range matches {
		if match == expectedFilename {
			return true, nil
		}
	}
	return false, nil
}

func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true, nil
	}
	return false, err
}
