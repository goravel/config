package queue

import "time"

type Driver interface {
	// Connection returns the connection name for the driver.
	Connection() string
	// Driver returns the driver name for the driver.
	Driver() string
	// Push pushes the job onto the queue.
	Push(job Job, args []any, queue string) error
	// Bulk pushes a slice of jobs onto the queue.
	Bulk(jobs []Jobs, queue string) error
	// Later pushes the job onto the queue after a delay.
	Later(delay time.Duration, job Job, args []any, queue string) error
	// Pop pops the next job off of the queue.
	Pop(queue string) (Job, []any, error)
}
