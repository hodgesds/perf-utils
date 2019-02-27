package perf

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSoftwareProfiler(t *testing.T) {
	swProfiler := NewSoftwareProfiler(os.Getpid(), -1)
	defer func() {
		if err := swProfiler.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := swProfiler.Start(); err != nil {
		t.Fatal(err)
	}

	profile, err := swProfiler.Profile()
	if err != nil {
		t.Fatal(err)
	}
	_, err = json.Marshal(profile)
	if err != nil {
		t.Fatal(err)
	}

	if err := swProfiler.Stop(); err != nil {
		t.Fatal(err)
	}
}
