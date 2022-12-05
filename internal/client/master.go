package client

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	"net"
)

var AlreadyAuthorizedErr = errors.New("already authorized")
var BadAuthErr = errors.New("bad credentials")

// MasterSettings is the info needed to connect to master
type MasterSettings struct {
	// The address to connect to
	Address  string
	Username string
	Password string
	// The TCP connection to master
	conn net.Conn
}

// Auth will authorize the user
func (m *MasterSettings) Auth() error {
	// Check already authorized
	if m.conn != nil {
		return AlreadyAuthorizedErr
	}
	// Connect to master
	c, err := net.Dial(m.Address, "tcp")
	if err != nil {
		return errors.Wrap(err, "cannot dial master")
	}
	// Authorize
	err = util.WriteProtobuf(c, &proto.ClientAuthorization{
		Username: m.Username,
		Password: m.Password,
	})
	if err != nil {
		_ = c.Close()
		return errors.Wrap(err, "cannot send auth data")
	}
	// Check status
	var authResult proto.ClientAuthorizationResult
	err = util.ReadProtobuf(c, &authResult)
	if err != nil {
		_ = c.Close()
		return errors.Wrap(err, "cannot read auth result")
	}
	if !authResult.Ok {
		_ = c.Close()
		return BadAuthErr
	}
	// Apply
	m.conn = c
	return nil
}

// Close will close the socket if it's not open
func (m *MasterSettings) Close() {
	if m.conn != nil {
		_ = m.conn.Close()
	}
}
