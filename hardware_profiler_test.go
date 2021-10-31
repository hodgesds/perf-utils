package perf

import (
	"encoding/json"
	"os"
	"testing"
)

func TestHardwareProfiler(t *testing.T) {
	hwProfiler, err := NewHardwareProfiler(os.Getpid(), -1, AllHardwareProfilers)
	if err != nil && !hwProfiler.HasProfilers() {
		t.Fatal(err)
	}
	defer func() {
		if err := hwProfiler.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := hwProfiler.Start(); err != nil {
		t.Fatal(err)
	}

	profile := &HardwareProfile{}
	err = hwProfiler.Profile(profile)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(profile)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatalf("Expected data, got: %+v\n", data)
	}

	if err := hwProfiler.Stop(); err != nil {
		t.Fatal(err)
	}
}
