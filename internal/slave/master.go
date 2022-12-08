package slave

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

// handleMasterConnection will get the request which master wants from us
func (s *Slave) handleMasterConnection(conn net.Conn) {
	defer conn.Close()
	// Get the request
	var request proto.MasterToSlaveRequest
	err := util.ReadProtobuf(conn, &request)
	if err != nil {
		log.WithError(err).Error("cannot read request packet of master")
		return
	}
	// Do the thing
	switch req := request.Request.(type) {
	case *proto.MasterToSlaveRequest_Ping:
		// Just send a message
		err = util.WriteProtobuf(conn, new(emptypb.Empty))
		if err != nil {
			log.WithError(err).Warn("cannot send pong to master")
			return
		}
		log.WithField("remote", conn.RemoteAddr()).Debug("sent pong")
	case *proto.MasterToSlaveRequest_GetTop:
		// Get utilization of system
		err = util.WriteProtobuf(conn, &proto.SlaveTop{
			JobLimit:    s.MaxTasks,
			RunningJobs: 0, // TODO: fill this somehow
			Cores:       getCPUCores(),
			FreeMemory:  getFreeMemory(),
			FreeDisk:    getFreeDisk(),
		})
		if err != nil {
			log.WithError(err).Warn("cannot send top to master")
			return
		}
		log.WithField("remote", conn.RemoteAddr()).Debug("sent top")
	case *proto.MasterToSlaveRequest_NewJob:
		// Check if job is fine to be run on this server (utilization)
		if !canJobRun(req.NewJob.NewJob) {
			err = util.WriteProtobuf(conn, &proto.SlaveJobAck{
				Result: &proto.SlaveJobAck_InsufficientResource{
					InsufficientResource: new(emptypb.Empty),
				},
			})
			if err != nil {
				log.WithError(err).Warn("cannot send ack (insufficient resource) of job to master")
			}
			return
		}
		// Check if this job will exceed the limits of task limit
		s.mu.Lock()
		if s.runningJobCount == s.MaxTasks {
			s.mu.Unlock()
			// Send error
			err = util.WriteProtobuf(conn, &proto.SlaveJobAck{
				Result: &proto.SlaveJobAck_TasksFull{
					TasksFull: new(emptypb.Empty),
				},
			})
			if err != nil {
				log.WithError(err).Warn("cannot send ack (tasks full) of job to master")
			}
			return
		}
		// If everything is fine, increase jobs
		s.runningJobCount++
		s.mu.Unlock()
		// Ack the master that we will run the job
		err = util.WriteProtobuf(conn, &proto.SlaveJobAck{
			Result: &proto.SlaveJobAck_Ack{
				Ack: new(emptypb.Empty),
			},
		})
		if err != nil {
			log.WithError(err).Warn("cannot send ack (job id) of job to master")
			s.reduceRunningJobs() // we failed to run it :(
			return
		}
		// Now run the job
		go s.runJob(req.NewJob.Id.Value, req.NewJob.NewJob)
	case *proto.MasterToSlaveRequest_GetJobLogs:
		// TODO
	}
}
