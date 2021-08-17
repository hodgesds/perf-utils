// +build linux

package perf

import (
	"testing"
)

func TestTraceFSMount(t *testing.T) {
	mount, err := TraceFSMount()
	if err != nil {
		t.Fatal(err)
	}
	if len(mount) == 0 {
		t.Fatal("tracefs not mounted")
	}
}

func TestDebugFSMount(t *testing.T) {
	mount, err := DebugFSMount()
	if err != nil {
		t.Fatal(err)
	}
	if len(mount) == 0 {
		t.Fatal("debugfs not mounted")
	}
}
