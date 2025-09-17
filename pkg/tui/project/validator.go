package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func validatePath(path string) error {
	if path == "" {
		return nil
	}

	//Check if parent path exists
	parentPath := filepath.Dir(path)
	_, err := os.Stat(parentPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("provided path does not exist")
		}
		return fmt.Errorf("invalid path: %e", err)
	}

	//Check if the provided path is a directory
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return fmt.Errorf("`%s` does not point to a directory", path)
	}
	return nil
}
