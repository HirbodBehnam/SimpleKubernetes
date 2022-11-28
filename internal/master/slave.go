package master

import "sync/atomic"

// A simple counter which will be used to assign ID of slaves
var slaveCounter atomic.Uint32

// getNextSlaveID will increase the internal ID counter and returns the new slave ID
func getNextSlaveID() uint32 {
	return slaveCounter.Add(1)
}
