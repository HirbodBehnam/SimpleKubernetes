package slave

import (
	sigar "github.com/cloudfoundry/gosigar"
	"math/rand"
)

func getFreeMemory() uint64 {
	mem := sigar.Mem{}
	_ = mem.Get() // fuck it
	return mem.Free
}

func getCPUUtilization() uint8 {
	return uint8(rand.Int()%90) + 1
}

func getFreeDisk() uint64 {
	disk := sigar.FileSystemUsage{}
	_ = disk.Get(".")
	return disk.Free
}
