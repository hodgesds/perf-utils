package perf

import (
	"testing"
)

func TestCPUInstructions(t *testing.T) {
	_, err := CPUInstructions(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCPUCycles(t *testing.T) {
	_, err := CPUCycles(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCacheRef(t *testing.T) {
	_, err := CacheRef(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCacheMiss(t *testing.T) {
	_, err := CacheMiss(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBusCycles(t *testing.T) {
	_, err := BusCycles(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestStalledFrontendCycles(t *testing.T) {
	t.Skip()
	_, err := StalledFrontendCycles(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestStalledBackendCycles(t *testing.T) {
	t.Skip()
	_, err := StalledBackendCycles(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCPURefCycles(t *testing.T) {
	_, err := CPURefCycles(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}
func TestCPUClock(t *testing.T) {
	_, err := CPUClock(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}
func TestCPUTaskClock(t *testing.T) {
	_, err := CPUTaskClock(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}
func TestPageFaults(t *testing.T) {
	_, err := PageFaults(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}
func TestContextSwitches(t *testing.T) {
	_, err := ContextSwitches(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCPUMigrations(t *testing.T) {
	_, err := CPUMigrations(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMinorPageFaults(t *testing.T) {
	_, err := MinorPageFaults(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMajorPageFaults(t *testing.T) {
	_, err := MajorPageFaults(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestAlignmentFaults(t *testing.T) {
	_, err := AlignmentFaults(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEmulationFaults(t *testing.T) {
	_, err := EmulationFaults(
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkCPUCycles(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		CPUCycles(
			func() error { return nil },
		)
	}
}
