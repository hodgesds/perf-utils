package skylake

import (
	"testing"

	"github.com/hodgesds/perf-utils"
	"golang.org/x/sys/unix"
)

func TestSkylake(t *testing.T) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_RAW,
		Config:      0x1B0,
		Size:        perf.EventAttrSize,
		Bits:        unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	_, err := unix.PerfEventOpen(
		eventAttr,
		unix.Gettid(),
		-1,
		-1,
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
}
