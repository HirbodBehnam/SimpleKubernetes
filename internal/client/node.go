package client

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/go-faster/errors"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	"strconv"
)

// PrintNodeList will print all nodes from master
func (m *MasterSettings) PrintNodeList() error {
	var result proto.SlavesStatus
	err := util.RequestWithProtobuf(m.conn, &proto.ClientRequest{
		Request: &proto.ClientRequest_NodeList{
			NodeList: new(emptypb.Empty),
		},
	}, &result)
	if err != nil {
		return err
	}
	// Check empty
	if len(result.Status) == 0 {
		fmt.Println("No slaves are connected to master!")
		return nil
	}
	// Print
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Remote Address"})
	for _, entry := range result.Status {
		status := "alive"
		if entry.Dead {
			status = "dead"
		}
		table.Append([]string{"N" + strconv.FormatInt(int64(entry.Id), 10), status, entry.Address})
	}
	table.Render()
	return nil
}

// PrintNodeTop will print nodes status
func (m *MasterSettings) PrintNodeTop() error {
	var result proto.SlavesTop
	err := util.RequestWithProtobuf(m.conn, &proto.ClientRequest{
		Request: &proto.ClientRequest_NodeTop{
			NodeTop: new(emptypb.Empty),
		},
	}, &result)
	if err != nil {
		return err
	}
	// Check if there is any slave
	if len(result.SlaveTops) == 0 {
		fmt.Println("No alive slaves!")
		return nil
	}
	// Create a table from them
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Job Limit", "Active Jobs", "CPU Cores", "Free Memory", "Free Disk"})
	for i := range result.SlaveTops {
		table.Append([]string{
			"N" + strconv.FormatInt(int64(result.SlaveIds[i]), 10),
			strconv.FormatInt(int64(result.SlaveTops[i].JobLimit), 10),
			strconv.FormatInt(int64(result.SlaveTops[i].RunningJobs), 10),
			strconv.FormatInt(int64(result.SlaveTops[i].Cores), 10),
			humanize.IBytes(result.SlaveTops[i].FreeMemory),
			humanize.IBytes(result.SlaveTops[i].FreeDisk),
		})
	}
	table.Render()
	return nil
}

// LogsOfJob will print logs of a job
func (m *MasterSettings) LogsOfJob(jobID string, count uint32, live, tail, stderr bool) error {
	// Create the request
	request := &proto.GetJobLogsRequest{
		JobId:     &proto.UUID{Value: jobID},
		LineCount: count,
		Stderr:    stderr,
	}
	if live {
		request.Type = proto.GetJobLogsRequestType_LIVE
	} else if tail {
		request.Type = proto.GetJobLogsRequestType_TAIL
	} else {
		request.Type = proto.GetJobLogsRequestType_HEAD
	}
	// Send the request
	err := util.WriteProtobuf(m.conn, &proto.ClientRequest{
		Request: &proto.ClientRequest_JobLog{
			JobLog: request,
		},
	})
	if err != nil {
		return errors.Wrap(err, "cannot send logs request")
	}
	// Stream logs if needed
	if live {
		return streamLiveLogs(m.conn)
	}
	// Otherwise just read the data
	return readNextLog(m.conn)
}

// streamLiveLogs will write live logs of user
func streamLiveLogs(r io.Reader) error {
	for {
		if err := readNextLog(r); err != nil {
			return err
		}
	}
}

// readNextLog will read the next log from a reader
func readNextLog(r io.Reader) error {
	var logs proto.GetJobLogsResult
	err := util.ReadProtobuf(r, &logs)
	if err != nil {
		return errors.Wrap(err, "cannot read logs response")
	}
	switch data := logs.Result.(type) {
	case *proto.GetJobLogsResult_Error:
		return errors.New(data.Error)
	case *proto.GetJobLogsResult_Results:
		for _, line := range data.Results.Logs {
			fmt.Println(line)
		}
	}
	return nil
}
