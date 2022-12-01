package util

import (
	"encoding/binary"
	"github.com/go-faster/errors"
	"io"
)

// ReadBigBuffer reads a very big buffer from a reader
func ReadBigBuffer(r io.Reader) ([]byte, error) {
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

// WriteBigBuffer writes a very big buffer into a writer
func WriteBigBuffer(w io.Writer, buffer []byte) error {
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
