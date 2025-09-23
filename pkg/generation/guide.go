package gen

import (
	"os"
	"path/filepath"
)

const DefaultDirPerms = 0775
const DefaultFilePerms = 0664

type Guide struct {
	RootPath  string
	DirPerms  os.FileMode
	FilePerms os.FileMode
}

func NewGuide(rootPath string, dirPerms, filePerms os.FileMode) *Guide {
	rootPath = filepath.Clean(rootPath)
	rootPath, err := filepath.Localize(rootPath)
	if err != nil {
		rootPath = filepath.Clean(rootPath)
	}
	return &Guide{
		RootPath:  rootPath,
		DirPerms:  dirPerms,
		FilePerms: filePerms,
	}
}

func DefaultGuide(rootPath string) *Guide {
	return NewGuide(rootPath, DefaultDirPerms, DefaultFilePerms)
}
