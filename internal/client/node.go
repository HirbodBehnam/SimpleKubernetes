package client

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/protobuf/types/known/emptypb"
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
