package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {

	get, _ := UnpackString("f4h3g5t3")
	want := "ffffhhhgggggttt"
	if get != want {
		t.Errorf("Output %q not equal to expected %q", get, want)
	}

}
