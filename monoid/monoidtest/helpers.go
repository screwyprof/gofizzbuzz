package monoidtest

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, want, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func StrPtr(s string) *string {
	return &s
}
