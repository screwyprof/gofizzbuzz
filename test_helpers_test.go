package gofizzbuzz_test

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	stdOut := os.Stdout
	stdErr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)

	return func() {
		defer null.Close()

		os.Stdout = stdOut
		os.Stderr = stdErr

		log.SetOutput(os.Stderr)
	}
}

func assertEquals(t *testing.T, want, got string) {
	t.Helper()

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func funcName(i interface{}) string {
	funcPkgPath := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	nameEnd := filepath.Ext(funcPkgPath) // .foo

	return strings.TrimPrefix(nameEnd, ".") // foo
}
