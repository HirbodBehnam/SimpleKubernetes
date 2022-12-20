package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	log "github.com/sirupsen/logrus"
	"io"
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
		s.Salves.AddSlave(SlaveListElement{
			Address: data.Hello.ConnectAddress,
			Id:      slaveID,
			Dead:    false,
		})
		// Dispatch jobs if needed
		go s.dispatchJobsToSlave(slaveID, data.Hello.ConnectAddress)
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
		job.Stdout = util.NewLineLogger(data.JobFinished.Stdout)
		job.Stderr = util.NewLineLogger(data.JobFinished.Stderr)
		s.jobs[data.JobFinished.JobId.Value] = job
		s.jobsMutex.Unlock()
		// Send another job to clients
		go s.dispatchRandomJob()
	}
}

// proxyJobLogs will proxy the job log results from slave to client.
// It will also create a new connection to slave.
func proxyJobLogs(client io.Writer, slaveAddress string, request *proto.GetJobLogsRequest) error {
	// Connect to slave
	slave, err := net.Dial("tcp", slaveAddress)
	if err != nil {
		return errors.Wrap(err, "cannot connect to slave")
	}
	defer slave.Close()
	// Send request
	err = util.WriteProtobuf(slave, &proto.MasterToSlaveRequest{
		Request: &proto.MasterToSlaveRequest_GetJobLogs{
			GetJobLogs: request,
		},
	})
	if err != nil {
		return errors.Wrap(err, "cannot send get log request to slave")
	}
	// Proxy result
	_, _ = io.Copy(client, slave)
	return nil
}
