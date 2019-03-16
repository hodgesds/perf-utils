package perf

import (
	"os"
	"testing"
)

func TestCacheProfiler(t *testing.T) {
	p := NewCacheProfiler(os.Getpid(), 0)
	defer func() {
		if err := p.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := p.Reset(); err != nil {
		t.Fatal(err)
	}
	if err := p.Start(); err != nil {
		t.Fatal(err)
	}
	_, err := p.Profile()
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Stop(); err != nil {
		t.Fatal(err)
	}
}
