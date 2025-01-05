package queue

import (
	"time"

	"github.com/goravel/framework/contracts/queue"
)

type Task struct {
	config     *Config
	connection string
	chain      bool
	delay      time.Duration
	driver     *DriverImpl
	jobs       []queue.Jobs
	queue      string
}

func NewTask(config *Config, job queue.Job, args []any) *Task {
	return &Task{
		config:     config,
		connection: config.DefaultConnection(),
		driver:     NewDriverImpl(config.DefaultConnection(), config),
		jobs: []queue.Jobs{
			{
				Job:  job,
				Args: args,
			},
		},
		queue: config.Queue(config.DefaultConnection(), ""),
	}
}

func NewChainTask(config *Config, jobs []queue.Jobs) *Task {
	return &Task{
		config:     config,
		connection: config.DefaultConnection(),
		chain:      true,
		driver:     NewDriverImpl(config.DefaultConnection(), config),
		jobs:       jobs,
		queue:      config.Queue(config.DefaultConnection(), ""),
	}
}

// Delay sets a delay time for the task
func (receiver *Task) Delay(delay time.Duration) queue.Task {
	receiver.delay = delay

	return receiver
}

// Dispatch dispatches the task
func (receiver *Task) Dispatch() error {
	driver, err := receiver.driver.New()
	if err != nil {
		return err
	}

	if receiver.chain {
		return driver.Bulk(receiver.jobs, receiver.queue)
	} else {
		job := receiver.jobs[0]
		if receiver.delay > 0 {
			return driver.Later(receiver.delay, job.Job, job.Args, receiver.queue)
		}
		return driver.Push(job.Job, job.Args, receiver.queue)
	}
}

// DispatchSync dispatches the task synchronously
func (receiver *Task) DispatchSync() error {
	if receiver.chain {
		for _, job := range receiver.jobs {
			if err := job.Job.Handle(job.Args...); err != nil {
				return err
			}
		}

		return nil
	} else {
		job := receiver.jobs[0]

		return job.Job.Handle(job.Args...)
	}
}

// OnConnection sets the connection name
func (receiver *Task) OnConnection(connection string) queue.Task {
	receiver.connection = connection
	receiver.driver = NewDriverImpl(connection, receiver.config)

	return receiver
}

// OnQueue sets the queue name
func (receiver *Task) OnQueue(queue string) queue.Task {
	receiver.queue = receiver.config.Queue(receiver.connection, queue)

	return receiver
}
