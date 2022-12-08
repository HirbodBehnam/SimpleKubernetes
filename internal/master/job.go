package master

import (
	"WLF/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type jobStatus uint8

const (
	// Job is not assigned to any slaves
	jobStatusPending jobStatus = iota
	// Job is running in a slave
	jobStatusRunning
	// Job is done with an exit code
	jobStatusDone
	// Job has not even started and has been exited with an error
	jobStatusError
)

// A job which is running in a client
type job struct {
	// The ID of this job
	ID string
	// What is the status of this job
	Status jobStatus
	// In which slave is this job running.
	// Is only meaningful if Status is not jobStatusPending
	SlaveID uint32
	// The job itself
	JobData *proto.NewJobMessage
	// Exit code of the process if Status is jobStatusDone
	ExitCode int32
	// The error of run if Status is jobStatusError
	RunError string
}

// ToJobData will convert a job to a proto.JobData object.
func (j job) ToJobData() *proto.JobData {
	result := &proto.JobData{
		Id:     &proto.UUID{Value: j.ID},
		Cmd:    j.JobData.Cmd,
		NodeId: j.SlaveID,
	}
	switch j.Status {
	case jobStatusPending:
		result.Status = &proto.JobData_Pending{Pending: new(emptypb.Empty)}
	case jobStatusRunning:
		result.Status = &proto.JobData_Running{Running: new(emptypb.Empty)}
	case jobStatusDone:
		result.Status = &proto.JobData_ExitCode{ExitCode: j.ExitCode}
	case jobStatusError:
		result.Status = &proto.JobData_RunError{RunError: j.RunError}
	}
	return result
}

// getJobList will get the job list from master
func (s *Server) getJobList() *proto.JobList {
	s.jobsMutex.RLock()
	defer s.jobsMutex.RUnlock()
	result := &proto.JobList{Jobs: make([]*proto.JobData, 0, len(s.jobs)+len(s.pendingJobs))}
	for _, job := range s.pendingJobs {
		result.Jobs = append(result.Jobs, job.ToJobData())
	}
	for _, job := range s.jobs {
		result.Jobs = append(result.Jobs, job.ToJobData())
	}
	return result
}

func (s *Server) dispatchJob(job *proto.NewJobMessage) error {
	// TODO
	return nil
}
