package client

import (
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	protobuf "github.com/golang/protobuf/proto"
	"io"
)

func requestWithProtobuf(conn io.ReadWriter, request, response protobuf.Message) error {
	// Send request
	err := util.WriteProtobuf(conn, request)
	if err != nil {
		return errors.Wrap(err, "cannot send request")
	}
	// Read result
	err = util.ReadProtobuf(conn, response)
	if err != nil {
		return errors.Wrap(err, "cannot read response")
	}
	return nil
}
