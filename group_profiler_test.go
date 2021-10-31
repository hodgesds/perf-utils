package perf

import (
	"os"
	"testing"
)

func TestGroupProfiler(t *testing.T) {
	p, err := NewGroupProfiler(
		os.Getpid(),
		-1,
		0,
		CPUMigrationsEventAttr(),
		MinorPageFaultsEventAttr(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Reset(); err != nil {
		t.Fatal(err)
	}

	if err := p.Start(); err != nil {
		t.Fatal(err)
	}

	err = p.Profile(&GroupProfileValue{})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Stop(); err != nil {
		t.Fatal(err)
	}

	if err := p.Close(); err != nil {
		t.Fatal(err)
	}
}
