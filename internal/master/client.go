package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	log "github.com/sirupsen/logrus"
	"net"
)

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()
	// Authorize user
	if !s.authorizeClient(conn) {
		return
	}
	// Read the command
	var command proto.ClientRequest
	err := util.ReadProtobuf(conn, &command)
	if err != nil {
		log.WithError(err).Error("cannot read job of client")
		return
	}
	// Get type of command
	switch data := command.Request.(type) {
	case *proto.ClientRequest_NewJob: // send a job to client
		if err = s.dispatchJob(data.NewJob); err != nil {
			log.WithError(err).Error("cannot dispatch job")
		}
		return
	case *proto.ClientRequest_JobList:
		err = util.WriteProtobuf(conn, s.getJobList())
		if err != nil {
			log.WithError(err).Warn("cannot write job list result")
		}
		return
	case *proto.ClientRequest_JobLog:
	case *proto.ClientRequest_NodeList:
		err = util.WriteProtobuf(conn, s.getNodeStatus())
		if err != nil {
			log.WithError(err).Warn("cannot write node list result")
		}
		return
	case *proto.ClientRequest_NodeTop:
		err = util.WriteProtobuf(conn, s.getNodeTop())
		if err != nil {
			log.WithError(err).Warn("cannot write node top result")
		}
		return
	}
}

// authorizeClient will read the first message from stream and authorize the user
func (s *Server) authorizeClient(conn net.Conn) bool {
	var auth proto.ClientAuthorization
	err := util.ReadProtobuf(conn, &auth)
	if err != nil {
		log.WithError(err).Error("cannot read client hello")
		return false
	}
	// Validate
	pass, exists := s.Users[auth.Username]
	var credOk bool
	if exists && pass == auth.Password {
		log.WithField("cred", &auth).Info("user logged in")
		credOk = true
	} else {
		log.WithField("cred", &auth).Warn("invalid creds")
		credOk = false
	}
	// Send result
	err = util.WriteProtobuf(conn, &proto.ClientAuthorizationResult{Ok: credOk})
	if err != nil {
		log.WithError(err).Error("cannot write back auth result")
		return false
	}
	return credOk
}

// getNodeStatus gets all of nodes status
func (s *Server) getNodeStatus() *proto.SlavesStatus {
	slaves := s.Salves.ToList()
	result := new(proto.SlavesStatus)
	// Check alive status
	for i, slave := range slaves {
		// Add results
		result.Status = append(result.Status, &proto.SlaveStatus{
			Id:      slave.Id,
			Address: slave.Address,
			Dead:    slave.Dead,
		})
		// Don't worry about the dead ones
		if slave.Dead {
			continue
		}
		// Check status of live ones
		err := pingSlave(slave.Address)
		if err != nil {
			log.WithError(err).WithField("id", slave.Id).WithField("address", slave.Address).
				Warn("slave died")
			s.Salves.MakeDead(slave.Address)
			slave.Dead = true
			result.Status[i].Dead = true
		}
	}
	return result
}

// getNodeStatus gets all of nodes status
func (s *Server) getNodeTop() *proto.SlavesTop {
	slaves := s.Salves.ToList()
	result := new(proto.SlavesTop)
	// Check alive status
	for _, slave := range slaves {
		// Don't worry about the dead ones
		if slave.Dead {
			continue
		}
		// Get top of live ones
		top, err := topSlave(slave.Address)
		if err != nil {
			log.WithError(err).WithField("id", slave.Id).WithField("address", slave.Address).
				Warn("slave died")
			s.Salves.MakeDead(slave.Address)
			continue
		}
		// Add to result
		result.SlaveIds = append(result.SlaveIds, slave.Id)
		result.SlaveTops = append(result.SlaveTops, top)
	}
	return result
}
