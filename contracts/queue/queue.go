package queue

type Queue interface {
	// Chain creates a chain of jobs to be processed one by one, passing
	Chain(jobs []Jobs) Task
	// GetJob get job by signature
	GetJob(signature string) (Job, error)
	// GetJobs get all jobs
	GetJobs() []Job
	// Job add a job to queue
	Job(job Job, args ...[]any) Task
	// Register register jobs
	Register(jobs []Job)
	// Worker create a queue worker
	Worker(payloads ...Args) Worker
}

type Worker interface {
	Run() error
	Shutdown() error
}

type Args struct {
	// Specify connection
	Connection string
	// Specify queue
	Queue string
	// Concurrent num
	Concurrent int
}
