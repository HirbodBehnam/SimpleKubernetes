package master

import (
	"WLF/pkg/proto"
	"bytes"
	"encoding/binary"
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
	// Read the job
	job, err := s.readJob(conn)
	if err != nil {
		log.WithError(err).Error("cannot read job of client")
		return
	}
	// Dispatch the job
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

// readJob will read the job from connection
func (s *Server) readJob(conn net.Conn) (*proto.Job, error) {
	// Read the size of payload
	var sizeBuffer [4]byte
	n, err := conn.Read(sizeBuffer[:])
	if err != nil {
		return nil, errors.Wrap(err, "cannot read size of job")
	}
	if n != 4 {
		return nil, errors.New("short read on size")
	}
	// Create buffer
	buffer := make([]byte, binary.BigEndian.Uint32(sizeBuffer[:]))
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read job")
	}
	if n != len(buffer) {
		return nil, errors.New("short read on job")
	}
	// Parse
	job := new(proto.Job)
	err = json.NewDecoder(bytes.NewReader(buffer)).Decode(&job)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse job")
	}
	return job, nil
}
