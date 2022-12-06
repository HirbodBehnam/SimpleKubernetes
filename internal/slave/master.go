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
			log.WithError(err).Warn("cannot send pong to client")
			return
		}
	case *proto.MasterToSlaveRequest_GetTop:
		// Get utilization of system
		err = util.WriteProtobuf(conn, &proto.SlaveTop{
			JobLimit:    uint32(s.MaxTasks),
			RunningJobs: 0, // TODO: fill this somehow
			Cores:       getCPUCores(),
			FreeMemory:  getFreeMemory(),
			FreeDisk:    getFreeDisk(),
		})
		if err != nil {
			log.WithError(err).Warn("cannot send pong to client")
			return
		}
	case *proto.MasterToSlaveRequest_NewJob:
		_ = req
	case *proto.MasterToSlaveRequest_GetJobLogs:

	}
}
