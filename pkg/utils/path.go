package utils

import "strings"

// TrimPath removes leading and trailing slashes from a file path.
func TrimPath(path string) string {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	return path
}
