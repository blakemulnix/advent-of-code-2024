package utils

import (
	"runtime"
	"path/filepath"
)

// GetInputPath returns the path to input.txt relative to the caller's file location
func GetInputPath(skip int) string {
	_, filename, _, ok := runtime.Caller(skip)
	if !ok {
		panic("Failed to get current file path")
	}
	return filepath.Join(filepath.Dir(filename), "input.txt")
} 