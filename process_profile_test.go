package perf

import (
	"os"
	"testing"

	"golang.org/x/sys/unix"
)

func TestProfiler(t *testing.T) {
	profiler, err := NewProfiler(
		unix.PERF_TYPE_HARDWARE,
		unix.PERF_COUNT_HW_INSTRUCTIONS,
		os.Getpid(),
		-1,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := profiler.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := profiler.Start(); err != nil {
		t.Fatal(err)
	}

	_, err = profiler.Profile()
	if err != nil {
		t.Fatal(err)
	}

	if err := profiler.Stop(); err != nil {
		t.Fatal(err)
	}
}
