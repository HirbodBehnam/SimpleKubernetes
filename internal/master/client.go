package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	protobuf "github.com/golang/protobuf/proto"
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
		// TODO: handle full slaves
		if err = s.dispatchJob(data.NewJob); err != nil {
			log.WithError(err).Error("cannot dispatch job")
		}
		return
	case *proto.ClientRequest_JobList: // Get status of all nodes
		err = util.WriteProtobuf(conn, s.getNodeStatus())
		if err != nil {
			log.WithError(err).Warn("cannot write node list result")
		}
		return
	case *proto.ClientRequest_JobLog:
	case *proto.ClientRequest_NodeList:
	case *proto.ClientRequest_NodeTop:
	}
}

// authorizeClient will read the first message from stream and authorize the user
func (s *Server) authorizeClient(conn net.Conn) bool {
	// Read the message
	data, err := util.ReadBigBuffer(conn)
	if err != nil {
		log.WithError(err).Warn("cannot read buffer in clients first message")
		return false
	}
	// Parse it
	var auth proto.ClientAuthorization
	err = protobuf.Unmarshal(data, &auth)
	if err != nil {
		log.WithError(err).Warn("cannot parse authorization packet")
		return false
	}
	// Validate
	pass, exists := s.Users[auth.Username]
	if exists && pass == auth.Password {
		log.WithField("cred", &auth).Info("user logged in")
		return true
	}
	log.WithField("cred", &auth).Warn("invalid creds")
	return false
}

func (s *Server) dispatchJob(job *proto.NewJobMessage) error {
	// TODO
	return nil
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
