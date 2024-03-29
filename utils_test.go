package perf

import (
	"os"
	"runtime"
	"testing"

	"golang.org/x/sys/unix"
)

func TestMaxOpenFiles(t *testing.T) {
	_, err := MaxOpenFiles()
	if err != nil {
		t.Fatal(err)
	}
}

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

func TestL1Data(t *testing.T) {
	_, err := L1Data(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestL1Instructions(t *testing.T) {
	_, err := L1Instructions(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_MISS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestLLCache(t *testing.T) {
	_, err := LLCache(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDataTLB(t *testing.T) {
	_, err := DataTLB(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestInstructionTLB(t *testing.T) {
	_, err := InstructionTLB(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBPU(t *testing.T) {
	_, err := BPU(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeCache(t *testing.T) {
	_, err := NodeCache(
		unix.PERF_COUNT_HW_CACHE_OP_READ,
		unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS,
		func() error { return nil },
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTids(t *testing.T) {
	tids, err := getTids(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	if len(tids) == 1 {
		t.Fatalf("expected multiple threads, got: %+v", tids)
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

func BenchmarkThreadLocking(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		runtime.LockOSThread()
		runtime.UnlockOSThread()
	}
}

func BenchmarkRunBenchmarks(b *testing.B) {
	eventAttrs := []unix.PerfEventAttr{
		CPUInstructionsEventAttr(),
		CPUCyclesEventAttr(),
	}
	RunBenchmarks(
		b,
		func(b *testing.B) {
			for n := 1; n < b.N; n++ {
				a := 42
				for i := 0; i < 1000; i++ {
					a += i
				}
			}
		},
		BenchStrict,
		eventAttrs...,
	)
}

func BenchmarkRunBenchmarksLocked(b *testing.B) {
	eventAttrs := []unix.PerfEventAttr{
		CPUInstructionsEventAttr(),
		CPUCyclesEventAttr(),
	}
	RunBenchmarks(
		b,
		func(b *testing.B) {
			for n := 1; n < b.N; n++ {
				a := 42
				for i := 0; i < 1000; i++ {
					a += i
				}
			}
		},
		BenchLock|BenchStrict,
		eventAttrs...,
	)
}

func BenchmarkBenchmarkTracepointsLocked(b *testing.B) {
	tracepoints := []string{
		"syscalls:sys_enter_getrusage",
	}
	BenchmarkTracepoints(
		b,
		func(b *testing.B) {
			for n := 1; n < b.N; n++ {
				if err := unix.Getrusage(0, &unix.Rusage{}); err != nil {
					b.Fatal(err)
				}
				if err := unix.Getrusage(0, &unix.Rusage{}); err != nil {
					b.Fatal(err)
				}
			}
		},
		BenchLock|BenchStrict,
		tracepoints...,
	)
}

func BenchmarkBenchmarkTracepoints(b *testing.B) {
	tracepoints := []string{
		"syscalls:sys_enter_getrusage",
	}
	BenchmarkTracepoints(
		b,
		func(b *testing.B) {
			for n := 1; n < b.N; n++ {
				if err := unix.Getrusage(0, &unix.Rusage{}); err != nil {
					b.Fatal(err)
				}
				if err := unix.Getrusage(0, &unix.Rusage{}); err != nil {
					b.Fatal(err)
				}
			}
		},
		BenchStrict,
		tracepoints...,
	)
}
