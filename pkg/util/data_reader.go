package util

import (
	"encoding/binary"
	"github.com/go-faster/errors"
	protobuf "google.golang.org/protobuf/proto"
	"io"
)

// readBigBuffer reads a very big buffer from a reader
func readBigBuffer(r io.Reader) ([]byte, error) {
	// Get size
	var sizeBuffer [4]byte
	n, err := r.Read(sizeBuffer[:])
	if err != nil {
		return nil, errors.Wrap(err, "cannot read size of buffer")
	}
	if n != 4 {
		return nil, errors.New("short read on size")
	}
	// Create buffer
	buffer := make([]byte, binary.BigEndian.Uint32(sizeBuffer[:]))
	n, err = r.Read(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read buffer")
	}
	if n != len(buffer) {
		return nil, errors.New("short read on buffer")
	}
	return buffer, nil
}

// writeBigBuffer writes a very big buffer into a writer
func writeBigBuffer(w io.Writer, buffer []byte) error {
	// Send the size
	var sizeBuffer [4]byte
	binary.BigEndian.PutUint32(sizeBuffer[:], uint32(len(buffer)))
	n, err := w.Write(sizeBuffer[:])
	if err != nil {
		return errors.Wrap(err, "cannot write size of buffer")
	}
	if n != 4 {
		return errors.New("short write on size")
	}
	// Send the buffer
	n, err = w.Write(buffer)
	if err != nil {
		return errors.Wrap(err, "cannot write buffer")
	}
	if n != len(buffer) {
		return errors.New("short write on buffer")
	}
	return nil
}

// ReadProtobuf will read a protobuf message from a stream using readBigBuffer
func ReadProtobuf(r io.Reader, message protobuf.Message) error {
	buffer, err := readBigBuffer(r)
	if err != nil {
		return errors.Wrap(err, "cannot read data")
	}
	err = protobuf.Unmarshal(buffer, message)
	if err != nil {
		return errors.Wrap(err, "cannot parse proto")
	}
	return nil
}

// WriteProtobuf will write a protobuf message in a writer using WriteBigBuffer
func WriteProtobuf(w io.Writer, message protobuf.Message) error {
	data, err := protobuf.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "cannot marshal message")
	}
	err = writeBigBuffer(w, data)
	if err != nil {
		return errors.Wrap(err, "cannot write data")
	}
	return nil
}
