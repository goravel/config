package event

import (
	"github.com/goravel/framework/contracts/event"
	queuecontract "github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/errors"
)

type Task struct {
	args      []event.Arg
	event     event.Event
	listeners []event.Listener
	queue     queuecontract.Queue
}

func NewTask(queue queuecontract.Queue, args []event.Arg, event event.Event, listeners []event.Listener) *Task {
	return &Task{
		args:      args,
		event:     event,
		listeners: listeners,
		queue:     queue,
	}
}

func (receiver *Task) Dispatch() error {
	if len(receiver.listeners) == 0 {
		return errors.EventListenerNotBind.Args(receiver.event)
	}

	handledArgs, err := receiver.event.Handle(receiver.args)
	if err != nil {
		return err
	}

	var mapArgs []any
	for _, arg := range handledArgs {
		mapArgs = append(mapArgs, arg.Value)
	}

	for _, listener := range receiver.listeners {
		var err error
		task := receiver.queue.Job(listener, eventArgsToQueueArgs(handledArgs))
		queue := listener.Queue(mapArgs...)
		if queue.Connection != "" {
			task.OnConnection(queue.Connection)
		}
		if queue.Queue != "" {
			task.OnQueue(queue.Queue)
		}
		if queue.Enable {
			err = task.Dispatch()
		} else {
			err = task.DispatchSync()
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func eventArgsToQueueArgs(args []event.Arg) []any {
	var queueArgs []any
	for _, arg := range args {
		queueArgs = append(queueArgs, arg.Value)
	}

	return queueArgs
}
