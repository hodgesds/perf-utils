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

func BenchmarkCPUCycles(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		CPUCycles(
			func() error { return nil },
		)
	}
}
