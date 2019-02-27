package perf

import (
	"encoding/binary"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

var (
	// EventAttrSize is the size of a PerfEventAttr
	EventAttrSize = uint32(unsafe.Sizeof(unix.PerfEventAttr{}))
)

// profileFn is a helper function to profile a function.
func profileFn(eventAttr *unix.PerfEventAttr, f func() error) (*ProfileValue, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	fd, err := unix.PerfEventOpen(
		eventAttr,
		unix.Gettid(),
		-1,
		-1,
		0,
	)
	if err != nil {
		return nil, err
	}
	if err := unix.IoctlSetInt(fd, unix.PERF_EVENT_IOC_RESET, 0); err != nil {
		return nil, err
	}
	if err := unix.IoctlSetInt(fd, unix.PERF_EVENT_IOC_ENABLE, 0); err != nil {
		return nil, err
	}
	if err := f(); err != nil {
		return nil, err
	}
	if err := unix.IoctlSetInt(fd, unix.PERF_EVENT_IOC_DISABLE, 0); err != nil {
		return nil, err
	}
	buf := make([]byte, 24)
	if _, err := syscall.Read(fd, buf); err != nil {
		return nil, err
	}
	return &ProfileValue{
		Value:       binary.LittleEndian.Uint64(buf[0:8]),
		TimeEnabled: binary.LittleEndian.Uint64(buf[8:16]),
		TimeRunning: binary.LittleEndian.Uint64(buf[16:24]),
	}, nil
}

// CPUInstructions is used to profile a function and return the number of CPU instructions.
// Note that it will call runtime.LockOSThread to ensure accurate profilng.
func CPUInstructions(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_INSTRUCTIONS,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CPUCycles is used to profile a function and return the number of CPU cycles.
// Note that it will call runtime.LockOSThread to ensure accurate profilng.
func CPUCycles(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_CPU_CYCLES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CacheRef is used to profile a function and return the number of cache
// references. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func CacheRef(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_CACHE_REFERENCES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CacheMiss is used to profile a function and return the number of cache
// misses. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func CacheMiss(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_CACHE_MISSES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// BusCycles is used to profile a function and return the number of bus
// cycles. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func BusCycles(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_BUS_CYCLES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// StalledFrontendCycles is used to profile a function and return the number of
// stalled frontend cycles. Note that it will call runtime.LockOSThread to
// ensure accurate profilng.
func StalledFrontendCycles(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_STALLED_CYCLES_FRONTEND,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// StalledBackendCycles is used to profile a function and return the number of
// stalled backend cycles. Note that it will call runtime.LockOSThread to
// ensure accurate profilng.
func StalledBackendCycles(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_STALLED_CYCLES_BACKEND,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CPURefCycles is used to profile a function and return the number of CPU
// references cycles which are not affected by frequency scaling. Note that it
// will call runtime.LockOSThread to ensure accurate profilng.
func CPURefCycles(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_HARDWARE,
		Config:      unix.PERF_COUNT_HW_REF_CPU_CYCLES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CPUClock is used to profile a function and return the CPU clock timer. Note
// that it will call runtime.LockOSThread to ensure accurate profilng.
func CPUClock(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_CPU_CLOCK,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CPUTaskClock is used to profile a function and return the CPU clock timer
// for the running task. Note that it will call runtime.LockOSThread to ensure
// accurate profilng.
func CPUTaskClock(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_TASK_CLOCK,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// PageFaults is used to profile a function and return the number of page
// faults. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func PageFaults(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_PAGE_FAULTS,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// ContextSwitches is used to profile a function and return the number of
// context switches. Note that it will call runtime.LockOSThread to ensure
// accurate profilng.
func ContextSwitches(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_CONTEXT_SWITCHES,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// CPUMigrations is used to profile a function and return the number of times
// the thread has been migrated to a new CPU. Note that it will call
// runtime.LockOSThread to ensure accurate profilng.
func CPUMigrations(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_CPU_MIGRATIONS,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// MinorPageFaults is used to profile a function and return the number of minor
// page faults. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func MinorPageFaults(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_PAGE_FAULTS_MIN,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// MajorPageFaults is used to profile a function and return the number of major
// page faults. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func MajorPageFaults(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_PAGE_FAULTS_MAJ,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// AlignmentFaults is used to profile a function and return the number of alignment
// faults. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func AlignmentFaults(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_ALIGNMENT_FAULTS,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}

// EmulationFaults is used to profile a function and return the number of emulation
// faults. Note that it will call runtime.LockOSThread to ensure accurate
// profilng.
func EmulationFaults(f func() error) (*ProfileValue, error) {
	eventAttr := &unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_SOFTWARE,
		Config:      unix.PERF_COUNT_SW_EMULATION_FAULTS,
		Size:        EventAttrSize,
		Bits:        unix.PerfBitDisabled | unix.PerfBitExcludeKernel | unix.PerfBitExcludeHv,
		Read_format: unix.PERF_FORMAT_TOTAL_TIME_RUNNING | unix.PERF_FORMAT_TOTAL_TIME_ENABLED,
	}
	return profileFn(eventAttr, f)
}
