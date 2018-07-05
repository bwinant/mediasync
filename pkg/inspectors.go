package pkg

import (
	"os"
	"strings"
	"path/filepath"
)

// Defines what files to scan
type Inspector interface {
	// Returns filepath.SkipDir if directory should be ignored during scan
	acceptDir(path string, info os.FileInfo) error

	// Returns true if file should be processed during scan
	acceptFile(path string, info os.FileInfo) bool

	// Performs any optional scan processing of the file
	process(file *File) *File
}

// Default implementation of Inspector that does nothing
type DefaultInspector struct {
}

// Always returns nil
func (i DefaultInspector) acceptDir(path string, info os.FileInfo) error {
	return nil
}

// Always returns true
func (i DefaultInspector) acceptFile(path string, info os.FileInfo) bool {
	return true
}

// A no-op
func (i DefaultInspector) process(file *File) *File {
	return file
}


type ConfigurableInspector struct {
	ignoreHidden bool
	ignorePaths *Set
	extensions *Set
}

func ConfigureInspector (ignoreHidden bool, ignorePaths []string, extensions []string) *ConfigurableInspector {
	ins := &ConfigurableInspector{ignoreHidden: ignoreHidden, ignorePaths:NewSet(), extensions:NewSet()}
	if ignorePaths != nil {
		ins.ignorePaths.AddAll(ignorePaths)
	}
	if extensions != nil {
		for _, ext := range extensions {
			ins.extensions.Add(strings.ToLower(ext))
		}
	}
	return ins
}

func (i ConfigurableInspector) acceptDir(path string, info os.FileInfo) error {
	if info.IsDir() {
		if i.ignoreHidden && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if i.ignorePaths.Contains(path) {
			return filepath.SkipDir
		}
	}

	return nil
}

func (i ConfigurableInspector) acceptFile(path string, info os.FileInfo) bool {
	if i.ignoreHidden &&  strings.HasPrefix(info.Name(), ".") {
		return false
	}

	ext := filepath.Ext(info.Name())
	if ext != "" {
		ext = strings.ToLower(ext)

		if i.extensions.Contains(ext) {
			return true
		}
	}

	return false
}

func (i ConfigurableInspector) process(file *File) *File {
	return file
}
