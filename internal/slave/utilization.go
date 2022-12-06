package slave

import (
	sigar "github.com/cloudfoundry/gosigar"
	"runtime"
)

func getFreeMemory() uint64 {
	mem := sigar.Mem{}
	_ = mem.Get() // fuck it
	return mem.Free
}

func getCPUCores() uint32 {
	return uint32(runtime.NumCPU())
}

func getFreeDisk() uint64 {
	disk := sigar.FileSystemUsage{}
	_ = disk.Get("C://")
	return disk.Free * 1024
}
