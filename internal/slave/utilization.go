package slave

import (
	"WLF/pkg/proto"
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
	_ = disk.Get(".")
	return disk.Free * 1024
}

// canJobRun will check if a job can be run on this slave or not based on system utilization
func canJobRun(job *proto.NewJobMessage) bool {
	// Check CPU
	if job.NeededCores > getCPUCores() {
		return false
	}

	// Check ram
	if job.NeededMemory > getFreeMemory() {
		return false
	}
	// Check HDD
	if job.NeededSpace > getFreeDisk() {
		return false
	}
	// Done
	return true
}
