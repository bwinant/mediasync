package pkg

import (
	"fmt"
	"strings"
	"path/filepath"
	"flag"
	"os"
)

type stringList []string

func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(v string) error {
	*s = strings.Split(v, ",")
	return nil
}

func buildIgnoreList(parentPath string, paths []string) []string {
	for i, p := range paths {
		paths[i] = filepath.Join(parentPath, p)
	}
	return paths
}

type Options struct {
	Path string
	Ignore []string
}

func GetOptions(progName string) (Options, Options, string, bool, bool) {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", progName)
		flag.PrintDefaults()
	}

	var dir1, dir2 string
	flag.StringVar(&dir1, "dir1", "", "Directory 1 (Required)")
	flag.StringVar(&dir2, "dir2", "", "Directory 2 (Required)")

	var ignore1, ignore2 stringList
	flag.Var(&ignore1, "ignore1", "A comma-separated list of relative paths to ignore while scanning dir1")
	flag.Var(&ignore2, "ignore2", "A comma-separated list of relative paths to ignore while scanning dir2")

	var mode string
	flag.StringVar(&mode, "mode", "", "How to process directory diff. Values: [sync|import]")

	var verbose, dryRun bool
	flag.BoolVar(&dryRun, "dryRun", false,"Do not copy any files")
	flag.BoolVar(&verbose, "v", false,"Enable verbose output")

	flag.Parse()

	if dir1 == "" || dir2 == "" {
		flag.Usage()
		os.Exit(1)
	}

	if mode != "" && (mode != "sync" && mode != "import") {
		flag.Usage()
		os.Exit(1)
	}

	ignore1 = buildIgnoreList(dir1, ignore1)
	ignore2 = buildIgnoreList(dir2, ignore2)

	return Options{Path:dir1, Ignore: ignore1}, Options{Path:dir2, Ignore:ignore2}, mode, dryRun, verbose
}
