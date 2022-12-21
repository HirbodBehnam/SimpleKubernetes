package util

import (
	"WLF/pkg/proto"
	"encoding/json"
	"github.com/go-faster/errors"
	"io"
	"strings"
)

// SendJobRequest will send a job request to master
func SendJobRequest(w io.Writer, jobConfig []byte, payloadName string, payloadData []byte) error {
	// Send initial request
	request := &proto.ClientRequest{
		Request: &proto.ClientRequest_NewJob{
			NewJob: &proto.ClientNewJobMessage{
				PayloadName: stringNilOnEmpty(payloadName),
			},
		},
	}
	err := WriteProtobuf(w, request)
	if err != nil {
		return errors.Wrap(err, "cannot send metadata of job")
	}
	// Now send the json payload
	err = writeBigBuffer(w, jobConfig)
	if err != nil {
		return errors.Wrap(err, "cannot send job json")
	}
	// If needed, send the payload data
	if payloadName != "" && payloadData != nil {
		err = writeBigBuffer(w, payloadData)
		if err != nil {
			return errors.Wrap(err, "cannot send job payload")
		}
	}
	return nil
}

// ReadJobRequest will parse the request sent with SendJobRequest (except it's first message)
// as proto.NewJobMessage
func ReadJobRequest(r io.Reader, payloadName *string) (*proto.NewJobMessage, error) {
	result := new(proto.NewJobMessage)
	// Get the next message as the job json
	data, err := readBigBuffer(r)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read job json")
	}
	var jobData intermediateJobResult
	err = json.Unmarshal(data, &jobData)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse job json")
	}
	// Populate the result
	result.ProgramName, result.Arguments, err = parseCmd(jobData.Cmd)
	result.NeededCores = jobData.Cores
	result.NeededSpace = jobData.Disk
	result.NeededMemory = jobData.Memory
	// Get the payload if needed
	if payloadName != nil {
		payload, err := readBigBuffer(r)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read big payload")
		}
		result.Payload = &proto.PayloadProgram{
			ProgramBin:  payload,
			ProgramName: *payloadName,
		}
	}
	// Done
	return result, nil
}

func stringNilOnEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

type intermediateJobResult struct {
	Cmd    string `json:"cmd,omitempty"`
	Cores  uint32 `json:"cores"`
	Memory uint64 `json:"memory"`
	Disk   uint64 `json:"disk"`
}

// parseCmd will parse the command given to us from job request
func parseCmd(cmd string) (program string, arguments []string, err error) {
	if cmd == "" {
		return "", nil, errors.New("empty command")
	}
	program, args, found := strings.Cut(cmd, " ")
	if !found {
		arguments = make([]string, 0)
		return
	}
	arguments = strings.Split(args, " ")
	return
}
