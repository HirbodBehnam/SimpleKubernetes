package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"bytes"
	"encoding/json"
	"github.com/go-faster/errors"
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
	command, err := s.readRequest(conn)
	if err != nil {
		log.WithError(err).Error("cannot read job of client")
		return
	}
	// Get type of command
	if command.NewJob != nil {
		if err = s.dispatchJob(command.NewJob); err != nil {
			log.WithError(err).Error("cannot dispatch job")
		}
		return
	}
	// Get status of all nodes
	if command.NodeList != nil {
		data, _ := json.Marshal(s.getNodeStatus())
		err = util.WriteBigBuffer(conn, data)
		if err != nil {
			log.WithError(err).Warn("cannot write node list result")
		}
		return
	}
}

// authorizeClient will read the first message from stream and authorize the user
func (s *Server) authorizeClient(conn net.Conn) bool {
	// Read the message
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.WithError(err).Warn("cannot read buffer in clients first message")
		return false
	}
	// Parse it
	var auth proto.ClientAuthorization
	err = json.NewDecoder(bytes.NewReader(buffer[:n])).Decode(&auth)
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

// readRequest will read the request from client connection
func (s *Server) readRequest(conn net.Conn) (*proto.ClientRequest, error) {
	// Read big buffer
	buffer, err := util.ReadBigBuffer(conn)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read request")
	}
	// Parse
	job := new(proto.ClientRequest)
	err = json.NewDecoder(bytes.NewReader(buffer)).Decode(&job)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse job")
	}
	return job, nil
}

func (s *Server) dispatchJob(job *proto.NewJobMessage) error {
	// TODO
	return nil
}

// getNodeStatus gets all of nodes status
func (s *Server) getNodeStatus() []util.SlaveListElement {
	slaves := s.Salves.ToList()
	// Check alive status
	for _, slave := range slaves {
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
		}
	}
	return slaves
}
