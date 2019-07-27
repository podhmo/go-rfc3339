package rfc3339_test

import (
	"testing"

	rfc3339 "github.com/podhmo/go-rfc3339"
)

func TestIt(t *testing.T) {
	now := rfc3339.MustParse("2000-01-01T00:00:00Z")

	want := "2000-01-01T00:00:00Z"
	got := rfc3339.Format(now)
	if got != want {
		t.Errorf("want %q, but %q", want, got)
	}
}
