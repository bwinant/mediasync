package pkg_test

import (
	"testing"
	"reflect"
	"runtime"
	"fmt"
	"path/filepath"
)

func assertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d  expected %#v, actual %#v\n", filepath.Base(file), line, expected, actual)
		t.FailNow()
	}
}

func assertTrue(t *testing.T, condition bool) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d  expected condition to be true", filepath.Base(file), line)
		t.FailNow()
	}
}

func assertFalse(t *testing.T, condition bool) {
	if condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d  expected condition to be false", filepath.Base(file), line)
		t.FailNow()
	}
}