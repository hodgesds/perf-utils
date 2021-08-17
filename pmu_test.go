// +build linux

package perf

import (
	"testing"
)

func TestAvailablePMUs(t *testing.T) {
	pmus, err := AvailablePMUs()
	if err != nil {
		t.Fatal(err)
	}
	if len(pmus) == 0 {
		t.Fatal("no PMU events")
	}
}
