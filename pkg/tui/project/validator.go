package project

import (
	"errors"
	"fmt"
	"os"

	"github.com/Ahu-Tools/AhuM/pkg/project"
)

func validateDbs(dbs []project.Database) error {
	if len(dbs) == 0 {
		return errors.New("You have to choose at least one database!")
	}

	return nil
}

func validateEdges(edges []project.Edge) error {
	if len(edges) == 0 {
		return errors.New("You have to choose at least one edge!")
	}

	return nil
}

func validatePath(path string) error {
	if path == "" {
		return nil
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("provided path does not exist")
		}
		return fmt.Errorf("invalid path: %e", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("`%s` does not point to a directory", path)
	}
	return nil
}

func validateHost(s string) error {
	if s == "" {
		return errors.New("Host should not be empty!")
	}
	return nil
}

func validateName(s string) error {
	if s == "" {
		return errors.New("Name should not be empty!")
	}
	return nil
}
