package perf

import (
	"runtime"
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

func BenchmarkCPUInstructions(b *testing.B) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	perfCtx, err := NewContext(0)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		perfCtx.Instructions(
			func() error { return nil },
		)
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
