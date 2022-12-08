package master

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
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

func (s *Server) dispatchJob(j *proto.NewJobMessage) error {
	jobID := uuid.NewString()
	// Lock jobs to edit them
	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()
	// Save job
	savedJob := job{
		ID:      jobID,
		Status:  jobStatusPending,
		JobData: j,
	}
	s.pendingJobs[jobID] = savedJob
	// Get list of slaves
	s.tryDispatchJobToSlaves(savedJob)
	return nil
}

// dispatchRandomJob will send a random job to all slaves and checks if they can accept it
func (s *Server) dispatchRandomJob() {
	// Lock jobs to edit them
	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()
	// Loop each job
	for _, j := range s.pendingJobs {
		s.tryDispatchJobToSlaves(j)
	}
}

// tryDispatchJobToSlaves will try to send a job to slaves. It updates
// s.pendingJobs and s.jobs if job is dispatched.
// Caller must lock s.jobsMutex before calling this thread
func (s *Server) tryDispatchJobToSlaves(j job) {
	slaves := s.Salves.ToList()
	for _, slave := range slaves {
		// We don't want to send jobs to dead slaves
		if slave.Dead {
			continue
		}
		// Otherwise try to send the job to this slave
		err := dispatchJobToSlave(j, slave.Address)
		if err == TasksFullErr || err == InsufficientResourceErr { // we don't case
			continue
		}
		if err != nil { // dead slave :(
			log.WithError(err).WithField("slave", slave).Warn("slave died")
			s.Salves.MakeDead(slave.Address)
			continue
		}
		// We dispatched the job! Yay
		log.WithField("jobID", j.ID).WithField("slaveID", slave.Id).Info("dispatched job")
		delete(s.pendingJobs, j.ID)
		j.Status = jobStatusRunning
		j.SlaveID = slave.Id
		j.JobData.Program = nil // free memory
		s.jobs[j.ID] = j
		break
	}
}

// dispatchJobToSlave tries to send a job to a slave.
// It might return TasksFullErr or InsufficientResourceErr if slaves says so.
// Otherwise, an error is returned and the slave is probably dead.
func dispatchJobToSlave(j job, slaveAddress string) error {
	// Connect to slave
	conn, err := net.Dial("tcp", slaveAddress)
	if err != nil {
		return errors.Wrap(err, "cannot dial slave")
	}
	defer conn.Close()
	// Connect to slave and send the request
	var response proto.SlaveJobAck
	err = util.RequestWithProtobuf(conn, &proto.MasterToSlaveRequest{
		Request: &proto.MasterToSlaveRequest_NewJob{
			NewJob: &proto.SlaveNewJobRequest{
				Id:     &proto.UUID{Value: j.ID},
				NewJob: j.JobData,
			},
		},
	}, &response)
	if err != nil {
		return err
	}
	// Check result of job
	if _, tasksFull := response.Result.(*proto.SlaveJobAck_TasksFull); tasksFull {
		return TasksFullErr
	}
	if _, insufficientResource := response.Result.(*proto.SlaveJobAck_InsufficientResource); insufficientResource {
		return InsufficientResourceErr
	}
	return nil
}
