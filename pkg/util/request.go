package util

import (
	"github.com/go-faster/errors"
	protobuf "github.com/golang/protobuf/proto"
	"io"
)

// RequestWithProtobuf will send a protobuffer request into a connection and then expects another one
func RequestWithProtobuf(conn io.ReadWriter, request, response protobuf.Message) error {
	// Send request
	err := WriteProtobuf(conn, request)
	if err != nil {
		return errors.Wrap(err, "cannot send request")
	}
	// Read result
	err = ReadProtobuf(conn, response)
	if err != nil {
		return errors.Wrap(err, "cannot read response")
	}
	return nil
}
