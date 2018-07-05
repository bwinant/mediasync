package main

import (
	ms "mediasync/pkg"
	"fmt"
	"os"
)

func main() {
	extensions := []string{".mp3", ".m4a", ".m4p", ".m4v", ".aif"}

	options1, options2, mode, dryRun, verbose := ms.GetOptions("itunesSync")
	ms.SetDryRun(dryRun)
	ms.SetVerbose(verbose)

	inspector1 := ms.ConfigureInspector(true, options1.Ignore, extensions)
	scan1, err := ms.Scan(options1.Path, inspector1)
	if err != nil {
		fatal(err)
	}

	inspector2 := ms.ConfigureInspector(true, options2.Ignore, extensions)
	scan2, err := ms.Scan(options2.Path, inspector2)
	if err != nil {
		fatal(err)
	}

	diff1 := ms.Compare(scan1, scan2)
	diff2 := ms.Compare(scan2, scan1)

	switch mode {
		case "sync":
			err = ms.Sync(options1.Path, diff1, options2.Path, diff2)

		case "import":
			err = ms.CreateImport(diff2, "import")
	}

	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}