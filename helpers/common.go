package helpers

import (
	"os"
	"path/filepath"
)

// IntToInt8Pointer converts integer to pinter of int8
func IntToInt8Pointer(i int) *int8 {
	i8 := int8(i)
	return &i8
}

// MakeFilePathAbsolute checks if provided file path is absolute
// or not. If provided path is absolute, it will be returned, if
// it is not, the absolute path will be calculated and returned.
func MakeFilePathAbsolute(path string) (string, error) {
	var absoluteFilePath = path
	if !filepath.IsAbs(path) {
		workingDirectory, err := os.Getwd()
		if err != nil {
			return "", err
		}
		absoluteFilePath = filepath.Join(workingDirectory, path)
	}

	return absoluteFilePath, nil
}
