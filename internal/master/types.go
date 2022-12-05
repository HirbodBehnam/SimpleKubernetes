package master

// SlaveTopResult contains the info about a slave
type SlaveTopResult struct {
	FreeRam       uint64
	FreeDiskSpace uint64
	// How many jobs is this slave running
	RunningJobs uint32
	// Maximum number of jobs this slave can run
	MaxRunningJobs uint32
	// CPU Utilization in percent
	CpuUtilization uint8
}
