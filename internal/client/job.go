package client

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"fmt"
	"github.com/go-faster/errors"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"path/filepath"
	"strconv"
)

// PrintJobList will request and print all jobs in master
func (m *MasterSettings) PrintJobList() error {
	// Request data
	var result proto.JobList
	err := util.RequestWithProtobuf(m.conn, &proto.ClientRequest{
		Request: &proto.ClientRequest_JobList{
			JobList: new(emptypb.Empty),
		},
	}, &result)
	if err != nil {
		return err
	}
	// Check empty jobs
	if len(result.Jobs) == 0 {
		fmt.Println("No jobs are available!")
		return nil
	}
	// Print the jobs
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Job ID", "CMD", "Status", "Node ID", "Exit Code/Run Error"})
	for _, entry := range result.Jobs {
		nodeID, status, reason := "-", "", "-"
		switch data := entry.Status.(type) {
		case *proto.JobData_Pending:
			status = "Pending"
		case *proto.JobData_Running:
			status = "Running"
			nodeID = "N" + strconv.FormatUint(uint64(entry.NodeId), 10)
		case *proto.JobData_ExitCode:
			status = "Exited"
			nodeID = "N" + strconv.FormatUint(uint64(entry.NodeId), 10)
			reason = strconv.FormatInt(int64(data.ExitCode), 10)
		case *proto.JobData_RunError:
			status = "Run Error"
			nodeID = "N" + strconv.FormatUint(uint64(entry.NodeId), 10)
			reason = data.RunError
		}
		table.Append([]string{
			entry.Id.Value,
			entry.Cmd,
			status,
			nodeID,
			reason,
		})
	}
	table.Render()
	return nil
}

// AddJob will get a job and dispatch it to master
func (m *MasterSettings) AddJob(configPath, executablePath string, neededMemory, neededDisk *uint64, neededCores *uint32) error {
	// Read the config file
	config, err := readJobConfigJson(configPath)
	if err != nil {
		return errors.Wrap(err, "cannot read job config file")
	}
	// Read the executable if needed
	var executable []byte
	if executablePath != "" {
		executable, err = os.ReadFile(executablePath)
		if err != nil {
			return errors.Wrap(err, "cannot read the executable")
		}
	}
	// Create job
	job := &proto.NewJobMessage{
		Cmd:          config.Cmd,
		NeededMemory: neededMemory,
		NeededCores:  neededCores,
		NeededSpace:  neededDisk,
	}
	if executable != nil {
		job.Program = &proto.PayloadProgram{
			ProgramBin:  executable,
			ProgramName: filepath.Base(executablePath),
		}
	}
	// Send job
	var jobID proto.UUID
	err = util.RequestWithProtobuf(
		m.conn,
		&proto.ClientRequest{Request: &proto.ClientRequest_NewJob{NewJob: job}},
		&jobID,
	)
	if err != nil {
		return errors.Wrap(err, "cannot dispatch job")
	}
	// Show ID
	fmt.Println("Job dispatched with ID", jobID.Value)
	return nil
}
