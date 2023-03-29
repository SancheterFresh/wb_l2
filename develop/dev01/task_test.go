package main

import (
	"testing"
	"time"
)

func TestPrintTime(t *testing.T) {
	got := PrintTime()
	want := time.Since(got)
	if want > 50*time.Millisecond {
		t.Errorf("got %q, wanted %q", got, want)
	}
}