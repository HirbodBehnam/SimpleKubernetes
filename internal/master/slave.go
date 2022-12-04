package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"encoding/json"
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

// handleSlaveHello will add the slave to our pool
func (s *Server) handleSlaveHello(conn net.Conn) {
	defer conn.Close()
	// Get the hello message
	data, err := util.ReadBigBuffer(conn)
	if err != nil {
		log.WithError(err).Error("cannot read slave hello")
		return
	}
	// Parse
	dataParsed := new(proto.SlaveToMasterRequest)
	err = json.Unmarshal(data, dataParsed)
	if err != nil {
		log.WithError(err).Error("cannot parse slave hello")
		return
	}
	// Send the server ack
	if dataParsed.Hello != nil {
		slaveID := getNextSlaveID()
		data, _ = json.Marshal(&proto.SlaveHelloMasterAck{
			Id: slaveID,
		})
		err = util.WriteBigBuffer(conn, data)
		if err != nil {
			log.WithError(err).Error("cannot write server ack for slave")
			return
		}
		// Done. Add to slaves
		s.Salves.AddSlave(util.SlaveListElement{
			Address: dataParsed.Hello.ConnectAddress,
			Id:      slaveID,
			Dead:    false,
		})
		return
	}
}
