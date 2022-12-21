package slave

import (
	"WLF/pkg/proto"
	"WLF/pkg/util"
	log "github.com/sirupsen/logrus"
	protobuf "google.golang.org/protobuf/proto"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// A job in slave
type slaveJob struct {
	// When was this job started
	started time.Time
	// When was this job ended.
	// This value can be used to check if job is done or not
	ended time.Time
	// stdout logs (each element represents a line)
	stdout util.LineLogger
	// stderr logs (each element represents a line)
	stderr util.LineLogger
}

// Dead will check if the job is dead or not
func (j *slaveJob) Dead() bool {
	return !j.ended.IsZero()
}

// runJob will get a job and run it
func (s *Slave) runJob(id string, requestedJob *proto.NewJobMessage) {
	log.WithField("id", id).Info("running job")
	// At first add the job to list of jobs
	job := &slaveJob{
		started: time.Now(),
	}
	s.mu.Lock()
	s.jobs[id] = job
	s.mu.Unlock()
	// Create a temp folder to run the application in there
	sandboxFolder, err := os.MkdirTemp("", "slave"+strconv.FormatUint(uint64(s.slaveID), 10)+"sandbox"+id)
	if err != nil {
		log.WithError(err).Error("cannot create sandbox for job")
		s.jobDone(id, job, 0, protobuf.String("cannot create sandbox: "+err.Error()))
		return
	}
	defer os.RemoveAll(sandboxFolder)
	// Save the application in sandbox if needed
	if requestedJob.Payload != nil {
		err = os.WriteFile(filepath.Join(sandboxFolder, requestedJob.Payload.ProgramName), requestedJob.Payload.ProgramBin, 0666)
		if err != nil {
			log.WithError(err).Error("cannot save embedded program")
			s.jobDone(id, job, 0, protobuf.String("cannot save embedded program: "+err.Error()))
			return
		}
	}
	// Finally, run the program
	cmd := exec.Command(requestedJob.ProgramName, requestedJob.Arguments...)
	cmd.Dir = sandboxFolder
	cmd.Stdout = &job.stdout
	cmd.Stderr = &job.stderr
	// Run it and wait
	err = cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.WithField("jobID", id).WithField("exitCode", exiterr.ExitCode()).Info("job done")
			s.jobDone(id, job, exiterr.ExitCode(), nil)
		} else {
			log.WithField("jobID", id).WithError(err).Info("cannot run job")
			s.jobDone(id, job, 0, protobuf.String("cannot run job: "+err.Error()))
		}
		return
	}
	// Finally!
	log.WithField("jobID", id).Info("job done")
	s.jobDone(id, job, 0, nil)
}

// jobDone will do anything which is needed after we finish a job.
// At first, it will decrease the count of running jobs.
// Then, it will set the ended time of the job.
// After, it will send the job result to master.
// If runError is not null, it will send a job failed message to master.
func (s *Slave) jobDone(jobID string, job *slaveJob, exitCode int, runError *string) {
	s.reduceRunningJobs()  // we finished a job
	job.ended = time.Now() // we ended it now
	// we send the result to master that we are done
	masterData := &proto.JobFinishedResult{
		SlaveId: s.slaveID,
		JobId:   &proto.UUID{Value: jobID},
		Stdout:  job.stdout.Copy(),
		Stderr:  job.stderr.Copy(),
	}
	if runError != nil {
		masterData.Status = &proto.JobFinishedResult_RunError{RunError: *runError}
	} else {
		masterData.Status = &proto.JobFinishedResult_ExitCode{ExitCode: int32(exitCode)}
	}
	s.sendJobResult(masterData)
}

// senJobResult will send a job result to master
func (s *Slave) sendJobResult(result *proto.JobFinishedResult) {
	// Create connection to master
	conn, err := net.Dial("tcp", s.MasterAddress)
	if err != nil {
		log.WithError(err).Error("cannot dial master to send job result")
		return
	}
	defer conn.Close()
	// Send the result and done
	err = util.WriteProtobuf(conn, &proto.SlaveToMasterRequest{
		Request: &proto.SlaveToMasterRequest_JobFinished{
			JobFinished: result,
		},
	})
	if err != nil {
		log.WithError(err).Error("cannot send job result to master")
		return
	}
}

// reduceRunningJobs will reduce the running job count by one atomically
func (s *Slave) reduceRunningJobs() {
	s.mu.Lock()
	s.runningJobCount--
	s.mu.Unlock()
}
