package pkg

import (
	"path/filepath"
	"os"
	"io"
)

var dryRun bool
func SetDryRun(v bool) {
	dryRun = v
}

// Copy files1 to path2 and files2 to path1
func Sync(path1 string, files1 []*File, path2 string, files2 []*File) error {
	if err := copyFiles(files2, path1); err != nil {
		return err
	}

	if err := copyFiles(files1, path2); err != nil {
		return err
	}

	return nil
}

func CreateImport(files []*File, destDir string) error {
	return copyFiles(files, destDir)
}

func copyFiles(files []*File, destDir string) error {
	n := len(files)

	if n > 0 {
		infof("Copying %d files to %s", n, destDir)

		for _, f := range files {
			dest := filepath.Join(destDir, f.Name)

			infof("Copying %s to %s", f.AbsolutePath, dest)
			if !dryRun {
				err := copyFile(f.AbsolutePath, dest)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func copyFile(src string, dest string) error {
	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, os.ModePerm)
	}

	s, err := os.Open(src)
	if err != nil {
		return nil
	}
	defer s.Close()

	d, err := os.Create(dest)

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	return d.Close()
}