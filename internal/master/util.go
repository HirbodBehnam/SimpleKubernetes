package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

// pingSlave will ping the slave to check if it's alive or not
func pingSlave(address string) error {
	// Connect to client
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Wrap(err, "cannot dial slave")
	}
	defer conn.Close()
	// Check status
	result := new(emptypb.Empty)
	err = util.RequestWithProtobuf(conn, &proto.MasterToSlaveRequest{
		Request: &proto.MasterToSlaveRequest_Ping{
			Ping: new(emptypb.Empty),
		},
	}, result)
	if err != nil {
		return errors.Wrap(err, "cannot ping slave")
	}
	return nil
}

// topSlave will get the top of the slave to get utilization of it
func topSlave(address string) (*proto.SlaveTop, error) {
	// Connect to client
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, errors.Wrap(err, "cannot dial slave")
	}
	defer conn.Close()
	// Check status
	result := new(proto.SlaveTop)
	err = util.RequestWithProtobuf(conn, &proto.MasterToSlaveRequest{
		Request: &proto.MasterToSlaveRequest_GetTop{
			GetTop: new(emptypb.Empty),
		},
	}, result)
	if err != nil {
		return nil, errors.Wrap(err, "cannot top slave")
	}
	return result, nil
}
