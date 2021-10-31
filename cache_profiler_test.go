package perf

import (
	"os"
	"testing"
)

func TestCacheProfiler(t *testing.T) {
	p, err := NewCacheProfiler(os.Getpid(), 0, AllCacheProfilers)
	if err != nil {
		t.Fatal(err)
	}
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
	err = p.Profile(&CacheProfile{})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Stop(); err != nil {
		t.Fatal(err)
	}
}
