package prof

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	profileFolder = "stats"
)

func MustProfCPU(fname string) (cancel func()) {
	var err error
	if cancel, err = cpuProf(fname); err == nil {
		return cancel
	}
	panic("ProfCPU panic! " + err.Error())
}

func MustProfMem(fname string) (cancel func()) {
	var err error
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
