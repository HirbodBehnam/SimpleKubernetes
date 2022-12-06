package slave

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	log "github.com/sirupsen/logrus"
	"net"
)

// Slave represents a single slave which can accept jobs from master
type Slave struct {
	MasterAddress string
	MaxTasks      int
	slaveID       uint32
}

// RunSlave is the entry point of the slave
func (s *Slave) RunSlave(listenAddress string) error {
	// At first try to listen on given address
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return errors.Wrap(err, "cannot listen on given address")
	}
	defer l.Close()
	log.WithField("address", l.Addr().String()).Debug("started listening for master")
	// Now connect to master and send the message
	s.slaveID, err = doMasterHandshake(s.MasterAddress, l.Addr().String())
	if err != nil {
		return errors.Wrap(err, "cannot connect to master")
	}
	log.WithField("id", s.slaveID).Debug("connected to master and got the ID")
	// Done. Listen for connections
	for {
		conn, err := l.Accept()
		if err != nil {
			log.WithError(err).Error("cannot accept connection")
			return err
		}
		go s.handleMasterConnection(conn)
	}
}

// doMasterHandshake will do the handshake with master so that it knows that we are a slave
func doMasterHandshake(masterAddress, localAddress string) (uint32, error) {
	// Connect to master
	conn, err := net.Dial("tcp", masterAddress)
	if err != nil {
		return 0, errors.Wrap(err, "cannot dial master")
	}
	defer conn.Close()
	// Send hello message
	messageAck := new(proto.SlaveHelloMasterAck)
	err = util.RequestWithProtobuf(conn, &proto.SlaveToMasterRequest{
		Request: &proto.SlaveToMasterRequest_Hello{
			Hello: &proto.SlaveHello{
				ConnectAddress: localAddress,
			},
		},
	}, messageAck)
	if err != nil {
		return 0, errors.Wrap(err, "cannot do handshake with master")
	}
	return messageAck.Id, nil
}
