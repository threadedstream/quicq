package prof

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

const (
	profileFolder = "stats"
)

// MustProfCPU runs cpu profiler
func MustProfCPU(fname string) (cancel func()) {
	var err error

	now := time.Now()
	parts := strings.Split(fname, ".")
	fname = fmt.Sprintf("%s_%d%d%d.%s", parts[0], now.Hour(), now.Minute(), now.Second(), parts[1])
	if cancel, err = cpuProf(fname); err == nil {
		return cancel
	}
	panic("ProfCPU panic! " + err.Error())
}

// MustProfMem runs memory profiler
func MustProfMem(fname string) (cancel func()) {
	var err error

	now := time.Now()
	parts := strings.Split(fname, ".")
	fname = fmt.Sprintf("%s_%d%d%d.%s", parts[0], now.Hour(), now.Minute(), now.Second(), parts[1])
	if cancel, err = memProf(fname); err == nil {
		return cancel
	}
	panic("ProfMem panic! " + err.Error())
}

func cpuProf(fname string) (cancel func(), err error) {
	f, err := os.Create(fmt.Sprintf("%s/%s", profileFolder, fname))
	if err != nil {
		return nil, err
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		return nil, err
	}
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}, nil
}

func memProf(fname string) (cancel func(), err error) {
	f, err := os.Create(fmt.Sprintf("%s/%s", profileFolder, fname))
	if err != nil {
		return nil, err
	}
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		return nil, err
	}
	return func() {
		f.Close()
	}, nil
}
