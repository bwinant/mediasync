package pkg

import (
	"fmt"
)

var verbose bool

func SetVerbose(v bool) {
	verbose = v
}

func debug(msg string) {
	if verbose {
		fmt.Println(msg)
	}
}

func debugf(format string, v ...interface{}) {
	if verbose {
		fmt.Printf(format + "\n", v...)
	}
}

func info(msg string) {
	fmt.Println(msg)
}

func infof(format string, v ...interface{}) {
	fmt.Printf(format + "\n", v...)
}

func errorf(format string, v ...interface{}) {
	fmt.Printf(format + "\n", v...)
}
