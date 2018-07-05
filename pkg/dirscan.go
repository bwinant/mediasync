package pkg

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Name         string
	AbsolutePath string
	Size         int64
	Metadata     map[string]string
}

func (f File) String() string {
	return f.Name
}


type DirScan struct {
	Name      string
	FileCount int
	IgnoredFileCount int
	DirCount  int
	IgnoredDirCount int
	Files     []*File
}

func (d DirScan) String() string {
	return d.Name
}


// Recursively scans a directory looking for files that meet criteria defined by the supplied Inspector
func Scan(startDir string, inspector Inspector) (*DirScan, error) {
	// Are we scanning a directory?
	stat, err := os.Stat(startDir)
	if err != nil {
		return nil, err
	}
	if !stat.Mode().IsDir() {
		return nil, errors.New(startDir + " is not a directory")
	}

	// Setup default if not specified
	if inspector == nil {
		inspector = DefaultInspector{}
	}

	// Track results of directory scan
	d := DirScan{Name: startDir}

	// Closure to walk directory
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Do we want to traverse this directory?
		if info.IsDir() {
			err := inspector.acceptDir(path, info)

			if err != nil {
				infof("Ignoring directory %s", path)
				d.IgnoredDirCount++
			} else if path != startDir { // Don't count the initial directory
				d.DirCount++
			}

			return err
		}

		if inspector.acceptFile(path, info) {
			// Found a file to process
			d.FileCount++

			if d.FileCount % 500 == 0 {
				debugf("Found %d files", d.FileCount)
			}

			// Will need file path relative to the starting directory
			relPath, err := filepath.Rel(startDir, path)
			if err != nil {
				infof("Unable to calculate relative path for file %s", path)
			}

			// Perform optional further processing of file
			file := inspector.process(&File{Name: relPath, AbsolutePath: path, Size: info.Size(), Metadata: make(map[string]string)})

			// Update directory scan results
			d.Files = append(d.Files, file)
		} else {
			debugf("Ignoring file %s", path)
			d.IgnoredFileCount++
		}

		return nil
	}

	infof("Scanning %s", startDir)

	start := time.Now()
	err = filepath.Walk(startDir, walkFn)
	elapsed := time.Since(start)

	infof("Found %d files in %s (ignored %d directories, %d files)", d.FileCount, elapsed, d.IgnoredDirCount, d.IgnoredFileCount)

	return &d, err
}