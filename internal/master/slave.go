package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	log "github.com/sirupsen/logrus"
	"net"
	"sync/atomic"
)

// A simple counter which will be used to assign ID of slaves
var slaveCounter atomic.Uint32

// getNextSlaveID will increase the internal ID counter and returns the new slave ID
func getNextSlaveID() uint32 {
	return slaveCounter.Add(1)
}

// handleSlave will handle a slave connection
func (s *Server) handleSlave(conn net.Conn) {
	defer conn.Close()
	// Parse
	var request proto.SlaveToMasterRequest
	err := util.ReadProtobuf(conn, &request)
	if err != nil {
		log.WithError(err).Error("cannot read the first message of slave")
		return
	}
	// Send the server ack
	switch data := request.Request.(type) {
	case *proto.SlaveToMasterRequest_Hello:
		slaveID := getNextSlaveID()
		err = util.WriteProtobuf(conn, &proto.SlaveHelloMasterAck{
			Id: slaveID,
		})
		if err != nil {
			log.WithError(err).Error("cannot write server ack for slave")
			return
		}
		// Done. Add to slaves
		s.Salves.AddSlave(util.SlaveListElement{
			Address: data.Hello.ConnectAddress,
			Id:      slaveID,
			Dead:    false,
		})
	case *proto.SlaveToMasterRequest_JobFinished:
		// Update job
		s.jobsMutex.Lock()
		job := s.jobs[data.JobFinished.JobId.Value]
		switch status := data.JobFinished.Status.(type) {
		case *proto.JobFinishedResult_ExitCode:
			job.Status = jobStatusDone
			job.ExitCode = status.ExitCode
		case *proto.JobFinishedResult_RunError:
			job.Status = jobStatusError
			job.RunError = status.RunError
		}
		s.jobs[data.JobFinished.JobId.Value] = job
		s.jobsMutex.Unlock()
		// Send another job to clients
		go s.dispatchRandomJob()
	}
}
